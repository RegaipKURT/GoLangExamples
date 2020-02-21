package main

import (
	"fmt"
	// "runtime"
)

/*
	KANALLARIN VE GORUTİNELERİN TEMELLERİ

	Gorutine'ler thread değildir. Gorutine ve Thread farkını anlamak için:
	https://www.geeksforgeeks.org/golang-goroutine-vs-thread/
	https://www.tutorialspoint.com/difference-between-goroutine-and-thread-in-golang

	Kanalları anlamak için: https://www.geeksforgeeks.org/channel-in-golang/

	Kanallar go routinelerin birbirleriyle iletişim kurma aracıdır. Bir kanalın belli bir
	tipi olur ve kanal kendi tipi dışında herhangi bir veri taşıyamaz. (ör: var chan int)
	Kanala değer atama işlemi <- işaretiyle yapılır ve bu işlem iki çeşit olabilir.

	Recieve Operation
	Eğer <- işareti kanalın sağında ise recieve (teslim alma) işlemi yapılıyordur.
	(Örn: channel<- value).

	Send Operation
	Eğer <- işlemi kanalın solunda ise kanal kendine gelen değeri gönderme işlemi yapıyordur.
	Bu durumda kendine bir değer gelene kadar bekler ve bir değer aldığında belirtilen
	değişkene o değeri atama işlemi yapılır. Eğer kanala değer gönderilmemişse deadlock
	hatası ile program sonlandırılır.
	(Örn: value:= <-channel)

	Kanalın Açık veya Kapalı olması durumu
	Kanaldan değer alınırken yani send operasyonu sırasında, ayı zamanda kanalın açık mı
	yoksa kapalı mı olduğu bilgisi de bool olarak döner. Eğer istenirse bu bool değer de
	alınıp kullanılabilir. Atama işlemi sırasında karşılanmazsa bu değer döndürülmez.
	Açık olma durumu: true  --- Kapalı olma durumu: false
	Kanalın açık veya kapalı olması o kanal üzerinden bir veri gönderilip gönderilemeyeceğini
	berlirler. Kapalı ise o kanala bir değer atanamaz. Kanal kapandığı anda değeri kendi
	tipindeki default değer olarak belirlenir. Bu değer ise kanaldan okunabilir.
	(Örneğin: string tipinde kanal için kapalı kanal değeri: "", int kanal için: 0 )

*/

// square aldığı kanal değerine her seferinde değerin karesini gönderir.
func square(c chan int) {
	for i := 0; i < 10; i++	{
		c <- i*i
	}
	close(c) // for döngüsü bitince kanalı kapatacak.
	deger := <-c
	fmt.Println("Kapalı kanal değeri: ", deger)
}

// okuyucu kanaldan gelen değerleri kanal açık olduğu sürece okur.
func okuyucu(c chan int) {
	for {
		val, ok := <-c //channel'dan gelen değerleri bekliyoruz.
		//kanal kendine gelen değeri ve açık olup olmadığı (bool) bilgisini döndürür.
		if ok == false { //eğer kanal kapalıysa döngüyü kırıyoruz.
			fmt.Println(val, ok, "<--- loop broke.")
			break
		}
		go fmt.Println(val, ok) // kanal açıksa, değeri ve açık olma durumunu yazıyoruz.
		//ama go keywordü ile ayrı bir go routine oluşturduğumun için sırası karışık olacak.
	}
}

func main() {
	fmt.Println("Main Başladı!")

	//farklı çekirdekler kullanarak paralel programlama örneği uygulayabiliriz.
	// böylece her goroutine farklı bir işlemciye veya çekirdeğe gönderirlir.
	// runtime.GOMAXPROCS(8) //bir goroutine tarafından kullanılacak max. işlemci veya çekirdek sayısı.

	c := make(chan int) //kanal oluşturma

	go square(c)  // kare alan fonksiyonu channel ile birlikte goroutine olarak çalıştırdık.
	go okuyucu(c) // okuyucu square fonksiyonundan gelen değerleri kanal kapanana kadar okur.
	//okuyucudan yazdırılan değerler de birer goroutine olduğundan yazdırma sırası karışıktır.
	// fmt.Println() fonksiyonlarından hangisi önce işini bitirdiyse o yazdırılır.

	/*okuyucu donksiyonunun başına go koyarak çalıştırır ve aşağıdaki for döngüsünü kapatmazsak
	karışık olarak çalışırlar. Çünkü kod okuyucudan sonra akmaya devam eder ve hangisi
	kanaldan gelen değeri önce okursa o bastıracaktır. Dolayısıyla okuyucudan yazılanlarda
	true veya false ile beraber değer yazılırken, aşağıdaki for döngüsünden
	yazılanlarda sadece kanalın integer değeri ekranda görünecektir.
	Hatta okuyucunun aldığı değerlerden bazıları görünmeyebilir. Çünkü okuyucu içindeki her
	fmt.Println() fonksiyonu da ayrı bir goroutine'dir ve bu goroutine'ler bitmeden main sonlanabilir.
	*/

	// BREAK KEYWORD'U KULLANAMDAN DA KANALDAN GELEN DEĞERLER OKUNABİLİR.
	//Burayı kullanırken değerleri okuyabilmek için yukarıdaki okuyucu fonksiyon kapatılmalı
	// veya kodun devam etmesi için başına go keywordü konulmalıdır.
	//Aksi takdirde buraya gelindiğinde okuyucu fonksiyonu bütün değerleri zaten okuduğu için
	// (Çünkü yazdığımız kod akışı gereği kanal kapanana kadar okumaya devam edecektir.)
	// ve sonunda kanal kapalı olduğu için hiçbir değer okunmadan bu döngüden çıkılır.
	for deger := range c {
		// sonsuz döngü kullanarak tıpkı bir dizi gibi gelen değerleri bekleyebiliriz
		// kanal kapalı olduğunda bu sonsuz döngü döngü otomatik olarak sonlanır.
		fmt.Println(deger)
	}
	// bu döngüyü kullanırken eğer kanal hiç kapatılmazsa program deadlock vererek kapanır.

	fmt.Println("Main Sonlandı!")
	//main donksiyonu içinde kanallardan değer bekleme işlemleri bitince kod sonlanır.
}
