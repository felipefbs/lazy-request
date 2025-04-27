package store_test

import (
	"testing"

	"github.com/felipefbs/lazy-request/store"
)

func TestStore(t *testing.T) {
	t.Run("GetAllRequests", func(t *testing.T) {
		path := ".."

		requests, err := store.ReadDirectory(path)
		if err != nil {
			t.Log(err)
			t.Error("error reading the requests")
		}

		t.Log(requests)
	})
}
