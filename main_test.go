package vantuz

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestRateLimit(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Write([]byte("Hello."))
	}))
	defer server.Close()

	const expected = 5
	req, err := http.NewRequestWithContext(context.Background(), "GET", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	for i := 0; i < expected; i++ {
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	if requestCount != expected {
		t.Errorf("Expected %d requests, got %d", expected, requestCount)
	}
}

func TestOneRequestManySend(t *testing.T) {
	const key = "grant_type"
	const val = "device_code"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		gotVal := r.Form.Get(key)
		if gotVal != val {
			t.Errorf("Expected %s=%s, got %s=%s", key, val, key, gotVal)
		}
	}))
	defer server.Close()

	form := map[string]string{
		key: val,
	}
	req := C().R().
		SetFormUrlMap(form)

	for i := 0; i < 10; i++ {
		fmt.Printf("Request: %d\n", i)
		_, err := req.Post(context.Background(), server.URL)
		if err != nil {
			t.Fatal(err)
		}
	}
}
