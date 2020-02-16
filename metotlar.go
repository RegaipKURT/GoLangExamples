package main

import (
	"fmt"
)

func main() {

	//1. YOL -- KULLANICI OLUŞTURMA
	user1 := &user{
		firstname: "kemal",
		lastname:  "kara",
		username:  "K.Kara",
		ID:        12343,
		age:       34,
		pay:       &payment{5342.0, 923.0}}

	fmt.Printf("Tam İsim: %s\nKullanıcı Adı: %s\nÖdeme Toplamı: %.2f\n",
		getFullName(*user1), getUserName(user1), getPayment(user1))

	// 2. YOL -- KULLANICI OLUŞTURMA
	user2 := newUser()
	user2.firstname = "Eşref"
	user2.lastname = "Sencer"
	user2.username = "GODER.CO"
	user2.pay = newPayment()
	user2.pay.salary = 3213.5
	user2.pay.bonus = 678.3

	fmt.Printf("\nTam İsim: %s\nKullanıcı Adı: %s\nÖdeme Toplamı: %.2f\n",
		getFullName(*user2), getUserName(user2), getPayment(user2))

}

//KULLANICI SINIFI VE METOTLARI
type user struct {
	ID        int
	username  string
	firstname string
	lastname  string
	age       int
	pay       *payment
}

//kullanıcı yapıcı metodu
func newUser() *user {
	u := new(user)
	u.pay = newPayment()
	return u
}
func getFullName(u user) string {
	return u.firstname + " " + u.lastname
}
func getUserName(u *user) string {
	return u.username
}

//PAYMENT SINIFI VE METOTLARI.
type payment struct {
	salary float64
	bonus  float64
}

//payment için yapıcı metod
func newPayment() *payment {
	p := new(payment)
	return p
}
func getPayment(u *user) float64 {
	return u.pay.salary + u.pay.bonus
}
