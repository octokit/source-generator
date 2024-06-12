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

var platformTypeMap = map[string]platformType{
	"dotcom": dotcom,
	"ghes":   ghes,
	"ghec":   ghec,
}

func (p platformType) String() string {
	return [...]string{"dotcom", "ghes", "ghec"}[p]
}

func isValidPlatform(value string) (platformType, bool) {
	platform, exists := platformTypeMap[value]
	return platform, exists
}

// Docs [here](https://pkg.go.dev/flag) recommends this approach.
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

	if _, valid := isValidPlatform(platform); valid == false {
		return fmt.Errorf("invalid platform provided. platform must be one of dotcom, ghes, ghec. given platform: %s", platform)
	}

	if platform == platformType.String(ghes) && version == "" {
		return fmt.Errorf("version is required for GHES platform")
	}

	if (platform == platformType.String(dotcom) || platform == platformType.String(ghec)) && version != "" {
		return fmt.Errorf("version may not be specified for given platform: %s", platform)
	}

	fileExt := ".json"
	fileName := "api.github.com"

	if platform == platformType.String(ghes) {
		fileName = platform + "-" + version
	} else if platform == platformType.String(ghec) {
		fileName = platform
	}

	downloadedFile := "schemas/" + fileName + fileExt

	out, err := os.Create(downloadedFile)
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
		// Clean up the created file
		err := os.Remove(downloadedFile)
		if err != nil {
			return fmt.Errorf("Error deleting file: %s. Error: %s", downloadedFile, err)
		}
		return fmt.Errorf("Received 404 Not Found for url: %s", url)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
