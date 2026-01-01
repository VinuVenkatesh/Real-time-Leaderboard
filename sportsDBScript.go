package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("list of activities.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Remove leading "number. "
		if _, activity, ok := strings.Cut(line, ". "); ok {
			fmt.Println(activity)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
