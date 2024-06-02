package process

import (
	"fmt"
	"os"
)

func uid() {
	fmt.Printf("ユーザID: %d\n", os.Geteuid())
	fmt.Printf("グループID: %d\n", os.Getegid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID: %v\n", groups)
}
