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
