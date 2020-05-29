package models

// TranscriptionResult struct.
type TranscriptionResult struct {
	SessionID   string `json:"sessionId"`
	CandidateID string `json:"candidateId"`
	Question    string `json:"question"`
	Result      string `json:"result"`
}
