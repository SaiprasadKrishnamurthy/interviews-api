package utils

import "time"

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

// ExtractResult extracts result from a channel with a max timeout in seconds.
func ExtractResult(c chan map[string]interface{}, timeoutSeconds int) map[string]interface{} {
	select {
	case res := <-c:
		return res
	case <-time.After(time.Duration(timeoutSeconds) * time.Second):
		return map[string]interface{}{}
	}
}
