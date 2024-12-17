package pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	)


// Update_Classoid - изменяет объект в БД по ID, присваивает Classoid
func (m *PgDescription) Update_Classoid() error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgDescription.Update_Classoid(m)

	return err
}

// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *PgDescription) Update_Description() error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgDescription.Update_Description(m)

	return err
}

// Update_Objoid - изменяет объект в БД по ID, присваивает Objoid
func (m *PgDescription) Update_Objoid() error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgDescription.Update_Objoid(m)

	return err
}

// Update_Objsubid - изменяет объект в БД по ID, присваивает Objsubid
func (m *PgDescription) Update_Objsubid() error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgDescription.Update_Objsubid(m)

	return err
}

// Update_VersionID - изменяет объект в БД по ID, присваивает VersionID
func (m *PgDescription) Update_VersionID() error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgDescription.Update_VersionID(m)

	return err
}
