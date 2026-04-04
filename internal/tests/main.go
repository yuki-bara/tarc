package main

import (
	"fmt"
	"os"

	"github.com/yuki-bara/tarc"
)

func main() {
	fmt.Print("RUN TEST")
	os.Mkdir("out", 0755)
	err := tarc.Compressfile("test", "files/test.tar", "*")
	if err != nil {
		os.Exit(1)
	}
	err = tarc.Extractfile("files/test.tar", "out", "*")
	if err != nil {
		os.Exit(2)
	}
}
