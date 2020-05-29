package config

// Configuration config object
type Configuration struct {
	Elastic     ElasticConfiguration
	Server      ServerConfiguration
	VideoStore  VideoStoreConfiguration
	Nats        NatsConfiguration
	Transcoding TranscodingConfiguration
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

// NatsConfiguration  config object
type NatsConfiguration struct {
	URL                         string
	InterviewCompletedSubject   string
	TranscodingCompletedSubject string
}

// ServerConfiguration  config object
type ServerConfiguration struct {
	Port string
}

// VideoStoreConfiguration  config object
type VideoStoreConfiguration struct {
	URI     string
	WorkDir string
}

// TranscodingConfiguration  config object
type TranscodingConfiguration struct {
	TimeoutInSeconds int
}
