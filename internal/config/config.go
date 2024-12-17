package config

import (
	"github.com/ManyakRus/starter/microl"
	"strconv"
)

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	PAUSE_TINKOFF_MS int
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	//var err error
	Settings = SettingsINI{}
	Name := ""
	s := ""

	//
	Name = "PAUSE_TINKOFF_MS"
	s = microl.Getenv(Name, true)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	Settings.PAUSE_TINKOFF_MS = i

}
