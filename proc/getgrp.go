package process

import (
	"fmt"
	"os"
	"syscall"
)

func grp() {
	sid, _ := syscall.Getsid(os.Getegid())
	fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n", os.Getegid(), sid)
}
