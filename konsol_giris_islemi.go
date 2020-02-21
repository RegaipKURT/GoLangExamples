package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Konsoldan girilen her veri bize string tipinde gelir.
	reader := bufio.NewReader(os.Stdin) //konsoldan girilen veriyi alan okuyucu
	fmt.Print("Lütfen bir yazı giriniz: ")
	str, _ := reader.ReadString('\n')
	fmt.Println("Girilen yazi : " + str)
	// belirttiğiimiz karakter gelene kadar okur. (Yeni satıra geçene kadar bu durumda)
	//reader içinde farklı çeşitlerde kullanabileceğimiz metotlar var.

	fmt.Print("Bir sayi giriniz: ")
	sayi, _ := reader.ReadString('\n')
	sayiFloat, _ := strconv.ParseFloat(strings.TrimSpace(sayi), 64)
	fmt.Printf("Girilen sayi float olarak: %.5v \n", sayiFloat)

}
