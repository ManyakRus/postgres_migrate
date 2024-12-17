package pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	)


// Update_Condeferrable - изменяет объект в БД по ID, присваивает Condeferrable
func (m *PgConstraint) Update_Condeferrable() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Condeferrable(m)

	return err
}

// Update_Condeferred - изменяет объект в БД по ID, присваивает Condeferred
func (m *PgConstraint) Update_Condeferred() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Condeferred(m)

	return err
}

// Update_Conexclop - изменяет объект в БД по ID, присваивает Conexclop
func (m *PgConstraint) Update_Conexclop() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conexclop(m)

	return err
}

// Update_Confdeltype - изменяет объект в БД по ID, присваивает Confdeltype
func (m *PgConstraint) Update_Confdeltype() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Confdeltype(m)

	return err
}

// Update_Conffeqop - изменяет объект в БД по ID, присваивает Conffeqop
func (m *PgConstraint) Update_Conffeqop() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conffeqop(m)

	return err
}

// Update_Confkey - изменяет объект в БД по ID, присваивает Confkey
func (m *PgConstraint) Update_Confkey() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Confkey(m)

	return err
}

// Update_Confmatchtype - изменяет объект в БД по ID, присваивает Confmatchtype
func (m *PgConstraint) Update_Confmatchtype() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Confmatchtype(m)

	return err
}

// Update_Confrelid - изменяет объект в БД по ID, присваивает Confrelid
func (m *PgConstraint) Update_Confrelid() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Confrelid(m)

	return err
}

// Update_Confupdtype - изменяет объект в БД по ID, присваивает Confupdtype
func (m *PgConstraint) Update_Confupdtype() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Confupdtype(m)

	return err
}

// Update_Conindid - изменяет объект в БД по ID, присваивает Conindid
func (m *PgConstraint) Update_Conindid() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conindid(m)

	return err
}

// Update_Coninhcount - изменяет объект в БД по ID, присваивает Coninhcount
func (m *PgConstraint) Update_Coninhcount() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Coninhcount(m)

	return err
}

// Update_Conislocal - изменяет объект в БД по ID, присваивает Conislocal
func (m *PgConstraint) Update_Conislocal() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conislocal(m)

	return err
}

// Update_Conkey - изменяет объект в БД по ID, присваивает Conkey
func (m *PgConstraint) Update_Conkey() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conkey(m)

	return err
}

// Update_Conname - изменяет объект в БД по ID, присваивает Conname
func (m *PgConstraint) Update_Conname() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conname(m)

	return err
}

// Update_Connamespace - изменяет объект в БД по ID, присваивает Connamespace
func (m *PgConstraint) Update_Connamespace() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Connamespace(m)

	return err
}

// Update_Connoinherit - изменяет объект в БД по ID, присваивает Connoinherit
func (m *PgConstraint) Update_Connoinherit() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Connoinherit(m)

	return err
}

// Update_Conparentid - изменяет объект в БД по ID, присваивает Conparentid
func (m *PgConstraint) Update_Conparentid() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conparentid(m)

	return err
}

// Update_Conpfeqop - изменяет объект в БД по ID, присваивает Conpfeqop
func (m *PgConstraint) Update_Conpfeqop() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conpfeqop(m)

	return err
}

// Update_Conppeqop - изменяет объект в БД по ID, присваивает Conppeqop
func (m *PgConstraint) Update_Conppeqop() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conppeqop(m)

	return err
}

// Update_Conrelid - изменяет объект в БД по ID, присваивает Conrelid
func (m *PgConstraint) Update_Conrelid() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Conrelid(m)

	return err
}

// Update_Contype - изменяет объект в БД по ID, присваивает Contype
func (m *PgConstraint) Update_Contype() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Contype(m)

	return err
}

// Update_Contypid - изменяет объект в БД по ID, присваивает Contypid
func (m *PgConstraint) Update_Contypid() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Contypid(m)

	return err
}

// Update_Convalidated - изменяет объект в БД по ID, присваивает Convalidated
func (m *PgConstraint) Update_Convalidated() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Convalidated(m)

	return err
}

// Update_Oid - изменяет объект в БД по ID, присваивает Oid
func (m *PgConstraint) Update_Oid() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_Oid(m)

	return err
}

// Update_VersionID - изменяет объект в БД по ID, присваивает VersionID
func (m *PgConstraint) Update_VersionID() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.Update_VersionID(m)

	return err
}
