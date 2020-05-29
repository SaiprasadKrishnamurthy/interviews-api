package models

import (
	"io"
)

// InterviewCompletedEvent struct.
type InterviewCompletedEvent struct {
	APIObject
	SessionID   string `json:"sessionId"`
	CandidateID string `json:"candidateId"`
}

// ToJSON - encodes the object to JSON
func (q *InterviewCompletedEvent) ToJSON(w io.Writer) error {
	return q.APIObject.ToJSON(q, w)
}

// FromJSON - encodes the object to JSON
func (q *InterviewCompletedEvent) FromJSON(r io.Reader) error {
	return q.APIObject.FromJSON(q, r)
}
