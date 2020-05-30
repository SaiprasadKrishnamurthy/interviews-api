package models

import (
	"io"
)

// QuestionResult struct.
type QuestionResult struct {
	APIObject
	SessionID                         string  `json:"sessionId"`
	CandidateID                       string  `json:"candidateId"`
	QuestionID                        string  `json:"questionId"`
	AutoAnswerSimilarityScore         float64 `json:"autoAnswerSimilarityScore"`
	AutoAnswerAbsoluteMatchingScore   float64 `json:"autoAnswerAbsoluteMatchingScore"`
	AutoKeywordsSimilarityScore       float64 `json:"autoKeywordsSimilarityScore"`
	AutoKeywordsAbsoluteMatchingScore float64 `json:"autoKeywordsAbsoluteMatchingScore"`
	ManualScore                       float64 `json:"manualScore"`
	Confidence                        float32 `json:"confidence"`
}

// ToJSON - encodes the object to JSON
func (q *QuestionResult) ToJSON(w io.Writer) error {
	return q.APIObject.ToJSON(q, w)
}

// FromJSON - encodes the object to JSON
func (q *QuestionResult) FromJSON(r io.Reader) error {
	return q.APIObject.FromJSON(q, r)
}
