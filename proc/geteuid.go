package process

import (
	"fmt"
	"os"
)

func euid() {
	fmt.Printf("実効ユーザID: %d\n", os.Geteuid())
	fmt.Printf("実効グループID: %d\n", os.Getegid())
}
