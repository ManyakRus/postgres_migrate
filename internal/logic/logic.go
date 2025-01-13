package logic

import (
	"fmt"
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/postgres_migrate/internal/database/database_fill"
	"github.com/ManyakRus/postgres_migrate/internal/database/database_is_changed"
	"github.com/ManyakRus/postgres_migrate/internal/sql/sql_attributes"
	"github.com/ManyakRus/postgres_migrate/internal/sql/sql_tables"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
	"github.com/ManyakRus/starter/log"
)

// Start - старт работы логики всей программы
func Start(Settings *config.SettingsINI) {
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

	//
	TextSQL := ""
	TextSQL1 := ""

	//Tables
	TextSQL1, err = sql_tables.Start_Tables(Settings, Version.ID)
	if err != nil {
		err = fmt.Errorf("Start_Tables() error: %w", err)
		log.Error(err)
		return
	}
	TextSQL = TextSQL + TextSQL1

	//Attributes
	TextSQL1, err = sql_attributes.Start_Attributes(Settings, Version.ID)
	if err != nil {
		err = fmt.Errorf("Start_Tables() error: %w", err)
		log.Error(err)
		return
	}
	TextSQL = TextSQL + TextSQL1

	//
	log.Info(TextSQL)
}
