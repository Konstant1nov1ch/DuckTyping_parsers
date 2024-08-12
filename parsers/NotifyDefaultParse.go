package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// NotifyParser - парсер для файлов с уведомлениями
type NotifyParser struct{}

// CustomTime - кастомный тип времени для обработки формата ISO 8601
type CustomTime struct {
	time.Time
}

// UnmarshalJSON - метод для разбора времени в формате ISO 8601
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		return fmt.Errorf("failed to parse time: %v", err)
	}

	ct.Time = parsedTime
	return nil
}

// Parse - реализация парсинга файлов с уведомлениями
func (c NotifyParser) Parse(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var notify struct {
		Timestamp CustomTime `json:"timestamp"`
		Email     string     `json:"email"`
		Total     float64    `json:"total"`
		IsTrue    bool       `json:"isTrue"`
	}
	if err := json.Unmarshal(data, &notify); err != nil {
		return fmt.Errorf("NotifyParser error: %v", err)
	}

	if notify.Email == "" {
		return fmt.Errorf("NotifyParser: missing required fields")
	}

	return nil
}
func (c NotifyParser) GetId() string {
	return "Notify parser"
}
