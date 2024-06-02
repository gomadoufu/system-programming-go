package process

import (
	"fmt"
	"os"
)

func pid() {
	fmt.Printf("プロセスID: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())
}
