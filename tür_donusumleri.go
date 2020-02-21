package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Tür Dönüşümleri
	var myString string = "1"
	var myInt int = 5
	var myFloat float64 = math.Pi

	fmt.Printf("myString : %T, myInt : %T, myFloat : %T\n", myString, myInt, myFloat)
	//%T gönderilen veririn tipini alan placeholder gösterimidir.
	fmt.Printf("String değeri: %s, integer değeri: %d, float değeri: %.4v\n",
		myString, myInt, myFloat)
	//FLOAT %v ve TOPLAM GÖSTERİLECEK RAKAM SAYISIYLA YAZILIR.
	// %v herhangi bir value için kullanılabilir.

	//With Convertion Methods
	//STRING TO INTEGER
	num, _ := strconv.Atoi(myString)
	fmt.Println("Sayi değerinin 10 fazlası: ", num+10)
	//INTEGER TO STRING
	str := strconv.Itoa(myInt)
	fmt.Println("Sayi Değeri: " + str) // "+" sadece aynı tipteki verilerde kullanılabilir.

	//With Casting Methods
	var i int = 55
	var f float64 = float64(i)
	var u uint = uint(i)

	fmt.Printf("i : %T, f : %T, u : %T\n", i, f, u)

}
