package util

import (
	"bytes"
	"encoding/json"
	"mcolomerc/cc-tools/pkg/log"
	"os"
	"time"
)

// BuildPath - Create directory if not exists
func BuildPath(outPath string) (string, error) {
	// log.Printf("Building... %s \n", outPath)
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		err := os.Mkdir(outPath, os.ModePerm)
		if err != nil {
			log.Error("Export output directory: %s - %v", outPath, err)
			return "", err
		}
	}
	return outPath, nil
}

// Indent JSON
func IndentJson(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "\t"); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

// Time spent
func Timer(name string) func() {
	start := time.Now()
	return func() {
		log.Info("%s took %v\n", name, time.Since(start))
	}
}
