package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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
	resp, err := http.Get("http://www.tcmb.gov.tr/kurlar/today.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	data_n := string(data)

	var kurlar struct {
		Tarih_Date []Currency `xml:"Currency"`
	}

	if err := xml.Unmarshal([]byte(data_n), &kurlar); err != nil {
		log.Fatal(err)
	}

	if len(os.Args[1:]) == 0 {
		for _, c := range kurlar.Tarih_Date[:len(kurlar.Tarih_Date)-1] {
			fmt.Println(c.Isim, "("+c.CurrencyName+")", " Alış:",
				c.ForexBuying, "- Satış:", c.ForexSelling)
		}
		for _, c := range kurlar.Tarih_Date[len(kurlar.Tarih_Date)-1 : len(kurlar.Tarih_Date)] {
			fmt.Println("Special Drawing Rights (SDR)", " Alış:",
				c.ForexBuying, "- Satış:", c.ForexSelling)
		}
	} else {
		for _, isim := range os.Args[1:] {
			for _, c := range kurlar.Tarih_Date {
				if strings.ToUpper(isim[:2]) == c.Isim[:2] || strings.ToUpper(isim[:2]) == c.CurrencyName[:2] {
					fmt.Println(strings.Replace(c.Isim, "  ", "", -1), "("+strings.Replace(c.CurrencyName, "  ", "", -1)+")", " Alış:",
						c.ForexBuying, "- Satış:", c.ForexSelling)
				}
			}
		}
	}
}
