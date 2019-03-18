package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "log")
	logBytes, err := cmd.Output()
	if err != nil {
		fmt.Printf("Git Error: %s\n", err)
		os.Exit(1)
	}
	log := string(logBytes)
	fmt.Println(log)
}
