package writer

import (
	"bufio"
	"os"
)

func bufferio() {

	// ほかの言語で「バッファ付き出力」と呼ばれている機能は、Go言語ではこの bufio.Writer が提供している
	// bytes.Bufferではないので注意！

	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
