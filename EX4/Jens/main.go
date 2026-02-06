package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	case "darwin":
		wd, _ := os.Getwd()
		cmd = exec.Command("osascript", "-e", `tell app "Terminal" to do script "cd `+wd+`; go run main.go"`)
		fmt.Printf("Running command on MAC...\n")
	case "linux":
		cmd = exec.Command("gnome-terminal", "--", "go", "run", "main.go")
		fmt.Printf("Running command on LINUX...\n")
	default:
		fmt.Printf("Unsupported OS: %s\n", runtime.GOOS)
		return
	}

	cmd.Run()

	i := 0
	for {
		// Open the file for appending (create if not exists)
		// f, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) // APPEND MODE
		f, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644) // TRUNCATE MODE
		if err != nil {
			log.Fatal(err)
		}

		// Write the counter as a string
		_, err = f.WriteString(strconv.Itoa(i) + "\n")
		if err != nil {
			f.Close()
			log.Fatal(err)
		}
		f.Close()

		i++
		time.Sleep(1 * time.Second)

		content, err := os.ReadFile("output.txt")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(content))
	}
}

func primary() {
	fmt.Printf("Sleeping...\n")
	time.Sleep(5 * time.Second)
}

func backup() {
	fmt.Printf("Sleeping...\n")
	time.Sleep(5 * time.Second)
}
