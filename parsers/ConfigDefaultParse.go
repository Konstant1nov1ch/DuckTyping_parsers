package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ConfigParser - парсер для конфигурационных файлов
type ConfigParser struct{}

// Parse - реализация парсинга конфигурационных файлов
func (c ConfigParser) Parse(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var config struct {
		AppName string `json:"app_name"`
		Version string `json:"version"`
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("ConfigParser error: %v", err)
	}

	if config.AppName == "" || config.Version == "" {
		return fmt.Errorf("ConfigParser: missing required fields")
	}

	return nil
}

func (c ConfigParser) GetId() string {
	return "Config parser"
}
