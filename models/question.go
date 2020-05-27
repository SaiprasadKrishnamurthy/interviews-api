package models

import (
	"io"
)

// Questions struct.
type Questions struct {
	APIObject
	Questions []Question `json:"questions"`
}

// Question struct.
type Question struct {
	APIObject
	SessionID           string `json:"sessionId"`
	Sequence            int    `json:"sequence"`
	QuestionName        string `json:"questionName"`
	QuestionID          string `json:"questionId"`
	QuestionText        string `json:"questionText"`
	AnswerTimeInSeconds int    `json:"answerTimeInSeconds"`
}

// ToJSON - encodes the object to JSON
func (q *Questions) ToJSON(w io.Writer) error {
	return q.APIObject.ToJSON(q, w)
}

// FromJSON - encodes the object to JSON
func (q *Questions) FromJSON(r io.Reader) error {
	return q.APIObject.FromJSON(q, r)
}
