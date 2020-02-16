package main

import (
	"fmt"
)

/*
Defer deyimi içinde bulunduğu metodun en sonunda çalıştırılması gereken kodu belirtir.
Yani defer ile belirttiğiniz işlem, içinde bulunduğu metodun diğer tüm işlemlerinin
bitmesini bekler ve ardından belirtilen işi yapar. Dolayısıyla defer veritabanı ve dosya
işlemlerinde sıkça kullanılır. Örneğin dosyayı veya bağlantıyı açtıktan hemen sonra defer ile
kapatma işlemini yaparsanız, o fonksiyon içindeki işlemler bittikten sonra kapatma işlemini
gerçekleştirir. Böylece bağlantıların veya dosyaların kapatılması unutulmuş olmaz ve kaynak
yönetiminde etkinliği sağlamış olursunuz.
*/

//bir vertiabanı işlemini basitçe simüle ederek görelim!
var isConnected bool = false

func main() {
	defer fmt.Println("Her şeyden sonra çalışacak! Bütün işlemler tamamlandı!")
	// main fonksiyonu bütün işini bitirdikten sonra üstteki yazı ekrana yazılacak!
	fmt.Printf("Connection Open: %v\n", isConnected)
	databaseProcess()
	fmt.Printf("Connection Open: %v\n", isConnected)
}

func databaseProcess() {
	connect()
	defer disconnect() //herşey bittikten sonra kapat dedik.
	//databaseProcess() fonksiyonu bittiğinde d,sconnect çalışacak yani.
	fmt.Printf("Connection Open: %v\n", isConnected)
	//bağlantı hala açık çünkü defer ile disconnect yaptık!
	fmt.Println("İşlem Yapılıyor.")
}

func connect() {
	isConnected = true
	fmt.Println("Bağlantı Gerçekleşti...")
}

func disconnect() {
	isConnected = false
	fmt.Println("Bağlantı Sonlandırıldı!")
}
