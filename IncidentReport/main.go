package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Austin interface {
	AFD
	APD
}

type AFD struct {
	title       string
	coordinates string
	description string
	pubDate     string
}

type APD struct {
	title       string
	coordinates string
	description string
	pubDate     string
}

func main() {
	respAFD, errAFD := http.Get("https://services.austintexas.gov/fact/fact_rss.cfm")
	respAPD, errAPD := http.Get("https://services.austintexas.gov/qact/qact_rss.cfm")

	if errAFD != nil {
		log.Fatal(errAFD)
		fmt.Println("Error:", errAFD)

	}

	if errAPD != nil {
		log.Fatal(errAPD)
		fmt.Println("Error:", errAPD)

	}

	bodyAFD, errAFD := io.ReadAll(respAFD.Body)
	respAFD.Body.Close()
	if respAFD.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", respAFD.StatusCode, bodyAFD)

	}

	if errAFD != nil {
		log.Fatal(errAFD)

	}
	fmt.Printf("%s", bodyAFD)

	bodyAPD, errAPD := io.ReadAll(respAPD.Body)
	respAPD.Body.Close()
	if respAPD.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", respAPD.StatusCode, bodyAPD)

	}

	if errAPD != nil {
		log.Fatal(errAPD)

	}
	fmt.Printf("%s", bodyAPD)

}
