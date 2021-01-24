package utils

type Config struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	TestDB struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"dbname"`
	} `json:"test-db"`
	LiveDB struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"dbname"`
	} `json:"live-db"`
	Deployment string `json:"deployment"`
	S3host     string `json:"s3-host"`
}

type ConfigError struct{}

func (*ConfigError) Error() string {
	return "Configuration Error."
}
