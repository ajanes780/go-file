package main

import (
	"bufio"
	"log"
	"os"
)

func fetchImagesFromCache() {
	file, err := os.Open("cacheTestCalls.txt")
	defer func() {
		err = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		err = fetch(url)
	}

	if err != nil {
		log.Fatalf("failed: %s", err)
	}

}
