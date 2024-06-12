package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var useSchemaNext bool
var platform string
var version string

type platformType int

const (
	dotcom platformType = iota
	ghes
	ghec
)

var allowedPlatforms = []string{"dotcom", "ghes", "ghec"}

// Normally I hate using init() but the docs [here](https://pkg.go.dev/flag)
// recommend it so this usage is safe.
func init() {
	flag.BoolVar(&useSchemaNext, "schema-next", false, "Set to true using --schema-next=true to use the descriptions-next directory for schema downloads")
	flag.StringVar(&platform, "platform", "dotcom", "Set the platform to download the schema for. Options are dotcom, ghes, ghec")
	flag.StringVar(&version, "version", "", "Set the version of the schema to download. Required for GHES")
}

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
	log.Println("execution finished successfully")
}

func realMain() error {
	flag.Parse()

	foundPlatform := false
	for _, p := range allowedPlatforms {
		if p == platform {
			foundPlatform = true
			break
		}
	}

	if foundPlatform == false {
		return fmt.Errorf("invalid platform provided. platform must be one of dotcom, ghes, ghec. given platform: %s", platform)
	}

	if platform == "ghes" && version == "" {
		return fmt.Errorf("version is required for GHES platform")
	}

	if (platform == "dotcom" || platform == "ghec") && version != "" {
		return fmt.Errorf("version may not be specified for given platform: %s", platform)
	}

	fileExt := ".json"
	fileName := "api.github.com"

	if platform == "ghes" {
		fileName = platform + "-" + version
	} else if platform == "ghec" {
		fileName = platform
	}

	out, err := os.Create("schemas/" + fileName + fileExt)
	if err != nil {
		return err
	}

	defer out.Close()

	logMsg := "Downloading latest schema from descriptions directory"
	url := "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/" + fileName + "/" + fileName + fileExt
	if useSchemaNext {
		logMsg = "Downloading latest schema from descriptions-next directory"
		url = "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions-next/" + fileName + "/" + fileName + fileExt
	}

	log.Printf(logMsg)
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("Received 404 Not Found for url: %s", url)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
