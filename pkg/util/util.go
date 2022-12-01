package util

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

// BuildPath - Create directory if not exists
func BuildPath(outPath string) (string, error) {
	// log.Printf("Building... %s \n", outPath)
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		err := os.Mkdir(outPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Export output directory: %s - %v", outPath, err)
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
