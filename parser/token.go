package parser

type Token int

const (
	ILLEGAL = iota
	EOF
	WS

	IDENT

	ASTERISK
	COMMA
	COLON
	SLASH

	GET
	POST
	PUT
	DELETE
	PATCH
	OPTIONS
	HEAD
)
