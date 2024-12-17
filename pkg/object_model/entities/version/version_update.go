package version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	)


// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *Version) Update_Description() error {
	if Crud_Version == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_Version.Update_Description(m)

	return err
}

// Update_Name - изменяет объект в БД по ID, присваивает Name
func (m *Version) Update_Name() error {
	if Crud_Version == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_Version.Update_Name(m)

	return err
}
