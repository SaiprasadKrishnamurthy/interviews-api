package utils

import (
	"time"

	"github.com/saiprasadkrishnamurthy/interviews-api/models"
)

// Unique returns unique elements in a slice.
func Unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// ExtractTranscodingResult extracts result from a channel with a max timeout in seconds.
func ExtractTranscodingResult(c chan models.TranscodingResult, timeoutSeconds int) models.TranscodingResult {
	select {
	case res := <-c:
		return res
	case <-time.After(time.Duration(timeoutSeconds) * time.Second):
		return models.TranscodingResult{Result: "TimeoutError"}
	}
}

// ExtractTranscriptionResult extracts result from a channel with a max timeout in seconds.
func ExtractTranscriptionResult(c chan models.TranscriptionResult, timeoutSeconds int) models.TranscriptionResult {
	select {
	case res := <-c:
		return res
	case <-time.After(time.Duration(timeoutSeconds) * time.Second):
		return models.TranscriptionResult{Result: "TimeoutError"}
	}
}

// Filter questions by value supplied by a mapper function m from a slice.
func Filter(vs []models.QuestionMetadata, m func(models.QuestionMetadata) string, f func(string) bool) []models.QuestionMetadata {
	vsf := make([]models.QuestionMetadata, 0)
	for _, v := range vs {
		if f(m(v)) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
