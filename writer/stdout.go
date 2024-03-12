package writer

import "os"

func stdout() {
	os.Stdout.Write([]byte("os.Stdout exmple\n"))
}
