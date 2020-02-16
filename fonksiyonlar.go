package main

import (
	"fmt"
	"strings"
)

func main() {
	mesaj := "Naber"
	fmt.Println(mesaj) // çıktı: Naber

	//mesajı içerde değiştirmemize rağmen aynı kalacak.
	//çünkü fonksiyon orjinal değeri kopyalayarak çalışır.
	degistirmeyen(mesaj)
	fmt.Println(mesaj) //çıktı : Naber

	//ama pointer üzerinden değiştirirsek hafızada o değişkenin değerini alıp değiştirecek.
	degistir(&mesaj)
	fmt.Println(mesaj) //çıktı: Değişen mesaj

	//birden fazla geri dönüşlü fonksiyon
	var x, y = "Önce", "Sonra"
	x, y = yerDegis(x, y)
	fmt.Println(x, y)

	//belirsiz sayıda parametre alan fonksiyon
	sonuc, uzunluk := topla(3, 5, 67, 3, 345, 432, 543, 4, 6, 1, 44)
	fmt.Printf("%d elemanlı listenin Toplam sonucu: %d \n", uzunluk, sonuc)

	//isimlendirilmiş dönüş yapan fonksiyon
	varMı, sıklık := bul("kemal", "arkadaş", "mustafa kemal", "kemaliyet", "dürüst", "insan")
	fmt.Printf("Geçme Durumu: %v, Geçme Sıklığı: %d \n", varMı, sıklık) //beklenen 2

	//Anonim Fonksiyonlar.
	birlestir := func(degerler ...string) (donus string) {
		for _, d := range degerler {
			donus += d
		}
		return
	}
	topluYazi := birlestir("ali", "veli", "deli")
	fmt.Println(topluYazi)
}

//Fonksiyonlar normalde içlerine aldığı değişkeni kopyalaıp işlem yapar.
// Yani değişkeni içerde değiştirsek bile orjinal değişken aynı kalır.
// Ama değişkeni değiştirmek istiyorsak pointer ile işlem yapabiliriz.
func degistir(mesaj *string) {
	*mesaj = "Değişen Mesaj!"
	//işlem mesaj değişkeninin hafızadaki adresi üzerinde gerçekleşti.
}

//değiştimediğini görmek için yeniden deneyelim.
func degistirmeyen(mesaj string) {
	mesaj = "Değişmeyen mesaj"
	//işlem kopyalanan bir mesaj değişkeni üzerinde gerçekleşti.
}

//birden fazla geri dönüş yapan fonksiyonlar (istediğiimiz kadar yazabiliriz.)
func yerDegis(x string, y string) (string, string) {
	return y, x
}

//belirsiz sayıda parametre alan fonksiyon
func topla(sayilar ...int) (int, int) {
	toplam := 0
	for _, deger := range sayilar {
		toplam = toplam + deger
	}
	return toplam, len(sayilar)
}

//isimlendirilmiş geri dönüş yapan fonksiyonlar
//ilk verilen değer sonradan gelenlerin arasında geçiyor mu? geçiyorsa kaç defa geçiyor?
func bul(taranacakDeger string, isimler ...string) (sonuc bool, toplam int) {
	for _, deger := range isimler {
		if strings.Contains(deger, taranacakDeger) {
			sonuc = true
			toplam++
		}
	}
	return //belirtmeden sonuc ve toplamı döndürdü.
}
