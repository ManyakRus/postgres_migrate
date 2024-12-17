package pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	)


// Update_Nspacl - изменяет объект в БД по ID, присваивает Nspacl
func (m *PgNamespace) Update_Nspacl() error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgNamespace.Update_Nspacl(m)

	return err
}

// Update_Nspname - изменяет объект в БД по ID, присваивает Nspname
func (m *PgNamespace) Update_Nspname() error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgNamespace.Update_Nspname(m)

	return err
}

// Update_Nspowner - изменяет объект в БД по ID, присваивает Nspowner
func (m *PgNamespace) Update_Nspowner() error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgNamespace.Update_Nspowner(m)

	return err
}

// Update_Oid - изменяет объект в БД по ID, присваивает Oid
func (m *PgNamespace) Update_Oid() error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgNamespace.Update_Oid(m)

	return err
}

// Update_VersionID - изменяет объект в БД по ID, присваивает VersionID
func (m *PgNamespace) Update_VersionID() error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgNamespace.Update_VersionID(m)

	return err
}
