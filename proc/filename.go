package process

import (
	"fmt"
	"os"
)

func fileName() {
	path, _ := os.Executable()
	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	fmt.Printf("実行ファイルのパス: %s\n", path)
}
