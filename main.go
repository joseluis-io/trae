package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/boltdb/bolt"
)

type Configuration struct {
	Path  string
	Delay uint32
}

func main() {

	// Reading configuration file
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Path)
	fmt.Println(configuration.Delay)

	// Open bolt database connection
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create bucket if not exists
	var bucket = []byte("trae")
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(bucket)
		return nil
	})

	// Get user input
	word := os.Args[1]
	// Prepare URL resource
	baseURL := "https://dle.rae.es/" + word

	// check if word exists in cache
	exists := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get([]byte(word))
		fmt.Printf("%s\n", v)
		if v != nil {
			os.Exit(0)
		}
		return err
	})

	// Web Scraping
	client := &http.Client{}
	client.Transport = getTLSConfiguration(client.Transport)

	res, err := client.Get(baseURL)
	checkErr(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	article := doc.Find("article").Text()
	results := doc.Find("div.item-list").Text()

	// Persist word in cache
	if exists == nil {
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
