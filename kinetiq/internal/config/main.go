package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port uint16

	TelemetryFileName string
	LapsFileName      string

	// OBS
	UseObs           uint8
	ObsBufferSeconds uint16
	ObsPassword      string
	ObsPort          uint16
	ObsAddr          string

	TelemetryFMBufferRows uint16
}

type ConfigUpdate struct {
	Port *uint16 `json:"port,omitempty"`

	TelemetryFileName *string `json:"telemetryFileName,omitempty"`
	LapsFileName      *string `json:"lapsFileName,omitempty"`

	// OBS
	UseObs           *uint8  `json:"useObs,omitempty"`
	ObsBufferSeconds *uint16 `json:"obsBufferSeconds,omitempty"`
	ObsPassword      *string `json:"obsPassword,omitempty"`
	ObsPort          *uint16 `json:"obsPort,omitempty"`
	ObsAddr          *string `json:"obsAddr,omitempty"`

	TelemetryFMBufferRows *uint16 `json:"telemetryFMBufferRows,omitempty"`
}

var (
	currentConfig *Config
	configPath    string
)

func NewConfig() (*Config, error) {
	// exePath, _ := os.Executable()
	configPath = "config.json"

	var err error
	currentConfig, err = loadConfig()
	if err != nil {
		return nil, err
	}

	return currentConfig, nil
}

func (c *Config) ChangeConfig(payload ConfigUpdate) Config {
	if payload.Port != nil {
		c.Port = *payload.Port
	}
	if payload.TelemetryFileName != nil {
		c.TelemetryFileName = *payload.TelemetryFileName
	}
	if payload.LapsFileName != nil {
		c.LapsFileName = *payload.LapsFileName
	}
	if payload.UseObs != nil {
		c.UseObs = *payload.UseObs
	}
	if payload.ObsBufferSeconds != nil {
		c.ObsBufferSeconds = *payload.ObsBufferSeconds
	}
	if payload.ObsPassword != nil {
		c.ObsPassword = *payload.ObsPassword
	}
	if payload.ObsPort != nil {
		c.ObsPort = *payload.ObsPort
	}
	if payload.ObsAddr != nil {
		c.ObsAddr = *payload.ObsAddr
	}
	if payload.TelemetryFMBufferRows != nil {
		c.TelemetryFMBufferRows = *payload.TelemetryFMBufferRows
	}

	saveConfig(c)

	return *c
}

func loadConfig() (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			defaultCfg := getDefaultConfig()
			saveConfig(defaultCfg)
			return defaultCfg, nil
		}
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func saveConfig(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}

func getDefaultConfig() *Config {
	return &Config{
		Port: 20779,

		TelemetryFileName: "telemetry.csv",
		LapsFileName:      "laps.csv",

		TelemetryFMBufferRows: 100,

		UseObs:           0,
		ObsBufferSeconds: 120,
		ObsPassword:      "",
		ObsPort:          4455,
		ObsAddr:          "localhost",
	}
}
