package bullet

import (
	"fmt"
	"net"
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

	return httptest.NewUnstartedServer(http.HandlerFunc(function))
}

func TestSendNote(t *testing.T) {
	server := fakeServer(http.StatusOK, "")
	server.Start()
	defer server.Close()

	b := Bullet{token: "", baseURL: server.URL}

	err := b.SendNote("test", "test")
	if err != nil {
		t.Error(err)
	}
}

func TestSendLink(t *testing.T) {
	server := fakeServer(http.StatusOK, "")
	server.Start()
	defer server.Close()

	b := Bullet{token: "", baseURL: server.URL}

	err := b.SendLink("test", "test", "url")
	if err != nil {
		t.Error(err)
	}
}

func TestSendFile(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:56022")
	if err != nil {
		t.Error(err)
	}

	response := `{
  "file_name": "cat.jpg",
  "file_type": "image/jpeg",
  "file_url": "https://dl.pushbulletusercontent.com/034f197bc6c37cac3cc03542659d458b/cat.jpg",
  "upload_url": "http://127.0.0.1:56022"}`

	server := fakeServer(http.StatusOK, response)
	server.Listener.Close()
	server.Listener = l
	server.Start()
	defer server.Close()

	b := Bullet{token: "", baseURL: server.URL}

	err = b.SendFile("test", "test", "./README.md")
	if err != nil {
		t.Error(err)
	}
}
