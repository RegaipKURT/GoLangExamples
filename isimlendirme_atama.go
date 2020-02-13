package main

import (
	"fmt"
)

/*
	Diğer programlama dillerinde bulunan "public" ve "private" gibi belirleyiciler
	go dilinde değişken isminin küçük veya büyük harfle başlamasıyla belirlenir.
	küçük harfle başlayan değişkenlere dışardan erişilemez. sadece aynı paket erişir.
	büyük harfle başlayan değişkenlere paket dışından erişilebilir.
*/

//BURASI GLOBAL ALANDIR!
//BURADA TANIMLANAN DEĞİŞKENLERE AŞAĞIDA YAZILAN FONKSİYONLARIN TAMAMI ERİŞEBİLİR!
//BURADA OLUŞTURULAN DEĞİŞKENLER KULLANILMASA BİLE GO DİLİ HATA VERMEDEN ÇALIŞIR!
var a = "kullanmadığım değişken"

func main() {
	//KULLANMADIĞIMIZ BİR DEĞİŞKEN TANIMLARSAK PROGRAM HATA VERİR VE ÇALIŞMAZ!
	//Verimlilik açısından böyle yapılmış

	// isimlendirme yaparken _ (alt çizgi) kullanılması önerilmez
	//Go dilinde verinin tipi isminden sonra yazılır!

	//1. İSİMLENDİRME YÖNTEMİ
	var dil string = "GoLang"     // başka paketten (package kısmında yazandan başka) erişilemez
	var Version string = "1.13.7" // başka paketlerden eişilebilir.
	fmt.Println(dil, Version, "\n")

	// 2. İSİMLENDİRME YÖNTEMİ
	// var message string // böyle tanımlandığında o tipin varsayılan değeryile ilklendirilir
	// //string için varsayılan değer ==> ""
	// message = "\nMerhaba Go!" // atamayı sonradan yaptık (\n alt satıra geçer)

	//ama const yani sabit tanımlarsak sonradan atama yapamayız! kod hata verir
	// çünkü const değeri sabittir sonradan değişmez. ilk atanan değeri değiştiremeyiz!
	// const yeniMesaj string // HATALI KULLANIM!!!
	// yeniMesaj = "Hello Go!"

	// 3. İSİMLENDİRME YÖNTEMİ
	//tip belirtmeden var keywordu ile tanımlama yapılabilir!
	// var message = "Merhaba Go!" //otomatik olarak string atadı kendisi

	// fmt.Println(message)

	//4. İSİMLENDİRME YÖNTEMİ (ÇOKLU TİP BELİRTME VE ÇOKLU ATAMA)
	// var a, b, c int = 3, 4, 5 //sadece belirtilip sonradan da atama yapılabilir!
	// fmt.Println(a, b, c)
	// şöyle de yapılabilir
	// var a, b, c, d = 3, true, 4.5, "kelime"
	// fmt.Println(a, b, c, d)

	// 5. İSİMLENDİRME (TANIMLAMA) YÖNTEMİ! (EN KISASI)
	// bu şekilde atama sadece fonksiyonların içinde yapılabilir!
	// sayi := 5
	// kelime := "boş kelime"
	// yeni_sayi, yeni_kelime := 4, "yeni kelime"
	// fmt.Println(sayi, kelime)
	// fmt.Println(yeni_sayi, yeni_kelime)

	// 6. İSİMLENDİRME YÖNTEMİ
	// var (
	// 	degisken1 = "kelime"
	// 	degisken2 = 5
	// 	degisken3 = true
	// )
	// fmt.Println(degisken1, degisken2, degisken3)

	//STRİNG İFADELERİN TIRNAKLARA GÖRE AYRIMI
	a := "Şiir yazma hastalığım hep böyle havalarda nüksetti,"
	// "" içinde yazılanlar normal string ifadedir!

	b := `Beni 
     bu güzel 
              havalar mahvetti!`
	// ALT + 96 yazarak olan ` bu işaret arasında çok satırlı string yazılabilir!

	c := ' ' //boşluk karakterinin ascii numarasını döndürecek
	// '' içinde yazılanlar char ifadedir bir tek karakter alır ascii değerini döndürür.

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
