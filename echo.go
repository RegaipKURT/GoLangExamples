package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//ECHO KOMUTU GİBİ ALDIĞI DEĞERİ EKRANA YAZDIRAN PROGRAM!
	//başlangıç zamanı tanımlayalım.
	time0 := time.Now() // iki yöntem deneyip zamana göre performansına bakacağız.
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	time1 := time.Now() //ilk kodun bitiş zamanı, aynı zamanda ikinci kodun başlangıcı.

	//aynı kodu for döngüsü kullanmadan strings paketi kullanarak da yazabiliriz!
	fmt.Println(strings.Join(os.Args[1:], " "))
	time2 := time.Now() //ikinci kodun bitiş zamanı

	// iki yöntemin aldığı zamanları yazdıralım.
	//birinci yöntem for ile yaptığı için verimsizdir aslında
	fmt.Println("Inefficient Version takes time: ", time1.Sub(time0))
	// ikinci yöntem gömülü paketlerle çalıştığından daha hızlı olacaktır.
	fmt.Println("Strings version takes time: ", time2.Sub(time1))
	/*
		PROGRAM BİRKAÇ KEZ ÇALIŞTIRILIP DEĞERLERE BAKILDIĞINDA GÖMÜLÜ FONKSİYON KULLANMANIN
		4-5 KAT DAHA HIZLI SONUÇLAR VERDİĞİ GÖRÜLECEKTİR. GO DİLİ EN VERİMLİ ŞEKİLDE ÇALIŞMAK
		İÇİN TASARLANDIĞINDAN GÖMÜLÜ FONKSİYONLARDA YÜKSEK PERFORMANSLA ÇALIŞMAK ÜZERE
		TASARLANMIŞTIR. GÖMÜLÜ FONKSİYON KULLANMAK PROGRAMINIZIN DAHA VERİMLİ ÇALIŞMASINI SAĞLAR.
	*/
}
