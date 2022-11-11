package util

import (
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
