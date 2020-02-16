package main

import (
	"fmt"
)

func main() {
	araba := new(ferrari) // ferrari ile diğer özelliklere de erişebileceğiz.
	araba.marka = "Ferrari"
	araba.model = "V350"
	araba.renk = "Kırmızı"
	araba.hız = 450
	araba.isSpecial = true
	araba.fiyat = 249999.99

	araba.information()
}

// KALITIM KULLANIMI
type car struct {
	id    string
	marka string
	model string
	renk  string
	hız   int
	fiyat float64
}

type specialProduction struct {
	isSpecial bool
}

//ferrari araba tanımlanıyıp kalıtım uygulayalım
// composition dediğimiz isimsiz sınıf tanımlama yöntemiyle kalıtım yapıp ebeveyn sınıfların
// özelliklerine eriişebiliyor olacağız.
type ferrari struct {
	car //composition (golang içinde bu kullanıma embedded type denir.)
	specialProduction
}

// INTERFACE KULLANIMI
type carFace interface {
	Run() bool
	Stop() bool
	information() string
}

//Run carFace için örnek metodumuz olsun.
func (ferrari) Run() bool {
	return true
}

//Stop carface için durdurma metodumuz.
func (ferrari) Stop() bool {
	return false
}

func (araba ferrari) information() {
	fmt.Printf("Marka: %s\nModel: %s\nRenk: %s\nHız: %d\nFiyat: %.2f\nÖzel Üretim: %v\n",
		araba.marka, araba.model, araba.renk, araba.hız, araba.fiyat, araba.isSpecial)
}
