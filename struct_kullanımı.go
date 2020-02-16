package main

import (
	"fmt"
)

// STRUCT KULLANIMI
/*
	Struct diğer dillerdeki class yapısına karşılık gelir.
*/

// Human ile bir insan sınıfı oluşturduk: Yaş, İsim, Soyisim bilgileri var.
type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//NewHuman bizim yapıcı (constructor) metodumuz olacak.
//Go dilinde constructor default olarak yoktur.
func NewHuman() *Human {
	h := new(Human)
	return h
}

//NewHumanP NewHuman constructor metodunun parametreli halidir. (ad, soyad, yaş)
//GoLang operatör overloading de desteklemez bu nedenle adı farklı olmalı.
func NewHumanP(firstname, lastname string, age int) *Human {
	h := new(Human)
	h.Age = age
	h.FirstName = firstname
	h.LastName = lastname
	return h
}

func main() {
	// //Nesne Örneği Üretmek
	// //v1
	// gaziPasa := Human{
	// 	FirstName: "Mustafa Kemal",
	// 	LastName: "Atatürk",
	// 	Age: 57,
	// }
	// fmt.Println("Yaş:", gaziPasa.Age)

	// //v2
	// insan1 := new(Human)
	// insan1.FirstName = "Abuzer"
	// insan1.LastName = "Kadayıf"
	// fmt.Println(insan1.FirstName, insan1.LastName)

	//v3 ve v4 teki yıldızlar pointer tanımlarken pointer kullanmamızdan kaynaklanıyor.
	//v3 - kendi constructor metodumuz ile
	insan := *NewHuman()
	insan.Age, insan.LastName, insan.FirstName = 23, "İnsan", "Garip"
	fmt.Println(insan)

	//v4 - Parametreli kendi yapıcı metodumuz ile
	insan2 := *NewHumanP("Ali", "Kaya", 35)
	//insan2.FirstName = "Ömer Faruk" // sonradan değiştirebiliriz.
	fmt.Println(insan2)
}
