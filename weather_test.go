package weather

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRead(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	}))
	defer ts.Close()

	c, e := Read(ts.URL)
	if e != nil {
		t.Error("Can't fetch content")
	}
	if string(c) != "Hello" {
		t.Error("Wrong content", string(c))
	}
}
