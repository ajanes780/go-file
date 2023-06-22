package main

func main() {
	file := readFile("newcurl.log")
	createParsedLogs(file, "parsed.txt")
	//createCacheCalls()

	//fmt.Println("starting s3 calls....")
	//
	//start := time.Now()
	//fetchImages()
	//duration := time.Since(start)
	//fmt.Println("Fetch images from s3 took: ", formatDuration(duration))
	//
	//fmt.Println("starting cache calls....")
	//start = time.Now()
	//fetchImagesFromCache()
	//duration = time.Since(start)
	//fmt.Println("Fetch Images from cache took: ", formatDuration(duration))

}
