package writer

import (
	"bytes"
	"fmt"
)

func buffer() {
	Stringio()
	StringBuilder()
}

func Stringio() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

func StringBuilder() {
	var builder bytes.Buffer
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}
