package reader

import (
	"io"
	"os"
)

func filein() {
	file, err := os.Open("reader/filein.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
