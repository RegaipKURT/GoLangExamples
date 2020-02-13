package main

import (
	"fmt"
	"math/rand"
	"time"
)

// zaman değerini kullanarak rasgeleliği artıracağız!
//kaynağa dayalı şekilde random değer üreteceğiz
var source = rand.NewSource(time.Now().UnixNano())

//şifremiz bu karakterler arasından seçilecek sembollerle oluşturulacak!
const _charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789*!+.()&%#[]{}=;:"

func main() {
	//sifre üretme fonksiyonunu çağırıp sifre değerine atayalım ve yazdıralım.
	sifre := GeneratePassword(32) // 32 karakterli bir string dönecek
	fmt.Println("Üretilen Şifre:", sifre)
}

// GeneratePassword : Uzunluk(length) değeri alıp o uzunlukta rasgele şifre üretir.
// Dönüş tipi string'dir
func GeneratePassword(length int) string {
	x := make([]byte, length) //x byte dizisi içine ekleyeceğiz seçtiğimiz karakterleri
	for i := range x {
		x[i] = _charset[source.Int63()%int64(len(_charset))]
	}
	return string(x) //byte dizisini stringe çevirip döndürüyoruz!
}
