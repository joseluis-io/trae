package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/boltdb/bolt"
)

func main() {

	// get config file
	configFile := readConfigFile()
	defer configFile.Close()

	// parse config file
	configuration := parseConfigFile(configFile)

	// Open bolt database connection
	db := openDatabase(*configuration)
	defer db.Close()

	var bucket = []byte("trae")
	// create bucket if not exists
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(bucket)
		return nil
	})

	progArgs := len(os.Args)
	if progArgs == 1 {
		fmt.Println("Comando no válido, si necesitas ayuda prueba con:\n\ttrae -h")
		os.Exit(0)
	}
	// Get user input
	word := os.Args[1]

	if word == "-h" {
		fmt.Println("Puede llamar a los siguientes comandos:\n\n\ttrae <palabra>\n\ttrae -h")
		fmt.Println("\nPara más información acceda a https://github.com/jl-hoz/trae")
		os.Exit(0)
	}

	// Prepare URL resource
	baseURL := "https://dle.rae.es/" + url.QueryEscape(word)

	// check if word exists in cache
	exists := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get([]byte(word))
		fmt.Printf("%s\n", v)
		if v != nil {
			os.Exit(0)
		}
		return nil
	})

	// Web Scraping
	client := &http.Client{}
	client.Transport = getTLSConfiguration(client.Transport)

	res, err := client.Get(baseURL)
	if err != nil {
		fmt.Printf("Error al solicitar la siguiente URL: %s\n\n", baseURL)
		fmt.Printf("Posibles causas del error: \n\n")
		fmt.Println("1. No está conectado a internet.")
		fmt.Println("2. El servicio de la RAE web no está disponible.")
		os.Exit(0)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	article := doc.Find("article").Text()
	results := doc.Find("div.item-list").Text()

	// Persist word in cache
	if configuration.DatabaseStore && exists == nil {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket(bucket)
			err := b.Put([]byte(word), []byte(article))
			return err
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	if article != "" {
		fmt.Println(article)
	} else {
		fmt.Println(results)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
