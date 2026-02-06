package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		_, err2 := os.Create("output.txt")
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	defer file.Close()
	i := 0
	for {
    file, err := os.Open("output.txt")
		_, err2 := file.WriteString(string(i) + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
		i++
		time.Sleep(2 * time.Second)
		content, err := os.ReadFile(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(content))
	}

}
