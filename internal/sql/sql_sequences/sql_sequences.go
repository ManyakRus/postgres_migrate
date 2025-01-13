package sql_attributes

import (
	"fmt"
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/starter/log"
)

// Start_Attributes - добавляет текст SQL в Text
func Start_Attributes(Settings *config.SettingsINI, VersionID int64) (string, error) {
	var err error
	Otvet := ""
	Otvet1 := ""

	//DELETE
	Otvet1, err = Start_Attributes_delete(Settings, VersionID)
	if err != nil {
		err = fmt.Errorf("Start_Attributes_delete() error: %w", err)
		log.Error(err)
		return Otvet, err
	}
	Otvet = Otvet + Otvet1

	//CREATE
	Otvet1, err = Start_Attributes_create(Settings, VersionID)
	if err != nil {
		err = fmt.Errorf("Start_Attributes_create() error: %w", err)
		log.Error(err)
		return Otvet, err
	}
	Otvet = Otvet + Otvet1

	//ALTER
	Otvet1, err = Start_Secuences_alter(Settings, VersionID)
	if err != nil {
		err = fmt.Errorf("Start_Secuences_alter() error: %w", err)
		log.Error(err)
		return Otvet, err
	}
	Otvet = Otvet + Otvet1

	return Otvet, err
}
