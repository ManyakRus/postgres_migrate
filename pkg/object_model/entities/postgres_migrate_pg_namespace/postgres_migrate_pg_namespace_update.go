package postgres_migrate_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
)

// Update_Nspacl - изменяет объект в БД по ID, присваивает Nspacl
func (m *PostgresMigratePgNamespace) Update_Nspacl() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Update_Nspacl(m)

	return err
}

// Update_Nspname - изменяет объект в БД по ID, присваивает Nspname
func (m *PostgresMigratePgNamespace) Update_Nspname() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Update_Nspname(m)

	return err
}

// Update_Nspowner - изменяет объект в БД по ID, присваивает Nspowner
func (m *PostgresMigratePgNamespace) Update_Nspowner() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Update_Nspowner(m)

	return err
}

// Update_Oid - изменяет объект в БД по ID, присваивает Oid
func (m *PostgresMigratePgNamespace) Update_Oid() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Update_Oid(m)

	return err
}

// Update_VersionID - изменяет объект в БД по ID, присваивает VersionID
func (m *PostgresMigratePgNamespace) Update_VersionID() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Update_VersionID(m)

	return err
}
