package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	word := "ambiguo"
	baseURL := "https://dle.rae.es/" + word

	client := &http.Client{}

	client.Transport = getTLSConfiguration(client.Transport)

	res, err := client.Get(baseURL)
	checkErr(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	article := doc.Find("article").Text()

	fmt.Println(article)
}

func checkErr(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

