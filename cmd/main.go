package main

import (
	"flag"

	"github.com/gcharalla/real-contributions/scan"
	"github.com/gcharalla/real-contributions/stats"
)

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" && email != "" {
		scan.Scan(folder)
		stats.Stats(email)
	}
}
