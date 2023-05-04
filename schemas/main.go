package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var useSchemaNext bool

// Normally I hate using init() but the docs [here](https://pkg.go.dev/flag)
// recommend it so this usage is safe.
func init() {
	flag.BoolVar(&useSchemaNext, "schema-next", true, "Set to false using --schema-next=false to disable using the descriptions-next directory for schema downloads")
}

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
	log.Println("execution finished successfully")
}

func realMain() error {
	flag.Parse()

	out, err := os.Create("schemas/downloaded.json")
	if err != nil {
		return err
	}
	defer out.Close()

	logMsg := "Downloading latest schema from descriptions directory"
	url := "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/api.github.com/api.github.com.json"
	if useSchemaNext {
		logMsg = "Downloading latest schema from descriptions-next directory"
		url = "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions-next/api.github.com/api.github.com.json"
	}

	log.Printf(logMsg)
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
