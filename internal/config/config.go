package config

type (
	Config struct {
		Database Database `yaml:"database"`
	}

	Database struct {
		Username string `yaml:"username" env:"USERNAME,file"`
		Password string `yaml:"password" env:"PASSWORD,file"`
		DBName   string `yaml:"db_name" env:"DBNAME,file"`
		Host     string `yaml:"host" env:"HOST,file"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"ssl_mode"`
		TimeZone string `yaml:"time_zone"`
		Charset  string `yaml:"charset"`
	}
)
