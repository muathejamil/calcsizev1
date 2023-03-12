package FileWordCount

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Count struct {
}

func CountTotalWordInDir(dirPath string) int {
	file, err := os.Open(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	list, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	total_words_count := 0
	for _, f := range list {
		total_words_count += CountWordFile(filepath.Join(dirPath, filepath.Base(f.Name())))
	}
	return total_words_count
}

func CountWordFile(path string) int {
	fileHandle, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanWords)
	count := 0
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			break
		}
		count++
	}
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}
	return count
}
