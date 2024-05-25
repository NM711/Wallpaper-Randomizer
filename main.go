package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

func main() {
	wallpaperPath := "/home/user/Pictures/Wallpapers"
	wallpapers, err := os.ReadDir(wallpaperPath)

	if err != nil {
		log.Println(err)
		log.Printf(`Error: Failed to open/read directory at "%s"!`, wallpaperPath)
		os.Exit(1)
	}

	length := len(wallpapers)

	index := rand.Intn(length)

	file := wallpapers[index].Name()

	log.Printf(`Wallpaper Selected: "%s"`, file)

	fullFilePath := fmt.Sprintf("%s/%s", wallpaperPath, file)

	cmd := exec.Command("swww", "img", fullFilePath)

	_, outErr := cmd.Output()

	if outErr != nil {
		log.Println("Error: Something went wrong after command execution!")
		os.Exit(1)
	}
}
