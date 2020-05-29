package models

import (
	"io"
)

// TranscriptionCompletedEvent struct.
type TranscriptionCompletedEvent struct {
	APIObject
	SessionID            string                `json:"sessionId"`
	CandidateID          string                `json:"candidateId"`
	TranscriptionResults []TranscriptionResult `json:"transcriptionResults"`
}

// ToJSON - encodes the object to JSON
func (q *TranscriptionCompletedEvent) ToJSON(w io.Writer) error {
	return q.APIObject.ToJSON(q, w)
}

// FromJSON - encodes the object to JSON
func (q *TranscriptionCompletedEvent) FromJSON(r io.Reader) error {
	return q.APIObject.FromJSON(q, r)
}
