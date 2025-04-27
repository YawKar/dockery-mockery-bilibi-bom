package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		fmt.Fprintln(os.Stderr, "path is empty")
		os.Exit(1)
	}

	visitor := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	}

	if err := filepath.WalkDir(path, visitor); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
