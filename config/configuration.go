package config

// Configuration config object
type Configuration struct {
	Elastic    ElasticConfiguration
	Server     ServerConfiguration
	VideoStore VideoStoreConfiguration
}

// ElasticConfiguration elastic config object
type ElasticConfiguration struct {
	URI                    string
	SkipTLS                bool
	Username               string
	Password               string
	QuestionsMetadataIndex string
	SessionIndex           string
}

// ServerConfiguration  config object
type ServerConfiguration struct {
	Port string
}

// VideoStoreConfiguration  config object
type VideoStoreConfiguration struct {
	URI string
}
