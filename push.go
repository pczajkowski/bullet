package bullet

import (
	"bytes"
	"encoding/json"
)

type Push struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

//GetReader
func (p Push) GetReader() (*bytes.Buffer, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(jsonBytes)
	return buffer, err
}

//NewNotePush
func NewNotePush(title, text string) Push {
	push := Push{Type: "note", Title: title, Body: text}
	return push
}
