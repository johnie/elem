package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/joho/godotenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}

	elemFilePath := filepath.Join(usr.HomeDir, ".elem")

	err = godotenv.Load(elemFilePath)
	if err != nil {
		fmt.Println("Error loading .elem file:", err)
		return
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: elem <on|off>")
		return
	}

	action := os.Args[1]
	if action != "on" && action != "off" {
		fmt.Println("Invalid argument. Use 'on' or 'off'")
		return
	}

	elemURL := os.Getenv("ELEMENT_URL")
	if elemURL == "" {
		fmt.Println("Please set the ELEMENT_URL environment variable.")
		return
	}

	url := fmt.Sprintf(elemURL, action)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	capitalizedAction := cases.Title(language.English).String(action)
	fmt.Println("Turning element:", capitalizedAction)
}
