package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	// Main project configs
	cfg                     interface{}
	ErrUnknownFileExtention error = errors.New("unknown file extention")
)

// Checks if extention of the file is supported or not and if supported
// its configurations will be read and will be saved in cfg variable
func Parse(path string, cfg interface{}, errorOnFileNotFound bool) error {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parseYaml(path, cfg, errorOnFileNotFound)
	default:
		return ErrUnknownFileExtention
	}
}

// Reads "env.yaml" config file(if exists) and save them in `cfg` variable
func ReadLocalConfigs(cfg interface{}) error {
	if err := Parse("env.yaml", cfg, false); err != nil {
		return err
	} else if err := Parse("env.yml", cfg, false); err != nil {
		return err
	}
	return nil
}

// Gets the `conf` pointer and makes `cfg` global variable to point out where `conf` points
func SetConfig(conf interface{}) {
	cfg = conf
}

// Parses configuration data of a yaml file
func parseYaml(path string, cfg interface{}, errorOnFileNotFound bool) error {
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
