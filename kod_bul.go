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
	args := ""
	for _, w := range os.Args[2:] {
		args += w + "+"
	}
	var url string
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
