package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Name struct {
	firstName string
	lastName  string
}

type Email struct {
	ID     int
	kind   string
	adress string
}

type Interests struct {
	Name string
}

type Person struct {
	Name     Name
	Email    []Email
	Interest []Interests
}

func GetPersonFullName(p Person) string {
	return p.Name.firstName + " " + p.Name.lastName
}

func GetPersonEmailAdd(p Person, i int) string {
	return p.Email[i].adress
}

func GetPersonEmails(p Person) []Email {
	return p.Email
}
func GetPersonEmail(p Person, i int) string {
	return p.Email[i].adress
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func SaveFile(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
	fmt.Println("Yazma işlemi Başarılı!")
}

func main() {
	kisi1 := Person{
		Name: Name{
			firstName: "Regaip",
			lastName:  "Kurt"},
		Email: []Email{Email{
			ID:     1,
			kind:   "GMAIL",
			adress: "mailimadresim.gmail.com"},
			Email{
				ID:     2,
				kind:   "WORK",
				adress: "garipcalisan@isadresi.com"}},
		Interest: []Interests{Interests{Name: "Programlama"},
			Interests{Name: "SiberGüvenlik"}}}
	fmt.Println(kisi1)
	fmt.Println(GetPersonFullName(kisi1), GetPersonEmailAdd(kisi1, 1))
	SaveFile("person1.json", kisi1)
}
