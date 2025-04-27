package parser_test

import (
	"os"
	"testing"

	"github.com/felipefbs/lazy-request/parser"
)

func TestParser(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		file, _ := os.Open("../get.http")

		request, err := parser.ParseHTTP(file, "/get.http")
		if err != nil || request == nil {
			t.Log(err)
			t.Error("request should not be empty and err should be nil")
		}

		t.Log(request)
	})
	t.Run("Post", func(t *testing.T) {
		file, _ := os.Open("../post.http")

		request, err := parser.ParseHTTP(file, "/post.http")
		if err != nil || request == nil {
			t.Log(err)
			t.Error("request should not be empty and err should be nil")
		}

		t.Log(request)
	})
	t.Run("Delete", func(t *testing.T) {
		file, _ := os.Open("../delete.http")

		request, err := parser.ParseHTTP(file, "delete.http")
		if err != nil || request == nil {
			t.Log(err)
			t.Error("request should not be empty and err should be nil")
		}

		t.Log(request)
	})
	t.Run("Put", func(t *testing.T) {
		file, _ := os.Open("../put.http")

		request, err := parser.ParseHTTP(file, "/put.http")
		if err != nil || request == nil {
			t.Log(err)
			t.Error("request should not be empty and err should be nil")
		}

		t.Log(request)
	})
	t.Run("Patch", func(t *testing.T) {
		file, _ := os.Open("../patch.http")

		request, err := parser.ParseHTTP(file, "/patch.http")
		if err != nil || request == nil {
			t.Log(err)
			t.Error("request should not be empty and err should be nil")
		}

		t.Log(request)
	})
}
