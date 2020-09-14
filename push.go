package bullet

import (
	"bytes"
)

type pushStruct struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
}

func (p pushStruct) getReader() (*bytes.Buffer, error) {
	return getReader(p)
}

func newNotePush(title, text string) pushStruct {
	push := pushStruct{Type: "note", Title: title, Body: text}
	return push
}

func newLinkPush(title, text, link string) pushStruct {
	push := pushStruct{Type: "note", Title: title, Body: text, Url: link}
	return push
}
