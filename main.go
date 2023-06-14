package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/araddon/dateparse"
)

func main() {
	now := time.Now()

	r, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(fmt.Errorf("reading from stdin: %w", err))
	}

	parsed, err := time.Parse(time.RFC3339Nano, string(r))
	if err != nil {
		parsed, err = dateparse.ParseLocal(string(r))
		if err != nil {
			log.Fatal(fmt.Errorf("parsing stdin as datetime: %w", err))
		}
	}

	since := now.Sub(parsed).Truncate(1 * time.Second).String()

	fmt.Printf("%s (%s ago)\n", parsed.Local().Format("Jan _2 2006 03:04:05PM MST"), since)
}
