package bullet

import (
	"fmt"
)

type BulletError struct {
	Error struct {
		Cat     string `json:"cat"`
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

func (be BulletError) GetError() error {
	return fmt.Errorf("%s: %s", be.Error.Type, be.Error.Message)
}
