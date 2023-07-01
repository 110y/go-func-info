package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/110y/go-func-info/internal/analysis"
)

func main() {
	ctx := context.Background()

	var pos int
	flag.IntVar(&pos, "pos", 0, "position of cursor")

	var path string
	flag.StringVar(&path, "file", "", "file")

	flag.Parse()

	funcInfo, err := analysis.GetFuncInfo(ctx, path, pos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to analyze: %s\n", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(funcInfo); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode json: %s\n", err)
	}
}
