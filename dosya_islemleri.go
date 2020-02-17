package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//main içinde de oluşturabiliriz bu değişkenleri ama şimdi oluşturalım.
var (
	newFile  os.File
	fileInfo os.FileInfo
	err      error
)

func main() {
	//DOSYA OLUŞTURMA
	newFile, err := dosyaOlustur("deneme1.txt")
	defer newFile.Close()
	if err == nil {
		log.Println(newFile.Name())
	} else if err != nil {
		log.Fatal(err)
	}

	// //DOSYA BİLGİSİ OKUMA
	// fileInfo, err = os.Stat("deneme1.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Dosya Adı: ", fileInfo.Name())
	// fmt.Printf("Dosya Boyutu: %d byte\n", fileInfo.Size())
	// fmt.Println("Dosya İzinleri: ", fileInfo.Mode())
	// fmt.Println("Değiştirilme Tarihi: ", fileInfo.ModTime())
	// fmt.Println("Klasör mü? : ", fileInfo.IsDir())
	// fmt.Printf("Sistem Bilgileri: %v\n", fileInfo.Sys())

	//DOSYA İSMİ DEĞİŞTİRME VE TAŞIMA (ikisi de rename ile yapılır.)
	err = os.Rename("deneme1.txt", "yeniDosya.txt") //isim değiştirme
	if err != nil {
		fmt.Printf("İsim değiştirme başarısız: %v\n", err.Error())
		os.Exit(2)
	} else {
		fmt.Println("İsim başarıyla değiştirildi.")
	}

	// //DOSYA TAŞIMA
	// newPath := "/home/pars/yeniDosya.txt" //burası değiştirilmezse hata verecektir!
	// err = os.Rename("yeniDosya.txt", newPath)
	// if err != nil {
	// 	fmt.Printf("Taşıma işlemi başarısız: %v\n", err.Error())
	// 	os.Exit(2)
	// } else {
	// 	fmt.Println("Taşıma işlemi başarılı: ", newPath)
	// }

	// //DOSYA VARLIK KONTROLÜ
	// fileinfo, errf := os.Stat("demo.txt")
	// if os.IsNotExist(errf) {
	// 	//bazı durumlarda da os.IsExist() üzerinden bakmamız gerekebilir. (ör:dosya açarken)
	// 	log.Fatal("File not found!")
	// }
	// fmt.Println(fileinfo)

	// //DOSYA KAPATMA! (KAPATMADIĞIMIZDA İŞLETİM SİSTEMİ KUILLANILIYORMUŞ GİBİ GÖREBİLİR.)
	// file, _ := os.Open("yeniDosya.txt")
	// //işlemler yapıldı
	// //file.Close()
	// defer file.Close() //bütün main fonksiyonu sonlandıktan sonra dosyayı kapatacak!

	// //DOSYAYI ÖZEL AYARLARLA AÇMAK! os.OpenFile()
	// //ikinci parametre dosya açılış amacı ve ayarları
	// //üçüncü parametre dosya izinlerini belirler
	// newFile, err = os.OpenFile("yeniDosya.txt", os.O_CREATE|os.O_RDWR, 0666)
	// //ctrl ile os.ayar yazan parametreye basıp orjinaline giderseniz açıklamaları var!
	// defer newFile.Close()
	// fileinfo, errf := os.Stat("yeniDosya.txt")
	// if errf != nil {
	// 	fmt.Println(errf)
	// }
	// fmt.Println(fileinfo)

	// //OKUMA VE YAZMA İZİNLERİNİ TEST ETMEK
	// f, ferr := os.OpenFile("yeni.txt", os.O_WRONLY, 0666) //yazma izniyle açmayı denedik.
	// defer f.Close()
	// //biraz daha detaylı bir hata anlizi yapalım
	// if ferr != nil { //hata varsa
	// 	if os.IsPermission(ferr) { //erişim izniyle ilgili hata varsa
	// 		fmt.Printf("Yazma izni reddedildi. (%v)\n", ferr.Error())
	// 		//okuma izniyle açsak okuma iznini reddetmiş olacaktı!
	// 		//logging.Error("Yazma izni reddedildi. (%v)\n", ferr.Error())
	// 		//logging paketi dışardan yüklendi. Her yerde çalışmayabilir!
	// 	} else if os.IsExist(ferr) { //dosya zaten varsa (burada gereksiz aslında)
	// 		fmt.Println("Dosya zaten var!")
	// 	} else if os.IsNotExist(ferr) { // dosya bulunamadıysa
	// 		fmt.Println("Dosya bulunamadı!")
	// 	} else { //bunların dışında başka bir hataysa
	// 		fmt.Println("Bilinmeyen Hata!\n", ferr.Error())
	// 	}
	// } else { //hata yoksa
	// 	fmt.Println(f.Name() + " dosyasına yazma yapılabilir.")
	// 	bilgi, _ := os.Stat(f.Name())
	// 	fmt.Println("Bilgiler:\n", bilgi)
	// }

	// // DOSYA KOPYALAMAK
	// originallink := "orjinal.txt"
	// copiedlink := "kopya_orj.txt"
	// eski, ferr := os.Open(originallink)
	// defer eski.Close()
	// if ferr != nil {
	// 	fmt.Printf(ferr.Error())
	// }
	// yeni, yerr := os.Create(copiedlink)
	// if yerr != nil {
	// 	fmt.Printf(yerr.Error())
	// }
	// defer yeni.Close()

	// //kopyalama işlemi buradan sonra başlıyor aslında!
	// bytesWritten, err := io.Copy(yeni, eski)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }
	// log.Printf("%d bytes copied.", bytesWritten)
	// // aşağısı yapılmasa da kopyalanabilir ama yapılmasında fayda var!
	// //dosya içeriğini işleme // belleği diske boşaltma
	// //e := yeni.Sync()
	// // if e != nil {
	// // 	fmt.Printf(e.Error())
	// // }

	// // DOSYAYA VERİ YAZMAK
	// f, ferr := os.OpenFile("deneme.txt", os.O_WRONLY|os.O_CREATE, 0666)
	// if ferr != nil {
	// 	log.Fatal(ferr.Error())
	// }
	// defer f.Close()
	// //byteslice şeklinde yazma
	// bytesSlice := []byte("Bu bir deneme amaçlı byteslice yazma işlemidir.\n")
	// bw, e := f.Write(bytesSlice)
	// if e != nil {
	// 	log.Fatal(e.Error())
	// }
	// //string olarak yazma
	// bws, er := f.WriteString("bu da string olarak yazma işleminin örneği\n")
	// if er != nil {
	// 	log.Fatal(er.Error())
	// }
	// //yazılan toplam byte türünden veri
	// fmt.Println(bw+bws, " byte dosyaya yazıldı!")

	// // DOSYAYA HIZLI VERİ YAZMAK (ERROR ALMAZSAK SADECE TEK SATIR!)
	// _ = ioutil.WriteFile("deneme.txt", []byte("Bu da hızlı yazma"), 0666)

	// GEÇİCİ KLASÖR VE DOSYALARLA ÇALIŞMAK
	//geçici dizin oluşturma
	tempDir, err := ioutil.TempDir("", "geciciKlasör")
	hataYakala(err) //sürekli error ile uğğraşmamak için yazılan fonksiyon
	fmt.Println("Geçiçi Klasör Oluşturuldu: ", tempDir)

	//geçici dosya oluşturma
	tempFile, err := ioutil.TempFile(tempDir, "geciciDosya.txt")
	hataYakala(err)
	fmt.Println("Geçiçi Dosya Oluşturuldu: ", tempFile.Name())
	//geçici dosyayı kapatma
	err = tempFile.Close()
	hataYakala(err)

	//oluşturulan geçici dosya ve klasörlerin silinmesi
	defer os.Remove(tempFile.Name())
	defer os.Remove(tempDir)

}

func dosyaOlustur(dosyaAdi string) (newFile *os.File, err error) {
	newFile, err = os.Create(dosyaAdi)
	return
}

//hataYakala her seferinde (if err != nil {...}) işlemini yapmamızı engeller!
func hataYakala(e error) {
	if err != nil {
		fmt.Errorf(e.Error())
	}
}
