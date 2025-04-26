package parser

import (
	"net/http"
)

type Parser struct{}

func (p *Parser) Parse() (*http.Request, error) {
	return &http.Request{}, nil
}
