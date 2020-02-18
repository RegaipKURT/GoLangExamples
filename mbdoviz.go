package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Currency struct {
	Unit            int
	Isim            string
	CurrencyName    string
	ForexBuying     float64
	ForexSelling    float64
	BanknoteBuying  string
	BanknoteSelling Banknote
}

type Banknote struct {
	CrossRateUSD   float64
	CrossRateOther float64
}

func main() {
	resp, err := http.Get("https://www.tcmb.gov.tr/kurlar/today.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	data_n := string(data[93:])
	var kurlar struct {
		Tarih_Date []Currency `xml:"Currency"`
	}
	if err := xml.Unmarshal([]byte(data_n), &kurlar); err != nil {
		log.Fatal(err)
	}
	if len(os.Args[1:]) == 0 {
		for _, c := range kurlar.Tarih_Date {
			fmt.Println(c.Isim, "-  Alış:", c.ForexBuying, "-Satış:", c.ForexSelling)
		}
	} else {
		for _, isim := range os.Args[1:] {
			for _, c := range kurlar.Tarih_Date {
				if isim == c.Isim {
					fmt.Println(c.Isim, "-  Alış:", c.ForexBuying, "-Satış:", c.ForexSelling)
				}
			}
		}
	}
}