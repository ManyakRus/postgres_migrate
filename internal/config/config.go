package config

import (
	"github.com/ManyakRus/starter/microl"
)

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	DB_SCHEME_DATABASE string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	//var err error
	Settings = SettingsINI{}
	Name := ""

	//
	Name = "DB_SCHEME_DATABASE"
	microl.Set_FieldFromEnv_String(&Settings, Name, true)

}
