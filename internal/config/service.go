package config

// Returns database username
func GetDBUsername() string {
	return cfg.Database.Username
}
