package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// read log file
// parse out the http request

func readFile(fileName string) *os.File {
	// open file
	file, err := os.Open(fileName)
	//defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open file: s%", err)

	}
	return file

}

func createParsedLogs(file *os.File, fileName string) {

	outfile, err := os.Create(fileName)

	if err != nil {
		log.Fatalln("Failed to create file s%", err)
	}
	re := regexp.MustCompile(`https.*$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		matches := re.FindString(line)
		parts := strings.Split(matches, " ")
		// If there are less than 2 parts, the string doesn't have a URL and status code
		if len(parts) < 2 {
			continue
		}
		//fmt.Println("part 1" + parts[0])
		//fmt.Println("part 2" + parts[1])
		//fmt.Println("part 3: " + parts[2])
		// The status code is the last part
		statusCode := parts[len(parts)-1]
		//fmt.Println(statusCode)
		if statusCode == "200" {
			// The URL is all the parts except the last one
			url := strings.Join(parts[:len(parts)-2], " ")

			if strings.HasPrefix(url, "https://projects-renoworks.s3.us-west-2.amazonaws.com") {
				_, err2 := fmt.Fprintln(outfile, url)
				if err2 != nil {
					return
				}
			}

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %s", err)
	}

}

func createCacheCalls() {
	// Take parse file and remove base url
	// add cache server address
	output, err := os.Create("cacheTestCalls.txt")
	file, err := os.Open("parsed.txt")

	if err != nil {
		log.Fatalln("Failed to creat files s%", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// replace the http url with the ec2 instance address
		// http://52.27.208.253:3000/

		// the 1 means the first occurance if you wanted all you can do  -1
		newUrl := strings.Replace(line, "https://renoworks-svn-product.s3.us-west-2.amazonaws.com/", "http://52.27.208.253:3000/", 1)
		_, err2 := fmt.Fprintln(output, newUrl)

		if err2 != nil {
			return
		}
	}

}
