package files

import (
	"fmt"
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/postgres_migrate/internal/files/files_tables"
	"github.com/ManyakRus/starter/log"
)

// Start_All - возвращает текст SQL для всех видов таблиц и колонок и др.
func Start_All(Settings *config.SettingsINI, VersionID int64) (string, error) {
	Otvet := ""
	var err error

	//
	Otvet1, err := files_tables.Start_Tables(Settings, VersionID)
	if err != nil {
		err = fmt.Errorf("Start_Tables() error: %w", err)
		log.Error(err)
		return Otvet, err
	}
	Otvet = Otvet + Otvet1

	return Otvet, err
}
