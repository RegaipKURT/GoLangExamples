package main

/*
Go'dilindeki map yapısı başka dillerdeki sözlük yapısına benzerdir.
Key-value şeklinde çalışır.
*/
import (
	"fmt"
)

//Konum Altitude ve Longtitude değerlerini float64 olarak tutar.
type Konum struct {
	ALT, LONG float64
}

func main() {
	//oluşturulurken key ve value tipi belirtilir.
	myMap := make(map[string]string /*sırasıyla eleman ve kapasite de yazılabilir*/)
	myMap["Ali"] = "İyi"
	myMap["Veli"] = "Kötü"
	myMap["Deli"] = "Mükemmel"
	fmt.Println(myMap)

	//kendi oluşturduğumuz tipleri de kullanabiliriz.
	var harita = make(map[string]Konum)
	harita["İstanbul"] = Konum{41.015137, 28.979530}
	harita["Diyarbakır"] = Konum{37.910000, 40.240002}
	harita["Adana"] = Konum{37.000000, 35.321335}
	harita["Giresun"] = Konum{40.917534, 38.392654}
	harita["Ankara"] = Konum{39.920944, 32.854547}
	fmt.Println(harita["Giresun"])
	for i, j := range harita {
		fmt.Println("Şehir:", i, "   ===>   Altitude:", j.ALT, "Longtitude:", j.LONG)
	}

	//silme işlemi delete komutuyla yapılabilir veya nil değer atanabilir.
	delete(harita, "İstanbul")
	fmt.Println(harita)
	harita["Tekirdağ"] = Konum{40.977779, 27.515278}
	fmt.Println(harita)

	//haritanın keyleri ile bir slice oluşturalım
	sehirler := make([]string, len(harita))
	i := 0
	for k := range harita {
		sehirler[i] = k
		i++
	}
	fmt.Println(sehirler)
}
