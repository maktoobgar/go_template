package config

type (
	Config struct {
		Database Database `yaml:"database"`
	}

	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"ssl_mode"`
		TimeZone string `yaml:"time_zone"`
		Charset  string `yaml:"charset"`
	}
)
