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

	/*
		switch platform {
			case "ghes":
				break
			case "ghec":
				break
			case "dotcom":
				break
			default:
				return fmt.Errorf("invalid platform provided. platform must be one of dotcom, ghes, ghec. given platform: %s", platform)
		}*/

	fileName := ""

	if platform == "ghes" {
		fileName = "schemas/" + platform + "-" + version + ".json"
	} else if platform == "ghec" {
		fileName = "schemas/" + platform + ".json"
	} else {
		fileName = "schemas/api.github.com.json"
	}

	out, err := os.Create(fileName)
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
