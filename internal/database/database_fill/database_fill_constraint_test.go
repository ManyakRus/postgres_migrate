package database_fill

import (
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestFill_constraint(t *testing.T) {
	//t.SkipNow() //деструктивный тест

	config_main.LoadEnvTest()
	config.FillSettings()

	postgres_gorm.Connect()
	defer postgres_gorm.CloseConnection()

	err := Fill_constraint(1)
	if err != nil {
		t.Error(err)
	}
}
