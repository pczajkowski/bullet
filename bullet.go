package bullet

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	shortTimeFormat = "2006-01-02 15:04"
)

//Bullet client for the Pushbullet API
type Bullet struct {
	token string
}

//NewBullet creates new Bullet using provided token
func NewBullet(token string) Bullet {
	b := Bullet{token: token}
	return b
}

func (b Bullet) newRequest(body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, "https://api.pushbullet.com/v2/pushes", body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Access-Token", b.token)
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}

func (b Bullet) send(push Push) error {
	reader, errReader := push.GetReader()
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

//SendNote sends push note with given title and text
func (b Bullet) SendNote(title, text string) error {
	note := NewNotePush(title, text)
	err := b.send(note)

	return err
}

//SendLink sends push link with given title, text and link
func (b Bullet) SendLink(title, text, link string) error {
	linkPush := NewLinkPush(title, text, link)
	err := b.send(linkPush)

	return err
}
