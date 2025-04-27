package parser

import (
	"bufio"
	"errors"
	"io"
	"net/http"
	"strings"
)

type RequestAttrs struct {
	Name    string
	Request *http.Request
	Body    []byte
}

func ParseHTTP(input io.Reader) (*RequestAttrs, error) {
	scanner := bufio.NewScanner(input)

	var (
		methodLine    string
		headers       []string
		bodyBuffer    strings.Builder
		isReadingBody bool
		requestName   string
	)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "###") {
			requestName = strings.TrimSpace(strings.TrimPrefix(line, "###"))
		} else if isAMethod(line) {
			requestName = line
		}

		if isReadingBody {
			bodyBuffer.WriteString(line)
			bodyBuffer.WriteString("\n")
			continue
		}

		if len(strings.TrimSpace(line)) == 0 {
			isReadingBody = true
			continue
		}

		if methodLine == "" && (strings.HasPrefix(line, "GET") ||
			strings.HasPrefix(line, "POST") ||
			strings.HasPrefix(line, "PUT") ||
			strings.HasPrefix(line, "DELETE") ||
			strings.HasPrefix(line, "PATCH") ||
			strings.HasPrefix(line, "HEAD") ||
			strings.HasPrefix(line, "OPTIONS")) {
			methodLine = line
			continue
		}

		if isHeaderLine(line) {
			headers = append(headers, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return createRequest(requestName, methodLine, headers, bodyBuffer.String())
}

func isHeaderLine(line string) bool {
	return strings.Contains(line, ":")
}

func createRequest(requestName, methodLine string, headers []string, body string) (*RequestAttrs, error) {
	if methodLine == "" {
		return nil, errors.New("método HTTP não encontrado na requisição")
	}

	parts := strings.SplitN(methodLine, " ", 2)
	if len(parts) < 2 {
		return nil, errors.New("formato inválido na linha do método HTTP")
	}

	method := parts[0]
	urlPath := strings.TrimSpace(parts[1])

	bodyReader := strings.NewReader(body)

	httpReq, err := http.NewRequest(method, urlPath, bodyReader)
	if err != nil {
		return nil, errors.New("falha ao criar a requisição HTTP: " + err.Error())
	}

	for _, header := range headers {
		parts := strings.SplitN(header, ":", 2)
		if len(parts) == 2 {
			headerName := strings.TrimSpace(parts[0])
			headerValue := strings.TrimSpace(parts[1])
			httpReq.Header.Add(headerName, headerValue)
		}
	}

	response := &RequestAttrs{
		Name:    requestName,
		Request: httpReq,
		Body:    []byte(body),
	}

	return response, nil
}

func isAMethod(line string) bool {
	s := strings.Split(line, " ")
	switch s[0] {
	case http.MethodGet, http.MethodDelete, http.MethodPatch, http.MethodPost, http.MethodPut:
		return true
	default:
		return false
	}
}
