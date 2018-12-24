package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("echo", "hello")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Err when executing uname command", err)
		return
	}

	fmt.Println("I am running on", out.String())
}
