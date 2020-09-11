package bullet

import (
	"bytes"
	"encoding/json"
)

type pushStruct struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
}

func (p pushStruct) getReader() (*bytes.Buffer, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(jsonBytes)
	return buffer, err
}

func newNotePush(title, text string) pushStruct {
	push := pushStruct{Type: "note", Title: title, Body: text}
	return push
}

func newLinkPush(title, text, link string) pushStruct {
	push := pushStruct{Type: "note", Title: title, Body: text, Url: link}
	return push
}
