package config

type (
	Config struct {
		Databases    []Database `yaml:"databases"`
		Postgres     Database   `yaml:"postgres"`
		Translator   Translator `yaml:"translator"`
		Logging      Logging    `yaml:"logging"`
		Api          Api        `yaml:"api"`
		Grpc         Grpc       `yaml:"grpc"`
		Debug        bool       `yaml:"debug"`
		Domain       string     `yaml:"domain"`
		PWD          string     `yaml:"pwd"`
		AllowOrigins string     `yaml:"allow_origins"`
		SecretKey    string     `yaml:"secret_key"`
	}

	Database struct {
		Name     string `yaml:"name"`
		Type     string `yaml:"type"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"ssl_mode"`
		TimeZone string `yaml:"time_zone"`
		Charset  string `yaml:"charset"`
	}

	Translator struct {
		Path string `yaml:"path"`
	}

	Logging struct {
		Path         string `yaml:"path"`
		Pattern      string `yaml:"pattern"`
		MaxAge       string `yaml:"max_age"`
		RotationTime string `yaml:"rotation_time"`
		RotationSize string `yaml:"rotation_size"`
	}

	Api struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
	}

	Grpc struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
	}
)
