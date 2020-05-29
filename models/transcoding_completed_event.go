package models

import (
	"io"
)

// TranscodingCompletedEvent struct.
type TranscodingCompletedEvent struct {
	APIObject
	SessionID          string              `json:"sessionId"`
	CandidateID        string              `json:"candidateId"`
	TranscodingResults []TranscodingResult `json:"transcodingResults"`
}

// ToJSON - encodes the object to JSON
func (q *TranscodingCompletedEvent) ToJSON(w io.Writer) error {
	return q.APIObject.ToJSON(q, w)
}

// FromJSON - encodes the object to JSON
func (q *TranscodingCompletedEvent) FromJSON(r io.Reader) error {
	return q.APIObject.FromJSON(q, r)
}
