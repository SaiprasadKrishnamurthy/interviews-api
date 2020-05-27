package models

import (
	"io"
)

// Session struct.
type Session struct {
	APIObject
	SessionID          string   `json:"sessionId"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Instructions       []string `json:"instructions"`
	TotalTimeInSeconds int      `json:"totalTimeInSeconds"`
}

// ToJSON - encodes the object to JSON
func (q *Session) ToJSON(w io.Writer) error {
	return q.APIObject.ToJSON(q, w)
}

// FromJSON - encodes the object to JSON
func (q *Session) FromJSON(r io.Reader) error {
	return q.APIObject.FromJSON(q, r)
}
