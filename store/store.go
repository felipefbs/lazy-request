package store

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/felipefbs/lazy-request/parser"
)

func getFilePaths(path string) []string {
	httpFiles := make([]string, 0, 10)

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && filepath.Ext(path) == ".http" {
			httpFiles = append(httpFiles, path)
		}

		return nil
	})

	return httpFiles
}

func ReadDirectory(path string) ([]parser.RequestAttrs, error) {
	httpFiles := getFilePaths(path)

	fmt.Printf("%+v", httpFiles)
	requests := make([]parser.RequestAttrs, 0, 10)

	for _, path := range httpFiles {
		reader, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		parsedRequest, err := parser.ParseHTTP(reader, path)
		reader.Close()

		if err != nil {
			return nil, err
		}

		requests = append(requests, *parsedRequest)
	}

	return requests, nil
}
