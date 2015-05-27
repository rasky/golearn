package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Click on this link to view the spreadsheet
// https://docs.google.com/a/develer.com/spreadsheets/d/1R7xpW4LTGj__B-W8qFKTxYst0LTJBR_AgNcyedbhvwI/edit?usp=sharing

const CSV_URL = "https://docs.google.com/spreadsheets/d/1R7xpW4LTGj__B-W8qFKTxYst0LTJBR_AgNcyedbhvwI/export?format=csv"

type Person struct {
	Nickname string
	Name     string
	Surname  string
	Birthday time.Time
	Tags     []string
}

func main() {

	res, err := http.Get(CSV_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// STEP 2: implement CSV parsing into data structure, and output to stdout
	records, err := csv.NewReader(res.Body).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var people []Person

	for _, r := range records {

		bday, err := time.Parse("02/01/2006", r[3])
		if err != nil {
			log.Println("WARNING: invalid birthday format, skipping record", r[3])
			continue
		}

		p := Person{
			Nickname: r[0],
			Name:     r[1],
			Surname:  r[2],
			Birthday: bday,
			Tags:     strings.Split(r[4], ","),
		}
		people = append(people, p)
	}

	fmt.Println(people)
}
