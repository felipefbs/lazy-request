package store

import (
	"errors"
	"net/http"

	"github.com/felipefbs/lazy-request/parser"
)

type Store struct {
	parser parser.Parser
}

func New(p parser.Parser) *Store {
	return &Store{p}
}

func (s *Store) GetRequest() (*http.Request, error) {
	request, err := s.parser.Parse()
	if err != nil {
		return &http.Request{}, errors.New("error reading requisition")
	}
	return request, nil
}

func (s *Store) ExecuteRequest() (http.Response, error) {
	request, err := s.GetRequest()
	if err != nil {
		return http.Response{}, err
	}

	c := http.DefaultClient

	res, err := c.Do(request)
	if err != nil {
		return http.Response{}, err
	}

	return *res, nil
}
