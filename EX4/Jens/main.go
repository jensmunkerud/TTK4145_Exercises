package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("Sleeping...\n")
	time.Sleep(5 * time.Second)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", "go run main.go")
		fmt.Printf("Running command on WINDOWS...\n")
	case "linux", "darwin":
		cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", "go run main.go")
		fmt.Printf("Running command on MAC...\n")
	}

	cmd.Run()
}

func primary() {
	fmt.Printf("Sleeping...\n")
	time.Sleep(5 * time.Second)
}

func backup() {
	fmt.Printf("Sleeping...\n")
	time.Sleep(5 * time.Second)
}
