package pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	)


// Update_Oid - изменяет объект в БД по ID, присваивает Oid
func (m *PgClass) Update_Oid() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Oid(m)

	return err
}

// Update_Relallvisible - изменяет объект в БД по ID, присваивает Relallvisible
func (m *PgClass) Update_Relallvisible() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relallvisible(m)

	return err
}

// Update_Relam - изменяет объект в БД по ID, присваивает Relam
func (m *PgClass) Update_Relam() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relam(m)

	return err
}

// Update_Relchecks - изменяет объект в БД по ID, присваивает Relchecks
func (m *PgClass) Update_Relchecks() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relchecks(m)

	return err
}

// Update_Relfilenode - изменяет объект в БД по ID, присваивает Relfilenode
func (m *PgClass) Update_Relfilenode() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relfilenode(m)

	return err
}

// Update_Relforcerowsecurity - изменяет объект в БД по ID, присваивает Relforcerowsecurity
func (m *PgClass) Update_Relforcerowsecurity() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relforcerowsecurity(m)

	return err
}

// Update_Relfrozenxid - изменяет объект в БД по ID, присваивает Relfrozenxid
func (m *PgClass) Update_Relfrozenxid() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relfrozenxid(m)

	return err
}

// Update_Relhasindex - изменяет объект в БД по ID, присваивает Relhasindex
func (m *PgClass) Update_Relhasindex() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relhasindex(m)

	return err
}

// Update_Relhasrules - изменяет объект в БД по ID, присваивает Relhasrules
func (m *PgClass) Update_Relhasrules() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relhasrules(m)

	return err
}

// Update_Relhassubclass - изменяет объект в БД по ID, присваивает Relhassubclass
func (m *PgClass) Update_Relhassubclass() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relhassubclass(m)

	return err
}

// Update_Relhastriggers - изменяет объект в БД по ID, присваивает Relhastriggers
func (m *PgClass) Update_Relhastriggers() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relhastriggers(m)

	return err
}

// Update_Relispartition - изменяет объект в БД по ID, присваивает Relispartition
func (m *PgClass) Update_Relispartition() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relispartition(m)

	return err
}

// Update_Relispopulated - изменяет объект в БД по ID, присваивает Relispopulated
func (m *PgClass) Update_Relispopulated() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relispopulated(m)

	return err
}

// Update_Relisshared - изменяет объект в БД по ID, присваивает Relisshared
func (m *PgClass) Update_Relisshared() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relisshared(m)

	return err
}

// Update_Relkind - изменяет объект в БД по ID, присваивает Relkind
func (m *PgClass) Update_Relkind() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relkind(m)

	return err
}

// Update_Relminmxid - изменяет объект в БД по ID, присваивает Relminmxid
func (m *PgClass) Update_Relminmxid() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relminmxid(m)

	return err
}

// Update_Relname - изменяет объект в БД по ID, присваивает Relname
func (m *PgClass) Update_Relname() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relname(m)

	return err
}

// Update_Relnamespace - изменяет объект в БД по ID, присваивает Relnamespace
func (m *PgClass) Update_Relnamespace() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relnamespace(m)

	return err
}

// Update_Relnatts - изменяет объект в БД по ID, присваивает Relnatts
func (m *PgClass) Update_Relnatts() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relnatts(m)

	return err
}

// Update_Reloftype - изменяет объект в БД по ID, присваивает Reloftype
func (m *PgClass) Update_Reloftype() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Reloftype(m)

	return err
}

// Update_Relowner - изменяет объект в БД по ID, присваивает Relowner
func (m *PgClass) Update_Relowner() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relowner(m)

	return err
}

// Update_Relpages - изменяет объект в БД по ID, присваивает Relpages
func (m *PgClass) Update_Relpages() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relpages(m)

	return err
}

// Update_Relpersistence - изменяет объект в БД по ID, присваивает Relpersistence
func (m *PgClass) Update_Relpersistence() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relpersistence(m)

	return err
}

// Update_Relreplident - изменяет объект в БД по ID, присваивает Relreplident
func (m *PgClass) Update_Relreplident() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relreplident(m)

	return err
}

// Update_Relrewrite - изменяет объект в БД по ID, присваивает Relrewrite
func (m *PgClass) Update_Relrewrite() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relrewrite(m)

	return err
}

// Update_Relrowsecurity - изменяет объект в БД по ID, присваивает Relrowsecurity
func (m *PgClass) Update_Relrowsecurity() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Relrowsecurity(m)

	return err
}

// Update_Reltablespace - изменяет объект в БД по ID, присваивает Reltablespace
func (m *PgClass) Update_Reltablespace() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Reltablespace(m)

	return err
}

// Update_Reltoastrelid - изменяет объект в БД по ID, присваивает Reltoastrelid
func (m *PgClass) Update_Reltoastrelid() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Reltoastrelid(m)

	return err
}

// Update_Reltuples - изменяет объект в БД по ID, присваивает Reltuples
func (m *PgClass) Update_Reltuples() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Reltuples(m)

	return err
}

// Update_Reltype - изменяет объект в БД по ID, присваивает Reltype
func (m *PgClass) Update_Reltype() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_Reltype(m)

	return err
}

// Update_VersionID - изменяет объект в БД по ID, присваивает VersionID
func (m *PgClass) Update_VersionID() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.Update_VersionID(m)

	return err
}
