package main

import "fmt"

func main() {
	dizi := [10]int{1, 2, 3, 4, 5}
	fmt.Println(dizi)

	var yeniDizi []string = []string{"ali", "veli", "deli"}
	fmt.Println(yeniDizi)

	var baskaDizi [10]float64
	baskaDizi[0], baskaDizi[1] = 1.5, 3.5
	fmt.Println(baskaDizi)

	//2 dönüşlü for dizi kullanımı
	for indis, eleman := range yeniDizi { //ilk indis, sonra eleman döndürür!
		fmt.Println(indis, eleman)
	}
	//tek dönüşlü for dizi kullanımı
	for eleman := range baskaDizi { // sadece indisleri döndürür.
		fmt.Println(eleman)
	}
	//sadece elemanları almak istiyorsak:
	for _, eleman := range yeniDizi {
		fmt.Println(eleman)
	}

	//uzunluk yazmadan ... koyarak da otomatik ayarlanabilir.
	myArray := [...]int{2, 34, 12, 2, 3, 1, 43, 65, 322, 432, 4254, 4}
	fmt.Println(myArray)
	fmt.Println("Uzunluk: ", len(myArray))

	var cars []string = []string{"BMW", "TESLA", "ŞAHİN", "ANADOL", "JAGUAR"}
	i := 0
	for i < len(cars) {
		fmt.Println(cars[i])
		i++
	}

	for ind, el := range cars {
		fmt.Println("İndex:", ind, " ==>  Eleman:", el)
	}

	//silice oluşturma (slice pointer mantığıyla çalışır. dizi değişirse slice da değişir.)
	var dilim = cars
	var dilimP = dilim[1:4]
	fmt.Println(dilimP)

	//array olmadan da slice tanımlanabilir
	slice := []int{}
	fmt.Println(len(slice), cap(slice)) //ilk kapasite ve uzunluk değerleri
	for i := 0; i < 10; i++ {
		slice = append(slice, i+i*12) //append slicelara eleman eklemeye yarar.
		//eleman ekledikçe kapasite ve uzunluk değişir.
		//eklenen eleman kapasiteyi aşarsa, kapasite otomatik olarak 2 katına çıkarılır.
	}
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice)) //son kapasite ve uzunluk değerleri

}
