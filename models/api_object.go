package models

import (
	"bufio"
	"encoding/json"
	"io"
)

// APIObject the base API object.
type APIObject struct {
}

// ToJSON converts to json.
func (a *APIObject) ToJSON(obj interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(obj)

}

// FromJSON converts to object from json.
func (a *APIObject) FromJSON(obj interface{}, r io.Reader) error {
	d := json.NewDecoder(bufio.NewReader(r))
	return d.Decode(&obj)
}
