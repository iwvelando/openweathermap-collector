// Package config defines the data structures related to configuration and
// includes functions for loading and parsing the config.
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Configuration holds all configuration.
type Configuration struct {
	OpenWeatherMap OpenWeatherMap
	InfluxDB       InfluxDB
}

// OpenWeatherMap holds the openweathermap.org configuration
type OpenWeatherMap struct {
	BaseURL       string
	SkipVerifySsl bool
	ApiKey        string
	Latitude      float64
	Longitude     float64
	Units         string
	Language      string
}

// InfluxDB holds the connection parameters for InfluxDB
type InfluxDB struct {
	Address           string
	Username          string
	Password          string
	MeasurementPrefix string
	Database          string
	RetentionPolicy   string
	Token             string
	Organization      string
	Bucket            string
	SkipVerifySsl     bool
	FlushInterval     uint
}

// LoadConfiguration takes a file path as input and loads the YAML-formatted
// configuration there.
func LoadConfiguration(configPath string) (*Configuration, error) {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	var configuration Configuration
	err := viper.Unmarshal(&configuration)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %s", err)
	}

	return &configuration, nil
}
