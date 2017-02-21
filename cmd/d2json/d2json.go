package main

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/kelvyne/d2protocolparser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %v path.swf\n", os.Args[0])
		os.Exit(1)
	}

	swfPath := os.Args[1]
	p, err := d2protocolparser.Build(swfPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "build: %v\n", err)
		os.Exit(2)
	}

	jsonEncoder := json.NewEncoder(os.Stdout)
	if err := jsonEncoder.Encode(p); err != nil {
		fmt.Fprintf(os.Stderr, "json encoding failed: %v\n", err)
		os.Exit(2)
	}
}
