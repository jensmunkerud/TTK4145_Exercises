package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	filePath := "counting.txt"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, err2 := os.Create(filePath)
		if err2 != nil {
			fmt.Println("Error creating file:", err2)
			return
		}
	}
	file, err := os.Open("counting.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	isPrimary := flag.Bool("primary", true, "Run as primary")
	flag.Parse()
	if *isPrimary {
		primary(file)
	} else {
		backup(file)
	}
	os.Exit(0)
}

func primary(file *os.File) {
	defer file.Close()
	fmt.Printf("Running primary process...\n")
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", "go run main.go -primary=false")
		fmt.Printf("Running command on WINDOWS...\n")
	case "darwin":
		wd, _ := os.Getwd()
		cmd = exec.Command("osascript", "-e", `tell app "Terminal" to do script "cd `+wd+`; go run main.go -primary=false"`)
		fmt.Printf("Running command on MAC...\n")
	case "linux":
		cmd = exec.Command("gnome-terminal", "--", "go", "run", "main.go", "-primary=false")
		fmt.Printf("Running command on LINUX...\n")
	default:
	}

	if cmd != nil {
		cmd.Start()
	}

	content, _ := os.ReadFile("counting.txt")
	var count int
	if len(content) != 0 {
		num, err := strconv.Atoi(strings.TrimSpace(string(content)))
		if err == nil {
			count = num + 1
		}
	} else {
		count = 1
	}
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Count:", count)
		count++
		err := os.WriteFile("counting.txt", []byte(strconv.Itoa(count)), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func backup(file *os.File) {
	defer file.Close()
	fmt.Printf("Running backup..\n")

	for {
		initialContent, _ := os.ReadFile("counting.txt")
		time.Sleep(2 * time.Second)
		currentContent, _ := os.ReadFile("counting.txt")

		if string(currentContent) == string(initialContent) {
			fmt.Println("No file change detected, running primary...")
			primary(file)
			break
		}
	}
}
