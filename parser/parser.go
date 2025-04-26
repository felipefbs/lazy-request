package parser

import (
	"fmt"
	"io"
	"net/http"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return tok, lit
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) Parse() (*http.Request, error) {
	var method string

	t, lit := p.scanIgnoreWhitespace()
	switch t {
	case GET:
		method = http.MethodGet
	case PATCH:
		method = http.MethodPatch
	default:
		return nil, fmt.Errorf("expected a method but got: %v", lit)
	}

	var path string
	for {
		t, lit := p.scanIgnoreWhitespace()
		if t == EOF {
			break
		}
		if t < IDENT {
			return nil, fmt.Errorf("illegal literal: %v", lit)
		}

		path += lit
	}

	return http.NewRequest(method, path, nil)
}
