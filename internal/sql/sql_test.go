package sql

import (
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestStart_All(t *testing.T) {
	config_main.LoadEnvTest()
	config.FillSettings()

	postgres_gorm.Connect()
	defer postgres_gorm.CloseConnection()

	Otvet, err := Start_All(&config.Settings, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log("TestStart_All() len(Otvet): ", len(Otvet))
}
