package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
							SELECT

Go dilinde kanallar ile çalışırken kullanılan select isimli bir ifade bulunur.
Kullanımı switch case yapısına benzer ama case kısmı bir kanal işleminin sonucunu bekler.

select {
case 1:
	ifade
case2:
	ifade
}

Yukarıdaki sözdizimi ile kullanılır. İstenirse aşağıdaki örnekte gösterildiği gibi
default durumu da eklenebilir.

Select ifadesi eğer default durum belirtilmediyse case ifadelerinde belirtilen kanallardan
en az birisi sonuç dönene kadar bekler. select ifadesi çalıştığında iki kanal da sonuç dönmüşse
sonuçlardan veya case durumlarından biri rastgele seçilir.

Default durum belirtildiğinde ise, eğer kanallardan hiçbiri değer dönmediyse çalıştırılacak ifadeyi
belirlemek için kullanılır.

Örneğin sorgulara hızlı cevap vermek için iki tane veritabanı server kullandığımızı düşünelim.
Aynı sorguyu gönderdiğimiz durumda iki kanaldan da sorguya cevap dönmesini bekleriz.
Fakat biz bu durumu hız için uyguladığımızdan, önce dönen cevap bizim için yeterlidir.
İkinci cevabı beklemeye ihtiyacımız yoktur.

Bu durumda şöyle bir select ifadesi kullanabiliriz:
select{
case sonuc1:= <-sorgu1:
	fmt.Println(sonuc1)
case sonuc2:= <- sorgu2:
	fmt.Println(sonuc2)
}
Yukarıdaki durumda önce dönen değer hangi kanaldan geldiyse o yazdırılır ve program akışına devam eder.
Eğer aşağıdaki gibi bir default durum kullanırsak:
select{
case sonuc1:= <-sorgu1:
	fmt.Println(sonuc1)
case sonuc2:= <- sorgu2:
	fmt.Println(sonuc2)
default:
	fmt.Println("Sonuçlar yetiştirilemedi!")
}
Eğer select ifadesine girildiğinde kanallardan ikisi de bir değer alamamışsa
o zaman default kısmı çalışacaktır.
*/

//MakeRandomInt bir kanala random integer değer atamak için kullanılacak
func makeRandomInt(c chan int) {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(4 * time.Nanosecond)
	a := rand.Intn(500)
	c <- a
	close(c)
}

//MakeRandomInt bir diğer kanala random integer değer atamak için kullanılacak
func makeRandomInt2(c chan int) {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(4 * time.Nanosecond)
	a := rand.Intn(100)
	c <- a
	close(c)
}

func main() {
	//iki integer kanal oluşturuyoruz.
	c1 := make(chan int)
	c2 := make(chan int)

	//fonksiyonlarımızı goroutine olacak şekilde bağımsız olarak çalıştırıyoruz.
	go makeRandomInt(c1)
	go makeRandomInt2(c2)

	//Yukarıdaki fonksiyonlar 4 nanosaniye beklerken, biz 2 nanosaniye bekliyoruz.
	time.Sleep(2 * time.Nanosecond)

	//Select ifadesi eğer kanallara bir değer verilmişse yazdıracak.
	//Yoksa c1 kanalının sonucunu bekleyip yazdıracak.
	select {
	case p1 := <-c1:
		fmt.Println("c1 alındı! Değeri: ", p1)
	case p2 := <-c2:
		fmt.Println("c2 alındı! Değeri: ", p2)
	default:
		fmt.Println("Değer için default beklendi! ", <-c1)
	}
}
