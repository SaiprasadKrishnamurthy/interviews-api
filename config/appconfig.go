package config

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

var tokens []string
var configuration Configuration

// InitConfigs initializes all the necessary configs once.
func InitConfigs() {
	env := os.Getenv("ENVIRONMENT")
	if len(env) == 0 {
		env = "default"
	}
	viper.SetConfigName(env + "_config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}

// GetPort returns the Web server port.
func GetPort() string {
	return configuration.Server.Port
}

// APIVersion public API version.
func APIVersion() string {
	return "v1"
}

// GetVideoQuestionsURL gets the url.
func GetVideoQuestionsURL() string {
	return ""
}

// GetConfig get configuration object.
func GetConfig() *Configuration {
	return &configuration
}

// ElasticClient elastic http client.
func ElasticClient() *http.Client {
	if configuration.Elastic.SkipTLS {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return &http.Client{
		Timeout: time.Second * 10,
	}
}
