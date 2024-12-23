package logic

import (
	"fmt"
	"github.com/ManyakRus/postgres_migrate/internal/database/database_fill"
	"github.com/ManyakRus/postgres_migrate/internal/database/database_is_changed"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
	"github.com/ManyakRus/starter/log"
)

// Start - старт работы логики всей программы
func Start() {
	//проверим есть ли изменения в метаданных
	TextChanged, err := database_is_changed.IsChanged_any()
	if err != nil {
		err = fmt.Errorf("IsChanged_any() error: %w", err)
		log.Error(err)
		return
	}

	if TextChanged == "" {
		log.Info("No changes")
		return
	}

	//создадим новую версию
	Version := postgres_migrate_version.PostgresMigrateVersion{}
	Version.Name = TextChanged
	err = Version.Save()
	if err != nil {
		err = fmt.Errorf("Version.Save() error: %w", err)
		log.Error(err)
		return
	}
	log.Infof("Saved new version id: %v, name: %v", Version.ID, Version.Name)

	//создадим файлы .sql

	//заполним все таблицы postgres_migrate
	err = database_fill.Fill_All(Version.ID)
	if err != nil {
		err = fmt.Errorf("Fill_All() error: %w", err)
		log.Error(err)
		return
	}

}
