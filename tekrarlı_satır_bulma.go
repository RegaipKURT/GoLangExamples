package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// BİR VEYA DAHA FAZLA DOSYA İÇİNDEN TEKRAR EDEN SATIRLARI BULAN PRGORAM!
	if len(os.Args[1:]) < 1 { // parametre girilmediyse mesaj gösterelim
		fmt.Println("Parametre Girilmedi.\nKullanım: Program [dosya isimleri]")
	}
	for _, file := range os.Args[1:] { //aldığı parametreleri dosya adlarının listesi olarak tuttuk.
		counts := make(map[string]int) //dosya satırlarını ve aynı satır geçme sayısını tutacağımız sözlük
		fmt.Println("\n******************************\n", "Taranan dosya: ", file)
		// yukarıdaki satırda dosyanın adını yazdırıp işlemlere başlıyoruz.

		f, err := os.Open(file) //dosyayı açıp f, hatayı err değerine atadık.
		if err != nil {         //hata nil değilse yani dosyayı açarken hata aldıysak.
			fmt.Fprintf(os.Stderr, "Hata: %v\n", err)
			continue
			//hatayı yazdırıp for döngüsünün başına döndük.
		} else { //hata olmadan açıldıysa
			countLines(f, counts) // contlines fonklsiyonuna dosyayı ve sözlüğü gönder.
			f.Close()             // dosyayı kapat.
		}

		for line, n := range counts { //saır adını ve geçme sıklığını counts içinden oku
			if n > 1 { // geçme sıklığı 1'den büyükse
				fmt.Printf("%d\t%s\n", n, line) // geçme sıklığı ve satırın kendisini yazdır.
			}
		}
		fmt.Println("******************************") // güzel görünsün diye ayraç yazdırıyoruz.
	}

}

func countLines(f *os.File, counts map[string]int) { //sözlük ve dosya alıyor içine
	input := bufio.NewScanner(f) // bir dosya alıp satırlarını döndüren obje
	for input.Scan() {           // taramayı başlat
		counts[input.Text()]++ //satıra karşılık gelen değeri 1 artır.
		// satır daha önce yoksa otomatik ekler.
	}
	if input.Err() != nil {
		fmt.Fprintf(os.Stderr, "Hata: %v\n", input.Err())
	}
}
