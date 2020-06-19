package main

import "fmt"

var (
	SghVersion = "unknown"
	GoVersion  = "unknown"
	BuildDate  = "unknown"
)

func main() {
	fmt.Printf("sgh %s (built: %s, go version: %s)\n", SghVersion, GoVersion, BuildDate)
}
