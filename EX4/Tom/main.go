package main

import(
	"os/exec"
	"time"
	"fmt"
)

func main() {
	fmt.Printf("Sleeping...\n")
	time.Sleep(10*time.Second)
	cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "go run main.go")
	fmt.Printf("Running command...\n")
	cmd.Run()
}

func primary() {
	fmt.Printf("Sleeping...\n")
	time.Sleep(10*time.Second)
}

func backup	() {
	fmt.Printf("Sleeping...\n")
	time.Sleep(10*time.Second)
}