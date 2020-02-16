package main

import "fmt"

/*
Channel goroutine'ler arasında haberleşmeyi sağlayan bir yapıdır!
*/

var toplanacakListe = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var toplanacakListeFarklı = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

// main fonksiyonunun içinde farklı listeyi farklı bir kanaldan kullanarak
// channel değerlerin karışmadığını göstermiş olacağız.
func main() {
	fmt.Println("Main Başladı!")
	farklıChannel := make(chan int) //bu kanalı karışmayacak işlem için kullanacağız
	c := make(chan int)
	go sum(toplanacakListeFarklı, farklıChannel)
	go sum(toplanacakListe[:len(toplanacakListe)], c)
	go sum(toplanacakListe[:len(toplanacakListe)-1], c)
	z := <-farklıChannel //EN BAŞTA OLMASINA RAĞMEN FARKLI KANALDAN BEKLEDİĞİ İÇİN İLK BİTENİ ALMAZ.
	y := <-c             // daha küçük olması lazım çünkü kısa listenin önce bitmesi gerekir.
	//y ilk aldığımız değer daha önce işi biten kanaldan gelir.
	x := <-c // büyük olmalı çünkü 2 katı kadar uzun bir listede işlem yapacak
	//x sonra aldığımız değer olduğu için sonra işi biten kanaldan gelen değer olacak
	//yukarda özetle ilk gelen c hangisi olursa olsun y'ye atadık.
	// beklenen değer kısa listeden gelen değer olması.
	//Çünkü daha kısa bir işlem ve daha önce bitmesi gerekir.
	//Fakat y'nin daha küçük olması GARANTİ DEĞİLDİR!
	//y'nin küçük listeden gelen değer olduğunun kesin olması için,
	// farklı isimde bir tane daha kanal kullanmak gerekir.
	//O fonksiyona başka kanal verip onun değerini bekleriz. Dolayısıyla karışmaz.
	fmt.Println("Küçük olması gereken y:", y)
	fmt.Println("Büyük olması gereken x:", x)

	//Buradaki fonksiyon farklı kanal kullanduığı için karışma ihtimali yoktur!
	fmt.Println("Farklı Listenin toplamı:", z)

	fmt.Println("Main Bitti!")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum
}
