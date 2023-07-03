package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	path = "https://dl.k8s.io/release/v1.27.3/bin/linux/amd64/kubectl"
	err := Downloader(path)
	if err != nil {
		panic(err)
	}
}

func Downloader(urlPath string) (err error) {

	fullURLFile := urlPath
	fileName := ""
	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		fmt.Println(" url.Parse failed in Downloader")
		return
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create in Downloader")
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		fmt.Println(" client.Get failed in downloader")
		return
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

	return nil
}
