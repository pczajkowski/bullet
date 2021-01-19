package bullet

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

var standardError string = `{
  "error": {
    "cat": "~(=^‥^)",
    "message": "The resource could not be found.",
    "type": "invalid_request"
  }
}`

var standardErrorMessage string = "invalid_request: The resource could not be found."

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

	err := b.SendNote("test", "test", "")
	if err != nil {
		t.Error(err)
	}
}

func TestSendNoteFail(t *testing.T) {
	server := fakeServer(http.StatusBadRequest, standardError)
	server.Start()
	defer server.Close()

	b := Bullet{token: "", baseURL: server.URL}

	err := b.SendNote("test", "test", "")
	if err == nil {
		t.Error("There should be error")
	}

	if err.Error() != standardErrorMessage {
		t.Error(err)
	}
}

func TestSendLink(t *testing.T) {
	server := fakeServer(http.StatusOK, "")
	server.Start()
	defer server.Close()

	b := Bullet{token: "", baseURL: server.URL}

	err := b.SendLink("test", "test", "url", "")
	if err != nil {
		t.Error(err)
	}
}

func TestSendLinkFail(t *testing.T) {
	server := fakeServer(http.StatusBadRequest, standardError)
	server.Start()
	defer server.Close()

	b := Bullet{token: "", baseURL: server.URL}

	err := b.SendLink("test", "test", "url", "")
	if err == nil {
		t.Error("There should be error")
	}

	if err.Error() != standardErrorMessage {
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

	err = b.SendFile("test", "test", "./README.md", "")
	if err != nil {
		t.Error(err)
	}
}

func TestSendFileFail(t *testing.T) {
	server := fakeServer(http.StatusBadRequest, standardError)
	server.Start()
	defer server.Close()

	b := Bullet{token: "", baseURL: server.URL}

	err := b.SendFile("test", "test", "./README.md", "")
	if err == nil {
		t.Error("There should be error")
	}

	if err.Error() != standardErrorMessage {
		t.Error(err)
	}
}

func TestListDevices(t *testing.T) {
	devicesJSON := `{
  "devices": [
    {
      "active": true,
      "app_version": 8623,
      "created": 1412047948.579029,
      "iden": "ujpah72o0sjAoRtnM0jc",
      "manufacturer": "Apple",
      "model": "iPhone 5s (GSM)",
      "modified": 1412047948.579031,
      "nickname": "Elon Musk's iPhone",
      "push_token": "production:f73be0ee7877c8c7fa69b1468cde764f"
    }
  ]
}`

	server := fakeServer(http.StatusOK, devicesJSON)
	server.Start()
	defer server.Close()

	b := Bullet{token: "", baseURL: server.URL}

	devices, err := b.ListDevices()
	if err != nil {
		t.Error(err)
	}

	if devices == nil {
		t.Fatal("Devices shouldn't be nil!")
	}

	expectedID := "ujpah72o0sjAoRtnM0jc"
	deviceID := devices.Items[0].ID
	if deviceID != expectedID {
		t.Errorf("Device ID should be %s, but is %s", expectedID, deviceID)
	}
}
