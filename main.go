package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	fmt.Printf("GoReleaser test! %s, commit %s, built at %s by %s", version, commit, date, builtBy)
}
