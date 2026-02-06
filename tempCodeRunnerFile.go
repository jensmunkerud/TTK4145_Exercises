package main

import (
	"os/exec"
	"path/filepath"
)

func main() {
	wd := "/Users/jens/Desktop/TTK4145_Exercises/EX4"
	cmd := exec.Command("osascript", "-e",
		`tell app "Terminal" to do script "cd `+filepath.ToSlash(wd)+`; ./dummyProgram"`)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}