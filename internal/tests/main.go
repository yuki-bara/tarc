package main

import (
	"os"

	"github.com/yuki-bara/tarc"
)

func main() {
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
