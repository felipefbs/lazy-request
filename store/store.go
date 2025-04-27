package store

import (
	"io/fs"
	"path/filepath"

	"github.com/felipefbs/lazy-request/parser"
)

type Store struct {
	path     string
	requests []parser.RequestAttrs
}

func New(path string) *Store {
	return &Store{path: path}
}

func (s *Store) ReadDirectory() ([]parser.RequestAttrs, error) {
	httpFiles := make([]string, 0, 10)

	filepath.WalkDir(s.path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && filepath.Ext(path) == ".http" {
			httpFiles = append(httpFiles, path)
		}

		return nil
	})

	return nil, nil
}
