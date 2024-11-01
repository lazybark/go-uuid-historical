package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lazybark/go-uuid-historical/uuidv7"
)

const (
	ansiCyan  = "\033[36m"
	ansiReset = "\033[0m"
)

func main() {
	layouts := flag.String("layouts", "2006-01-02 15:04:05.000 -0700,2006-01-02 15:04:05 -0700,2006-01-02T15:04:05.000Z,2006-01-02T15:04:05Z,2006-01-02 15:04:05,2006-01-02,2006.01.02 15:04:05", "Comma-separated list of timestamp layouts")
	flag.Parse()

	layoutList := splitLayouts(*layouts)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter timestamps (one per line):")

	for scanner.Scan() {
		timestampStr := scanner.Text()
		var t time.Time
		var err error

		for _, layout := range layoutList {
			t, err = time.Parse(layout, timestampStr)
			if err == nil {
				break
			}
		}

		if err != nil {
			fmt.Printf("Error parsing timestamp '%s': %v\n", timestampStr, err)
			continue
		}

		uuid, err := uuidv7.GenerateUUIDv7(t)
		if err != nil {
			fmt.Printf("Error generating UUID for timestamp '%s': %v\n", timestampStr, err)
			continue
		}

		fmt.Printf("UUID for timestamp '%s': %s%s%s\n", timestampStr, ansiCyan, uuid, ansiReset)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}

func splitLayouts(layouts string) []string {
	return strings.Split(layouts, ",")
}
