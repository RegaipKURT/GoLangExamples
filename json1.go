package main

import (
	"encoding/json"
	"fmt"
)

/*
json.Unmarshal json verisini go verisine döndürmeye yararken,
json.Marshal go verisini json verisine döndürmeye yarar.
*/
func main() {
	jsonString := `
		{
			"data": {
				"object": "card",
				"id" : "regkrt3203498",
				"firstname": "Haydar",
				"lastname": "Bozkurt",
				"balance": "26.780"
			}
		}
	`
	var m map[string]map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &m); err != nil {
		panic(err.Error())
	}
	// fmt.Println(m)
	// d := m["data"]
	// fmt.Println(d)
	// fmt.Println(d["id"])

	// m yi marshal metoduyla byte dizisi olarak alıp string yapabiliriz. böylece JSON formatında görünür
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
