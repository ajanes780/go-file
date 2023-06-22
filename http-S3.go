package main

import (
	"bufio"
	"log"
	"os"
)

func fetchImages() {
	file, err := os.Open("parsed.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing the file: %s", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()

		err := fetch(url)
		if err != nil {
			log.Printf("Request to %s failed: %s", url, err)
			continue
		}

	}
}
