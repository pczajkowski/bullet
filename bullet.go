package bullet

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	shortTimeFormat = "2006-01-02 15:04"
)

type Bullet struct {
	Token string
}

func (b Bullet) newRequest(body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, "https://api.pushbullet.com/v2/pushes", body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Access-Token", b.Token)
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}

//Send push note with given title and text
func (b Bullet) SendNote(title, text string) error {
	note := NewNotePush(title, text)
	reader, errReader := note.GetReader()
	if errReader != nil {
		return errReader
	}

	request, errRequest := b.newRequest(reader)
	if errRequest != nil {
		return errRequest
	}

	client := http.Client{}
	response, errResponse := client.Do(request)
	if errResponse != nil {
		return errResponse
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var errBullet BulletError
		decoder := json.NewDecoder(response.Body)
		errJSON := decoder.Decode(&errBullet)
		if errJSON != nil {
			return errJSON
		}

		return errBullet.GetError()
	}

	return nil
}
