package bullet

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func fakeServer(statusCode int, data string) *httptest.Server {
	function := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		fmt.Fprint(w, data)
	}

	return httptest.NewServer(http.HandlerFunc(function))
}

func TestSendNote(t *testing.T) {
	server := fakeServer(http.StatusOK, "")
	defer server.Close()

	b := Bullet{token: "", baseURL: server.URL}

	err := b.SendNote("test", "test")
	if err != nil {
		t.Error(err)
	}
}
