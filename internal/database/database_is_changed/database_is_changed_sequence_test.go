package database_is_changed

import (
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestIsChanged_Sequence(t *testing.T) {
	config_main.LoadEnvTest()
	config.FillSettings()

	postgres_gorm.Connect()
	defer postgres_gorm.CloseConnection()

	Otvet, err := IsChanged_Sequence()
	if err != nil {
		t.Error(err)
	}
	t.Log("IsChanged_Sequence(): ", Otvet)
}
