package main

// Bu program cheat.sh kullanarak aranan bir kod örneğini bulmaya yarar.
// orjinal program: https://github.com/chubin/cheat.sh
import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) <= 1 {
		cmd := exec.Command("curl", "cht.sh")
		res, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(res))

		fmt.Println("BU PROGRAM YUKARIDA GÖRÜLEN cht.sh PROGRAMI ÜZERİNDEN ARAMA YAPAR.\n")
		fmt.Println("En az bir programlama dili veya parametre belirtin!")
		fmt.Println("Kullanım: kodbul <programlama dili> <aranacak ifadeler>")
		os.Exit(0)
	}

	var args string
	var url string

	for _, w := range os.Args[2:] {
		args += w + "+"
	}

	if len(args) > 0 {
		url = "cht.sh/" + os.Args[1] + "/" + args[0:len(args)-1]
	} else {
		url = "cht.sh/" + os.Args[1]
	}

	cmd := exec.Command("curl", url)
	res, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}
