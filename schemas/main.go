package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("execution finished successfully")
}

func realMain() error {
	fmt.Printf("downloading latest schema version...\n")
	out, err := os.Create("schemas/downloaded.json")
	if err != nil {
		return err
	}
	defer out.Close()

	url := "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/api.github.com/api.github.com.json"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
