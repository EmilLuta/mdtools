package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "log")
	gitOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("Git Error: %s\n", err)
	}
	fmt.Println(string(gitOutput))
}
