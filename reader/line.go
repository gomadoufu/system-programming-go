package reader

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var lsource = `1行目
2行目
3行目`

func readLine() {
	reader := bufio.NewReader(strings.NewReader(lsource))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
