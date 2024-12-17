//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_pg_namespace

import (
	"fmt"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"context"
	"errors"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gorm.io/gorm"
	"time"
)

// UpdateManyFields - изменяет несколько полей в базе данных
// MassNeedUpdateFields - список полей структуры golang для обновления
func (crud Crud_DB) UpdateManyFields(m *pg_namespace.PgNamespace, MassNeedUpdateFields []string) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = UpdateManyFields_ctx(ctx, db, m, MassNeedUpdateFields)
	return err
}

// UpdateManyFields_ctx - изменяет несколько полей в базе данных
// с учётом контекста и соединения к БД
// MassNeedUpdateFields - список полей структуры golang для обновления
func UpdateManyFields_ctx(ctx context.Context, db *gorm.DB, m *pg_namespace.PgNamespace, MassNeedUpdateFields []string) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` UpdateManyFields() error: ID=0`)
		return err
	}

	err = create_update_ctx(ctx, db, m, MassNeedUpdateFields)
	return err
}

// Update_Nspacl - изменяет 1 поле Nspacl в базе данных
func (crud Crud_DB) Update_Nspacl(m *pg_namespace.PgNamespace) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Nspacl_ctx(ctx, db, m)
	return err
}

// Update_Nspacl_ctx - изменяет 1 поле Nspacl в базе данных
// с учётом контекста и соединения к БД
func Update_Nspacl_ctx(ctx context.Context, db *gorm.DB, m *pg_namespace.PgNamespace) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Nspacl() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Nspacl
	tx = db.Model(&m).Update("Nspacl", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Nspacl() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_namespace.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Nspname - изменяет 1 поле Nspname в базе данных
func (crud Crud_DB) Update_Nspname(m *pg_namespace.PgNamespace) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Nspname_ctx(ctx, db, m)
	return err
}

// Update_Nspname_ctx - изменяет 1 поле Nspname в базе данных
// с учётом контекста и соединения к БД
func Update_Nspname_ctx(ctx context.Context, db *gorm.DB, m *pg_namespace.PgNamespace) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Nspname() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Nspname
	tx = db.Model(&m).Update("Nspname", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Nspname() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_namespace.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Nspowner - изменяет 1 поле Nspowner в базе данных
func (crud Crud_DB) Update_Nspowner(m *pg_namespace.PgNamespace) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Nspowner_ctx(ctx, db, m)
	return err
}

// Update_Nspowner_ctx - изменяет 1 поле Nspowner в базе данных
// с учётом контекста и соединения к БД
func Update_Nspowner_ctx(ctx context.Context, db *gorm.DB, m *pg_namespace.PgNamespace) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Nspowner() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Nspowner
	tx = db.Model(&m).Update("Nspowner", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Nspowner() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_namespace.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Oid - изменяет 1 поле Oid в базе данных
func (crud Crud_DB) Update_Oid(m *pg_namespace.PgNamespace) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Oid_ctx(ctx, db, m)
	return err
}

// Update_Oid_ctx - изменяет 1 поле Oid в базе данных
// с учётом контекста и соединения к БД
func Update_Oid_ctx(ctx context.Context, db *gorm.DB, m *pg_namespace.PgNamespace) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Oid() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Oid
	tx = db.Model(&m).Update("Oid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Oid() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_namespace.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_VersionID - изменяет 1 поле VersionID в базе данных
func (crud Crud_DB) Update_VersionID(m *pg_namespace.PgNamespace) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_VersionID_ctx(ctx, db, m)
	return err
}

// Update_VersionID_ctx - изменяет 1 поле VersionID в базе данных
// с учётом контекста и соединения к БД
func Update_VersionID_ctx(ctx context.Context, db *gorm.DB, m *pg_namespace.PgNamespace) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_VersionID() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.VersionID
	tx = db.Model(&m).Update("VersionID", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_VersionID() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_namespace.StringIdentifier(m.Oid, m.VersionID))

	return err
}
