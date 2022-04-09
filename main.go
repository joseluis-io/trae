package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	word := os.Args[1]
	baseURL := "https://dle.rae.es/" + word

	// TODO: check if word exists in cache

	client := &http.Client{}
	client.Transport = getTLSConfiguration(client.Transport)

	res, err := client.Get(baseURL)
	checkErr(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	article := doc.Find("article").Text()
	results := doc.Find("div.item-list").Text()

	// TODO: persist word in cache

	if article != "" {
		fmt.Println(article)
	}else{
		fmt.Println(results)
	}


}

func checkErr(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

