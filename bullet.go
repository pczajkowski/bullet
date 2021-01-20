package bullet

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

const (
	shortTimeFormat = "2006-01-02 15:04"
)

//Bullet client for the Pushbullet API
type Bullet struct {
	token   string
	baseURL string
}

//NewBullet creates new Bullet using provided token
func NewBullet(token string) Bullet {
	b := Bullet{token: token, baseURL: "https://api.pushbullet.com/v2"}
	return b
}

func (b Bullet) newRequest(body io.Reader, URL string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, URL, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Access-Token", b.token)
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}

func (b Bullet) newRequestPush(body io.Reader) (*http.Request, error) {
	return b.newRequest(body, b.baseURL+"/pushes")
}

func (b Bullet) newRequestUpload(body io.Reader) (*http.Request, error) {
	return b.newRequest(body, b.baseURL+"/upload-request")
}

func doRequest(request *http.Request) (*http.Response, error) {
	client := http.Client{}
	response, errResponse := client.Do(request)
	if errResponse != nil {
		return nil, errResponse
	}

	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()

		var errBullet bulletError
		decoder := json.NewDecoder(response.Body)
		errJSON := decoder.Decode(&errBullet)
		if errJSON != nil {
			return nil, errJSON
		}

		return nil, errBullet.getError()
	}

	return response, nil
}

func (b Bullet) send(push pushInterface) error {
	reader, errReader := push.getReader()
	if errReader != nil {
		return errReader
	}

	request, errRequest := b.newRequestPush(reader)
	if errRequest != nil {
		return errRequest
	}

	response, errResponse := doRequest(request)
	if errResponse != nil {
		return errResponse
	}
	response.Body.Close()

	return nil
}

//SendNote sends push note with given title and text, use empty string as deviceID to send to all
func (b Bullet) SendNote(title, text, deviceID string) error {
	note := newNotePush(title, text, deviceID)
	err := b.send(note)

	return err
}

//SendLink sends push link with given title, text and link, use empty string as deviceID to send to all
func (b Bullet) SendLink(title, text, link, deviceID string) error {
	linkPush := newLinkPush(title, text, link, deviceID)
	err := b.send(linkPush)

	return err
}

func (b Bullet) requestUpload(file fileUpload) (*fileUpload, error) {
	reader, errReader := file.getReader()
	if errReader != nil {
		return nil, errReader
	}

	request, errRequest := b.newRequestUpload(reader)
	if errRequest != nil {
		return nil, errRequest
	}

	response, errResponse := doRequest(request)
	if errResponse != nil {
		return nil, errResponse
	}
	defer response.Body.Close()

	var result fileUpload
	decoder := json.NewDecoder(response.Body)
	errJSON := decoder.Decode(&result)
	if errJSON != nil {
		return nil, errJSON
	}

	return &result, nil
}

func (b Bullet) uploadFile(path string) (*fileUpload, error) {
	extension := filepath.Ext(path)
	mimeType := mime.TypeByExtension(extension)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	filename := filepath.Base(path)
	fileUpload := newFileUpload(filename, mimeType)

	uploadResponse, errUpload := b.requestUpload(fileUpload)
	if errUpload != nil {
		return nil, errUpload
	}

	request, errRequest := getFileRequest(path, *uploadResponse)
	if errRequest != nil {
		return nil, errRequest
	}

	client := http.Client{}
	response, errResponse := client.Do(request)
	if errResponse != nil {
		return nil, errResponse
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("Error uploading file: %s", response.Status)
	}

	return uploadResponse, nil
}

//SendFile sends file push with given title, text and file, use empty string as deviceID to send to all
func (b Bullet) SendFile(title, text, file, deviceID string) error {
	uploadResult, errUpload := b.uploadFile(file)
	if errUpload != nil {
		return errUpload
	}

	if title != "" {
		uploadResult.Title = title
	} else {
		uploadResult.Title = time.Now().Format(shortTimeFormat)
	}
	uploadResult.Body = text
	uploadResult.Type = "file"
	uploadResult.DeviceID = deviceID

	err := b.send(uploadResult)
	return err
}

// ListDevices returns Devices structure which contains slice of devices
func (b Bullet) ListDevices() (*Devices, error) {
	request, err := http.NewRequest(http.MethodGet, b.baseURL+"/devices", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Access-Token", b.token)

	response, errResponse := doRequest(request)
	if errResponse != nil {
		return nil, errResponse
	}
	defer response.Body.Close()

	var result Devices
	decoder := json.NewDecoder(response.Body)
	errJSON := decoder.Decode(&result)
	if errJSON != nil {
		return nil, errJSON
	}

	return &result, nil
}

// ListPushes returns Pushes structure which contains slice of pushes
func (b Bullet) ListPushes(active bool, modifiedAfter *time.Time, limit int, cursor string) (*Pushes, error) {
	params := url.Values{}

	if active {
		params.Add("active", "true")
	} else {
		params.Add("active", "false")
	}

	if modifiedAfter != nil {
		params.Add("modified_after", fmt.Sprint(modifiedAfter.Unix()))
	}

	if limit > 0 {
		params.Add("limit", fmt.Sprint(limit))
	}

	params.Add("cursor", cursor)
	URL := fmt.Sprintf("%s/pushes?%s", b.baseURL, params.Encode())

	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Access-Token", b.token)

	response, errResponse := doRequest(request)
	if errResponse != nil {
		return nil, errResponse
	}
	defer response.Body.Close()

	var result Pushes
	decoder := json.NewDecoder(response.Body)
	errJSON := decoder.Decode(&result)
	if errJSON != nil {
		return nil, errJSON
	}

	return &result, nil
}
