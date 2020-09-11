package bullet

import (
	"bytes"
)

type pushInterface interface {
	getReader() (*bytes.Buffer, error)
}
