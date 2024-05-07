package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func convertHLSFileToMP4(hlsURL string, outputName string, index int, wg *sync.WaitGroup) {
	defer wg.Done()

	outputFile := fmt.Sprintf("%s_%d.mp4", outputName, index)

	cmd := exec.Command("ffmpeg", "-i", hlsURL, "-c", "copy", outputFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Conversion for %s failed: %v\n", hlsURL, err)
		return
	}

	fmt.Printf("Conversion for %s successful!\n", hlsURL)
}

func main() {
	files, err := os.ReadDir("./vids/ss2/raw")

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var hlsURLs []string

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".html" {
			content := filepath.Join("./vids/ss2/raw", file.Name())
			hlsURLs = append(hlsURLs, content)
		}
	}
	fmt.Println(hlsURLs)

	var wg sync.WaitGroup
	for i, hlsURL := range hlsURLs {
		wg.Add(1)
		go convertHLSFileToMP4(hlsURL, "./vids/ss2/spongebob_ss2", i, &wg)
	}
	wg.Wait()
}
