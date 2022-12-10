package helpers

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filePath string) []string {
	var output []string
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}
