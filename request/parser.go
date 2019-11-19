package request

import (
	"encoding/json"
	"io"
)

func ParseJSON(body io.ReadCloser, destination interface{}) error {
	err := json.NewDecoder(body).Decode(destination)
	return err
}
