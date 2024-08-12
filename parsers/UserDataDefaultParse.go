package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// UserDataParser - парсер для файлов с данными пользователя
type UserDataParser struct{}

// Parse - реализация парсинга файлов с данными пользователя
func (u UserDataParser) Parse(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var userdata struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.Unmarshal(data, &userdata); err != nil {
		return fmt.Errorf("UserDataParser error: %v", err)
	}

	if userdata.Name == "" || userdata.Email == "" {
		return fmt.Errorf("UserDataParser: missing required fields")
	}

	return nil
}

func (c UserDataParser) GetId() string {
	return "UserData parser"
}
