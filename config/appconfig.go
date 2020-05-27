package config

import (
	"os"
)

var tokens []string

// InitConfigs initializes all the necessary configs once.
func InitConfigs() {

}

// GetPort returns the Web server port.
func GetPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8083"
	}
	return port
}

// APIVersion public API version.
func APIVersion() string {
	return "v1"
}

// GetVideoQuestionsURL gets the url.
func GetVideoQuestionsURL() string {
	return "https://github.com/SaiprasadKrishnamurthy/sample_video_files/raw/master/%s.mp4"
}
