package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	GO ROUTINELER'DEN THREAD DİYE BAHSETSEK BİLE ARKA PLAN MANTIĞI DAHA FARKLIDIR.
	GO DİLİNDE GOROUTINE LER ASLINDA İŞLETİM SİSTEMİ SEVİYESİNDE BİR THREAD DEĞİLDİR.
	BİR THREAD ALTINDA ÇALIŞAN FARKLI THREAD'LER GİBİDİR. (THREAD MANTIĞINDAN FARKLI BİRAZ!)
	BİR THREAD ÜZERİNDEN OLUŞTURULAN GOROUTINELERDİR. BU YÜZDEN, YANİ ASLINDA İŞLETİM
	SİSTEMİ SEVİYESİNDE GERÇEK BAŞKA THREAD OLMADIĞINDAN ÇOK HAFİFTİR.


				            GO ROUTINE VEYA THREAD KULLANIMI
	GO PROGRAMLAMA DİLİNDE KENDİ YAPISINA ÖZEL GO ROUTINE - THREADLER VARDIR.
	KENDİ YAPISINA ÖZEL ÇÜNKÜ, DİĞER DİLLERE ÇOK BENZESE DE ARKA PLANDA ÖZEL
	VE GÜÇLÜ BİR YAPI SUNAR. BU ÖZEL YAPI BİR TEK THREAD İÇİNDEN BİNLERCE THREAD ÇAĞRILMASINI
	DESTEKLEYEN HAFİF VE GÜÇLÜ BİR ALTYAPIYA SAHİPTİR.

				                  GO ROUTINE KULLANIMI
	GO ROUTINLER VEYA THREADLER "GO" KEYWORDU İLE BAŞLATIRLIR. "GO" KEYWORDÜ BİR FONKSİYONUN
	BAŞINA GELEREK ARTIK ONU BAŞKA BİR THREAD HALİNE GETİRİR VE O FONKSİYON AYRI BİR GO
	ROUTINE VEYA THREAD OLARAK ÇALIŞIR.
*/

func main() {
	// THREAD ŞEKLİNDE OLAN KULLANIM
	fmt.Println("---------------------- THREAD PROGRAMLAMA ----------------------")
	fmt.Println("Önce Yazdığım yazı")
	go harfler()
	fmt.Println("Sonra Yazdığım yazı")
	fmt.Println("Main: Waiting for worker to finish\n")
	time.Sleep(1 * time.Second) // thread şelindeki kısım bitene kadar bekleyelim.

	fmt.Println("\n\n---------------------- PARALEL PROGRAMLAMA ----------------------")
	runtime.GOMAXPROCS(8)       //MAXİMUM KAÇ İŞLEMCİ EŞZAMANLI KULLANILACAK.
	go harflerASYNC()           //burada fonksiyonda yapılan işlemlerde de go keywordü var.
	time.Sleep(1 * time.Second) //EĞER BEKLETMEZSEK MAIN FONKSİYON BİTİNCE PROGRAM KAPANIR.
	//eğer main thread threadler bitene kadar beklesin istersek channel kullanabiliriz.
	fmt.Println("Main: Tamamlandı!")
}

func harfler() {
	for l := byte('a'); l <= byte('z'); l++ {
		fmt.Print(string(l) + ", ")
	}
	fmt.Println("\nHarfler fonksiyonu tamamlandı!")
}

//PARALEL PROGRAMLAMADA HANGİ İŞLEMİN ÖNCE BİTECEĞİNİ GARANTİ EDEMEZSİNİZ.
// HANGİSİ ÖNCE BİTERSE ÖNCE ORASI DÖNECEKTİR. AŞAĞIDAKİ FONKSİYONDA,
// HARFLERİN SIRASI KARIŞIK ÇIKABİLİR VE TAMAMLANDI YAZISI FOR DÖNGÜSÜNDEN ÖNCE YAZILACAKTIR!
func harflerASYNC() {
	//BURADA İLK ÖNCE HARFLER TAMAMLANDI YAZACAK. ÇÜNKÜ ÖNCE BİTECEK.
	// YAZDIĞI HARFLER DE KARIŞIK BİR SIRAYLA YAZILMIŞ OLABİLİR.
	for l := byte('a'); l <= byte('z'); l++ {
		go fmt.Print(string(l) + ", ")
	}
	go fmt.Println("HarflerASYNC fonksiyonu tamamlandı!")
}
