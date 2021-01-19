package bullet

import (
	"bytes"
)

type pushStruct struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Url      string `json:"url"`
	DeviceID string `json:"device_iden"`
}

func (p pushStruct) getReader() (*bytes.Buffer, error) {
	return getReader(p)
}

func newNotePush(title, text, deviceID string) pushStruct {
	push := pushStruct{Type: "note", Title: title, Body: text, DeviceID: deviceID}
	return push
}

func newLinkPush(title, text, link, deviceID string) pushStruct {
	push := pushStruct{Type: "note", Title: title, Body: text, Url: link, DeviceID: deviceID}
	return push
}
