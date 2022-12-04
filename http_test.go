package errflow

import (
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	flow := NewHttpClientOf(&http.Client{})
	flow.Get("http://example.com")
}