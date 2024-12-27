package postgres_migrate_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	)


// Update_Classoid - изменяет объект в БД по ID, присваивает Classoid
func (m *PostgresMigratePgDescription) Update_Classoid() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgDescription.Update_Classoid(m)

	return err
}

// Update_Description - изменяет объект в БД по ID, присваивает Description
func (m *PostgresMigratePgDescription) Update_Description() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgDescription.Update_Description(m)

	return err
}

// Update_Objoid - изменяет объект в БД по ID, присваивает Objoid
func (m *PostgresMigratePgDescription) Update_Objoid() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgDescription.Update_Objoid(m)

	return err
}

// Update_Objsubid - изменяет объект в БД по ID, присваивает Objsubid
func (m *PostgresMigratePgDescription) Update_Objsubid() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgDescription.Update_Objsubid(m)

	return err
}

// Update_VersionID - изменяет объект в БД по ID, присваивает VersionID
func (m *PostgresMigratePgDescription) Update_VersionID() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgDescription.Update_VersionID(m)

	return err
}
