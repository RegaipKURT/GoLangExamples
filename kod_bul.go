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
	fmt.Println(len(os.Args), os.Args)
	args := ""
	for _, w := range os.Args[2:] {
		args += w + "+"
	}
	url := "cht.sh/" + os.Args[1] + "/" + args[0:len(args)-1]
	fmt.Println(url)

	cmd := exec.Command("curl", url)
	res, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
}
