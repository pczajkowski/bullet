package bullet

import (
	"bytes"
	"encoding/json"
)

type pushInterface interface {
	getReader() (*bytes.Buffer, error)
}

func getReader(object interface{}) (*bytes.Buffer, error) {
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(jsonBytes)
	return buffer, nil
}
