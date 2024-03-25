package reader

import (
	"fmt"
	"strings"
)

var fsource = "123 1.234 1.0e4 test"

func fscan() {
	reader := strings.NewReader(fsource)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
