package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
)

func main() {

	baseURL := "https://dle.rae.es/"

	client := &http.Client{}

	client.Transport = getTLSConfiguration(client.Transport)

	response, err := client.Get(baseURL)
	checkErr(err)

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	response.Body.Close()
}

func checkErr(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

