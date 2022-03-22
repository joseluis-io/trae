package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
    "github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
    // See if the href attribute exists on the element
    href, exists := element.Attr("href")
    if exists {
        fmt.Println(href)
    }
}

// Custom user agent.
const (
    userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) " +
        "AppleWebKit/537.36 (KHTML, like Gecko) " +
        "Chrome/53.0.2785.143 " +
        "Safari/537.36"
)

const url = "https://dle.rae.es/hola?m=form"

func main() {
	// HTTP Request


	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("failed to initiate request to rae: %v", err)
	}

	// Set HTTP User-Agent to server think is a user
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("failed to make request to rae: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("failed to read response body: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Errorf("failed to parse response from rae: %v", err)
	}

	doc.Find("a").Each(processElement)
}
