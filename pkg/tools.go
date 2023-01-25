package tools

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func Fuzz(targets []string) {
	var wg sync.WaitGroup
	for _, url := range targets {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			if err != nil {
				fmt.Print("Error Sending Request")
			}
		}(url)
	}
	wg.Wait()
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
