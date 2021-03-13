package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Currency struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}

func main() {
	var c Currency
	if len(os.Args) > 1 {
		if strings.ToUpper(os.Args[1]) == "HELP" || strings.ToUpper(os.Args[1]) == "--HELP" || strings.ToUpper(os.Args[1]) == "-HELP" || strings.ToUpper(os.Args[1]) == "-H" {
			helpExit()
		}
	}
	if len(os.Args) == 2 {
		base := strings.ToUpper(os.Args[1])
		resp, _ := http.Get("https://api.exchangeratesapi.io/latest?base=" + base)
		b, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(b, &c)

		if len(c.Rates) >= 1 {
			for i, j := range c.Rates {
				fmt.Println(i+": ", j)
			}
			fmt.Println("\nBase: ", c.Base, "\tDate: ", c.Date)
		} else {
			fmt.Println("Currency " + base + " not found!")
			helpExit()
		}
	} else if len(os.Args) == 3 {
		base := strings.ToUpper(os.Args[1])
		to := strings.ToUpper(os.Args[2])
		resp, _ := http.Get("https://api.exchangeratesapi.io/latest?base=" + base)
		b, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(b, &c)
		if len(c.Rates) >= 1 && c.Rates[to] != 0 {
			fmt.Println(base+" to "+to+": ", c.Rates[to], "\nDate: ", c.Date)
		} else {
			fmt.Println("Currency " + base + " or " + to + " not exists!")
			helpExit()
		}
	} else if len(os.Args) == 1 {
		resp, _ := http.Get("https://api.exchangeratesapi.io/latest")
		b, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(b, &c)

		for i, j := range c.Rates {
			fmt.Println(i+": ", j)
		}
		fmt.Println("\nBase: ", c.Base, "\tDate: ", c.Date)
	} else {
		helpExit()
	}

}

func helpExit() {
	fmt.Println("Usage: <program_name> or\nUsage: <program_name> <base> or\nUsage: <program_name> <from> <to> ")
	os.Exit(1)
}
