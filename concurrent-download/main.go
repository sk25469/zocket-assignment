package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func downloadSingleFile(url string, dest string, index int) ([]byte, error) {
	start := time.Now()
	fileName := "file" + strconv.Itoa(index)

	log.Printf("Downloading file %s from %s\n", fileName, url)

	var path bytes.Buffer
	path.WriteString(dest)
	path.WriteString("/")
	path.WriteString(fileName)

	outputPath, err := os.Create(path.String())

	if err != nil {
		fmt.Println(path.String())
		panic(err)
	}

	defer outputPath.Close()

	if err != nil {
		panic(err)
	}

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(outputPath, resp.Body)

	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	log.Printf("Single file download completed in %s", elapsed)

	return path.Bytes(), nil
}

func downloadMultipleFiles(urls []string) ([][]byte, error) {
	start := time.Now()
	doneChannel := make(chan []byte, len(urls))
	errorChannel := make(chan error, len(urls))
	for index, URL := range urls {
		go func(URL string, index int) {
			downloadedBytes, err := downloadSingleFile(URL, "./", index)
			if err != nil {
				errorChannel <- err
				doneChannel <- nil
				return
			}
			doneChannel <- downloadedBytes
			errorChannel <- nil
		}(URL, index)
	}
	elapsed := time.Since(start)
	log.Printf("Download completed in %s", elapsed)
	bytesArray := make([][]byte, 0)
	var errorMessage string
	for i := 0; i < len(urls); i++ {
		bytesArray = append(bytesArray, <-doneChannel)
		if err := <-errorChannel; err != nil {
			errorMessage = errorMessage + " " + err.Error()
		}
	}
	var err error
	if errorMessage != "" {
		err = errors.New(errorMessage)
	}
	return bytesArray, err
}

func main() {
	urls := []string{"https://drive.google.com/file/d/13_isyESbDNAPDh-wyQKZrtknaQCaZaE8/view?usp=share_link", "https://drive.google.com/file/d/12rHpVxwYQ4CsCu9qAnh9GZ_qeCsqTcAd/view?usp=share_link", "https://drive.google.com/file/d/12oM_PJyUmTQwrwrJvOXr3oC7wbR1TX0K/view?usp=share_link", "https://drive.google.com/file/d/12glUTOsaRihbGlc9BHwwRZXvYIpJqqmL/view?usp=share_link", "https://drive.google.com/file/d/12aIHx1izWKg5eYQL6z32MIf6NcwGBDW_/view?usp=share_link", "https://drive.google.com/file/d/12LEXa-kMGC1BJKFKp98u_1CRYN_ym6tS/view?usp=share_link", "https://drive.google.com/file/d/12PQtU9SpZTeJktUmyVwlRUtZVV2Vxsxr/view?usp=share_link", "https://drive.google.com/file/d/12RdhzfWubSJkNUh2aXAB9u5Dadzs0s38/view?usp=share_link", "https://drive.google.com/file/d/13-5foORv11gIFtDlRi4H0dxO93zcm8ue/view?usp=share_link", "https://drive.google.com/file/d/12z0Im57ng8uAuJRx4lBWkZrGgT2XU02W/view?usp=share_link", "https://drive.google.com/file/d/12K3OUq1HG9o8DRv6dMK1tOEivqqCXYdS/view?usp=share_link", "https://drive.google.com/file/d/12HNfvtvzTlwOJ9I_xfn2X2-xlojyXBYg/view?usp=share_link"}
	_, err := downloadMultipleFiles(urls)
	if err != nil {
		log.Printf("Something went wrong while downloading multiple files: %v", err.Error())
		return
	}
}
