package main

import (
	"fmt"
	"strings"
	"time"
)

// getFileName extracts the file name from URL
func getFileName(url string) string {
	segments := strings.Split(url, "/")
	return segments[len(segments)-1]
}

func formatDuration(d time.Duration) string {
	msec := d.Milliseconds()
	return fmt.Sprintf("%d ms", msec)
}
