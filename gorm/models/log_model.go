package models

import "encoding/json"

type Level int8

const (
	InfoLevel    Level = 1
	WarningLevel Level = 2
	ErrorLevel   Level = 3
)

func (s Level) MarshalJSON() ([]byte, error) {
	var str string
	switch s {
	case InfoLevel:
		str = "info"
	case WarningLevel:
		str = "warning"
	case ErrorLevel:
		str = "error"
	}
	//return json.Marshal(str)
	return json.Marshal(map[string]any{
		"value": int8(s),
		"label": str,
	})
}

type LogModel struct {
	ID    uint   `json:"id"`
	Title string `gorm:"size:32"`
	Level Level  `json:"level"`
}
