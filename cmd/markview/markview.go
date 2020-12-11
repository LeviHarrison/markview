package main

import (
	"fmt"
	"os"
	"time"

	"github.com/leviharrison/markview/generate"
	"github.com/leviharrison/markview/scan"
)

func main() {
	start := time.Now()

	files, err := scan.Scan()
	if err != nil {
		fmt.Printf("Error scanning for files: %v", err)
		os.Exit(1)
	}

	fmt.Println("Generating...")

	err = generate.Generate(files)
	if err != nil {
		fmt.Printf("Error generating site: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Done in %v seconds", time.Now().Sub(start).Seconds())
}
