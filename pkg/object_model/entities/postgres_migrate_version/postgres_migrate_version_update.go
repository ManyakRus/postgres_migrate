package postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	)


// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *PostgresMigrateVersion) Update_Description() error {
	if Crud_PostgresMigrateVersion == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigrateVersion.Update_Description(m)

	return err
}

// Update_Name - изменяет объект в БД по ID, присваивает Name
func (m *PostgresMigrateVersion) Update_Name() error {
	if Crud_PostgresMigrateVersion == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigrateVersion.Update_Name(m)

	return err
}
