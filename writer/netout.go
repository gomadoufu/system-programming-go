package writer

import (
	"io"
	"net"
	"os"
)

func netout() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
	file, err := os.Create("response.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(file, conn)
}
