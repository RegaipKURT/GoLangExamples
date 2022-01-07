package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var HOME = os.Getenv("HOME")

func firstRun() {

	if _, err := os.Stat(HOME + "/.kodbulConfig"); os.IsNotExist(err) {
		fmt.Println("Configuration File Not Found!")

		file, err := os.Create(HOME + "/.kodbulConfig")
		defer file.Close()
		if err != nil {
			log.Fatal("Error when creating configuration file: ", err.Error())
		}

		fmt.Println("Required updates in progress...")
		cmd := exec.Command("sudo", "apt", "install", "tldr", "curl")
		res, err := cmd.Output()
		x := string(res)
		fmt.Println(x)
		cmd2 := exec.Command("tldr", "--update")
		x2, err := cmd2.Output()
		if err != nil {
			fmt.Println("An error have raised: ", err.Error())
		} else {
			x3 := string(x2)
			fmt.Println(x3)
			fmt.Println("Update Successful!")
		}
		_, err = file.WriteString("1\n")
		if err != nil {
			log.Fatal("Error when writing configuration file: ", err.Error())
		}
	} else if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	firstRun()
	if len(os.Args) <= 1 {
		cmd := exec.Command("curl", "cht.sh")
		res, err := cmd.Output()
		if err != nil {
			fmt.Println("cheat.sh error: ", err.Error())
		}
		fmt.Println(string(res))

		cmd2 := exec.Command("tldr", "-h")
		res, err = cmd2.Output()
		if err != nil {
			fmt.Println("tldr error:", err.Error())
		}
		fmt.Println(string(res))

		fmt.Println("BU PROGRAM YUKARIDA GÖRÜLEN cht.sh ve tldr PROGRAMLARI ÜZERİNDEN ARAMA YAPAR.\ntldr komutunun yüklenmiş olması gerekir.")
		fmt.Println("En az bir programlama dili veya parametre belirtin!")
		fmt.Println("Kullanım (1): kodbul <programlama dili> <aranacak ifadeler>")
		fmt.Println("Kullanım (2): kodbul <bash girdisi>")

		os.Exit(0)
	}

	var args string
	var url string

	if len(os.Args) == 2 {
		if os.Args[1] != "-h" || os.Args[1] != "--help" {
			cmd := exec.Command("tldr", os.Args[1])
			res, err := cmd.Output()
			if err != nil {
				fmt.Println("tldr error: ", err.Error())
			}
			fmt.Println(string(res))
		} else {
			fmt.Println("BU PROGRAM cht.sh ve tldr PROGRAMLARI ÜZERİNDEN ARAMA YAPAR.\ntldr komutunun yüklenmiş olması gerekir.")
			fmt.Println("\nKullanım (1): kodbul <programlama dili> <aranacak ifadeler>")
			fmt.Println("Kullanım (2): kodbul <bash girdisi>")
			os.Exit(0)
		}
	}

	for _, w := range os.Args[2:] {
		args += w + " "
	}

	args, url = "", ""

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
