package models

import (
	"io"
)

// QuestionMetadata struct.
type QuestionMetadata struct {
	APIObject
	Question
	AnswerText        string `json:"answerText"`
	ImportantKeywords string `json:"importantKeywords"`
}

// ToJSON - encodes the object to JSON
func (q *QuestionMetadata) ToJSON(w io.Writer) error {
	return q.APIObject.ToJSON(q, w)
}

// FromJSON - encodes the object to JSON
func (q *QuestionMetadata) FromJSON(r io.Reader) error {
	return q.APIObject.FromJSON(q, r)
}
