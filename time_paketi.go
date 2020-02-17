package main

import (
	"fmt"
	"time"
)

func main() {
	// //the number of seconds elapsed since January 1, 1970 UTC.
	// fmt.Println("Unix Time:", time.Now().Unix())
	// fmt.Println("Uygulama iki saniye bekleyecek.")
	// time.Sleep(2 * time.Second)
	// //tarih oluşturma
	// tarihimiz :=time.Date(2020, 2, 18, 20, 45, 56, 0, time.Now().UTC().Local().Location())
	// fmt.Println(tarihimiz)
	// fmt.Println("20 Saniye için saat gösterilecek!")
	// time.Sleep(1*time.Second)

	//TARİHİ KENDİ BELİRLEDİĞİMİZ FORMATTA GÖSTERME
	benimFormatım := "02 January, Wednesday, 2006" //BU ÖRNEĞE BAKARAK DÜZEN OLUŞTURACAK.
	//yukarıya kendi istediğimiz düzeni yazdık örnek olarak.  değiştrilebilir.
	fmt.Println(time.Now().Format(benimFormatım))

	//TARİHLERİ KARŞILAŞTIRMA
	xTime := time.Date(2011, 12, 23, 23, 44, 33, 34, time.Now().UTC().Local().Location())
	//tarihleri karşılaştırırken aynı formatta olmasına dikkat edin yoksa hata verebilir!
	if xTime.Before(time.Now()) {
		fmt.Println("Şimdiki tarih daha sonra!")
	} else if xTime.After(time.Now()) {
		fmt.Println(xTime.Format(benimFormatım), "gelecekteki bir tarih!")
	} else if xTime.Equal(time.Now()) {
		fmt.Println("Tarihler Aynı!")
	}

	//aynı anda iki tarih alırsak biri önce biri sonra olur mu? Deneyelim.
	n, x := time.Now(), time.Now()
	if x.Equal(n) {
		fmt.Println("Alınan tarihler AYNI!")
	} else {
		fmt.Printf("Alınan tarihler FARKLI: \nn:%v\nx:%v\n", n, x)
	}

	// //tarih ekleme
	// yarın := time.Now().AddDate(0,0,1)
	// fmt.Println("Yarının tarihi: ",yarın)
}

// //bu fonksiyon saati her saniye sürekli gösterir.
// func saat()  {
// 	for {
// 		cmd := exec.Command("clear")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 		fmt.Print(time.Now().Hour(),":",time.Now().Minute(),":",time.Now().Second(),"\n")
// 		time.Sleep(1*time.Second)
// 	}
// }
