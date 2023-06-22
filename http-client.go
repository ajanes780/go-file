package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func fetch(url string) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	//url := "https://renoworks-svn-product.s3.us-west-2.amazonaws.com/waypoint/60743/Interior/Backsplash/Generic/~Generic_Tile_Backsplash/1x3inch_stone_himalayan_0014_Layer%2016.jpg" // replace with your URL
	resp, err := client.Get(url)

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalf("Error closing the response body: %s", err)
		}
	}()

	// Discard the response
	_, err = io.Copy(io.Discard, resp.Body)

	if err != nil {
		log.Fatalf("Error closing the response body: %s", err)
	}
	return err
}
