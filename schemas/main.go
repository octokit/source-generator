package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var useSchemaNext bool
var ghes bool

// Normally I hate using init() but the docs [here](https://pkg.go.dev/flag)
// recommend it so this usage is safe.
func init() {
	flag.BoolVar(&useSchemaNext, "schema-next", false, "Set to true using --schema-next=true to use the descriptions-next directory for schema downloads")
	flag.BoolVar(&ghes, "ghes", false, "Set to true using --ghes=true to pull down the descriptions in ghes-3.13")
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
	} else if ghes {
		logMsg = "Downloading the GHES 3.12 schema"
		url = "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/ghes-3.12/ghes-3.12.json"
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
