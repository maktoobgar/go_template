package config

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	// The address where all configuration files are stored.
	address string = "build/config/"
)

var (
	// Main project configs
	cfg                     *Config
	ErrUnknownFileExtention error = errors.New("unknown file extention")
)

// Checks if extention of the file is supported or not and if supported
// its configurations will be read and will be saved in cfg variable
func Parse(path string, cfg *Config, errorOnFileNotFound bool) error {
	switch fileExtention(path) {
	case "yaml":
		return parseYaml(path, cfg, errorOnFileNotFound)
	default:
		return ErrUnknownFileExtention
	}
}

// Reads "address_to_config_folder/config.yaml" file and save them in `cfg` variable
func ReadProjectConfigs(cfg *Config) error {
	if err := Parse(address+"config.yaml", cfg, true); err != nil {
		return err
	}
	return nil
}

// Reads "env.yaml" config file(if exists) and save them in `cfg` variable
func ReadLocalConfigs(cfg *Config) error {
	if err := Parse("env.yaml", cfg, false); err != nil {
		return err
	}
	return nil
}

// Gets the `conf` pointer and makes `cfg` global variable to point out where `conf` points
func SetConfig(conf *Config) {
	cfg = conf
}

// Returns extention of config file
func fileExtention(path string) string {
	var s []string = strings.Split(path, ".")
	return s[len(s)-1]
}

// Parses configuration data of a yaml file
func parseYaml(path string, cfg *Config, errorOnFileNotFound bool) error {
	file, err := os.Open(path)
	if err != nil {
		if errorOnFileNotFound {
			return err
		} else {
			return nil
		}
	}

	defer func() {
		e := file.Close()
		if e != nil {
			err = e
		}
	}()

	if err := yaml.NewDecoder(file).Decode(cfg); err != nil {
		return err
	}

	return err
}
