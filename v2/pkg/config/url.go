package config

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func URLValues(url string, array *[]string) {

	// Make an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	// Check if the status code is okay (e.g., 200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP status code", response.StatusCode)
		os.Exit(1)
	}

	// Create a bufio.Reader to read the response body
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\r")
		*array = append(*array, line)
	}
	// Read and print the body line by line
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		break
	// 	}
	// 	line = strings.TrimRight(line, "\n")
	// 	line = strings.TrimRight(line, "\r")
	// 	*array = append(*array, line)
	// }
}
