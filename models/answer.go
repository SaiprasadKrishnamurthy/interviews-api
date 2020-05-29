package models

import (
	"io"
)

// Answer struct.
type Answer struct {
	APIObject
	SessionID    string `json:"sessionId"`
	QuestionName string `json:"questionName"`
	QuestionID   string `json:"questionId"`
	CandidateID  string `json:"candidateId"`
	Timestamp    int32  `json:"timestamp"`
	AnswerID     string `json:"answerId"`
	AnswerText   string `json:"answerText"`
}

// ToJSON - encodes the object to JSON
func (q *Answer) ToJSON(w io.Writer) error {
	return q.APIObject.ToJSON(q, w)
}

// FromJSON - encodes the object to JSON
func (q *Answer) FromJSON(r io.Reader) error {
	return q.APIObject.FromJSON(q, r)
}
