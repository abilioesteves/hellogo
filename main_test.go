package main

import (
	"net/http"
	"testing"
)

type MockResponseWriter struct {
	LastWrite []byte
}

func (m *MockResponseWriter) Header() http.Header {
	return make(map[string][]string)
}

func (m *MockResponseWriter) Write(data []byte) (int, error) {
	m.LastWrite = data
	return 200, nil
}

func (m *MockResponseWriter) WriteHeader(statuCode int) {

}
func TestHello(t *testing.T) {
	w := new(MockResponseWriter)
	hello(w, nil)

	if len(w.LastWrite) == 0 {
		t.Errorf("Expecting buffer LastWrite to be filled with data")
	}

	t.Logf("Response writer: %v", w)
}
