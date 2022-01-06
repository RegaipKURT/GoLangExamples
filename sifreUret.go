package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// zaman değerini kullanarak rasgeleliği artıracağız!
//kaynağa dayalı şekilde random değer üreteceğiz
var source = rand.NewSource(time.Now().UnixNano())

//şifremiz bu karakterler arasından seçilecek sembollerle oluşturulacak!
const _charsetLow = "abcdefghijklmnopqrstuvwxyz"
const _charsetMidLow = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const _charsetMidHigh = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const _charsetHigh = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789*!+.()&%#[]{}=;:"

var L, C int

func main() {
	//şifre uzunluğunu argüman alma

	L := flag.Int("length", 4, "Length of password")
	//şifre kompleksliği argümanı
	C := flag.Int("complex", 4, "Complexity of password")
	flag.Parse()
	fmt.Println(*L, *C)
	//sifre üretme fonksiyonunu çağırıp sifre değerine atayalım ve yazdıralım.
	sifre := GeneratePassword(*L, *C) // karakterli bir string dönecek
	fmt.Println("Üretilen Şifre:", sifre)
}

// GeneratePassword : Uzunluk(length) değeri alıp o uzunlukta rasgele şifre üretir.
// Dönüş tipi string'dir
func GeneratePassword(length, complexity int) string {
	x := make([]byte, length) //x byte dizisi içine ekleyeceğiz seçtiğimiz karakterleri
	if complexity <= 1 {
		for i := range x {
			x[i] = _charsetLow[source.Int63()%int64(len(_charsetLow))]
		}
		return string(x)
	}
	if complexity == 2 {
		for i := range x {
			x[i] = _charsetMidLow[source.Int63()%int64(len(_charsetMidLow))]
		}
		return string(x)
	}
	if complexity == 3 {
		for i := range x {
			x[i] = _charsetMidHigh[source.Int63()%int64(len(_charsetMidHigh))]
		}
		return string(x)
	}
	if complexity >= 4 {
		for i := range x {
			x[i] = _charsetHigh[source.Int63()%int64(len(_charsetHigh))]
		}
		return string(x)
	}
	return string(x) //byte dizisini stringe çevirip döndürüyoruz!
}
