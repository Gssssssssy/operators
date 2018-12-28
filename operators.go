package operators

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func ExecShellCmd(s string) {
	cmd := exec.Command("/usr/bin/zsh", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

// 这里可使用goroutine
func CheckArrayFile(s string, array []interface{}) {
	for i, fn := range array {
		log.Println(i, fn)
	}
}
