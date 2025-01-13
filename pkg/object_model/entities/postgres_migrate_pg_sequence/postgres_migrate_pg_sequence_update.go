package postgres_migrate_pg_sequence

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
)

// Update_Seqcache - изменяет объект в БД по ID, присваивает Seqcache
func (m *PostgresMigratePgSequence) Update_Seqcache() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update_Seqcache(m)

	return err
}

// Update_Seqcycle - изменяет объект в БД по ID, присваивает Seqcycle
func (m *PostgresMigratePgSequence) Update_Seqcycle() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update_Seqcycle(m)

	return err
}

// Update_Seqincrement - изменяет объект в БД по ID, присваивает Seqincrement
func (m *PostgresMigratePgSequence) Update_Seqincrement() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update_Seqincrement(m)

	return err
}

// Update_Seqmax - изменяет объект в БД по ID, присваивает Seqmax
func (m *PostgresMigratePgSequence) Update_Seqmax() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update_Seqmax(m)

	return err
}

// Update_Seqmin - изменяет объект в БД по ID, присваивает Seqmin
func (m *PostgresMigratePgSequence) Update_Seqmin() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update_Seqmin(m)

	return err
}

// Update_Seqrelid - изменяет объект в БД по ID, присваивает Seqrelid
func (m *PostgresMigratePgSequence) Update_Seqrelid() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update_Seqrelid(m)

	return err
}

// Update_Seqstart - изменяет объект в БД по ID, присваивает Seqstart
func (m *PostgresMigratePgSequence) Update_Seqstart() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update_Seqstart(m)

	return err
}

// Update_Seqtypid - изменяет объект в БД по ID, присваивает Seqtypid
func (m *PostgresMigratePgSequence) Update_Seqtypid() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update_Seqtypid(m)

	return err
}

// Update_VersionID - изменяет объект в БД по ID, присваивает VersionID
func (m *PostgresMigratePgSequence) Update_VersionID() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update_VersionID(m)

	return err
}
