package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Click on this link to view the spreadsheet
// https://docs.google.com/a/develer.com/spreadsheets/d/1R7xpW4LTGj__B-W8qFKTxYst0LTJBR_AgNcyedbhvwI/edit?usp=sharing

const CSV_URL = "https://docs.google.com/spreadsheets/d/1R7xpW4LTGj__B-W8qFKTxYst0LTJBR_AgNcyedbhvwI/export?format=csv"

func main() {

	res, err := http.Get(CSV_URL)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)

	// STEP 2: implement CSV parsing into data structure, and output to stdout

}
