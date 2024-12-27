//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_description

import (
	"fmt"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
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
func (crud Crud_DB) UpdateManyFields(m *postgres_migrate_pg_description.PostgresMigratePgDescription, MassNeedUpdateFields []string) error {
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
func UpdateManyFields_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_description.PostgresMigratePgDescription, MassNeedUpdateFields []string) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	if (m.Classoid == 0) ||  (m.Objoid == 0) ||  (m.Objsubid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` UpdateManyFields() error: ID=0`)
		return err
	}

	err = create_update_ctx(ctx, db, m, MassNeedUpdateFields)
	return err
}

// Update_Classoid - изменяет 1 поле Classoid в базе данных
func (crud Crud_DB) Update_Classoid(m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Classoid_ctx(ctx, db, m)
	return err
}

// Update_Classoid_ctx - изменяет 1 поле Classoid в базе данных
// с учётом контекста и соединения к БД
func Update_Classoid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Classoid == 0) ||  (m.Objoid == 0) ||  (m.Objsubid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Classoid() error: Classoid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Classoid
	tx = db.Model(&m).Update("Classoid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Classoid() id: %v, error: %w", m.Classoid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_description.StringIdentifier(m.Classoid, m.Objoid, m.Objsubid, m.VersionID))

	return err
}

// Update_Description - изменяет 1 поле Description в базе данных
func (crud Crud_DB) Update_Description(m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Description_ctx(ctx, db, m)
	return err
}

// Update_Description_ctx - изменяет 1 поле Description в базе данных
// с учётом контекста и соединения к БД
func Update_Description_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Classoid == 0) ||  (m.Objoid == 0) ||  (m.Objsubid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Description() error: Classoid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Description
	tx = db.Model(&m).Update("Description", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Description() id: %v, error: %w", m.Classoid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_description.StringIdentifier(m.Classoid, m.Objoid, m.Objsubid, m.VersionID))

	return err
}

// Update_Objoid - изменяет 1 поле Objoid в базе данных
func (crud Crud_DB) Update_Objoid(m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Objoid_ctx(ctx, db, m)
	return err
}

// Update_Objoid_ctx - изменяет 1 поле Objoid в базе данных
// с учётом контекста и соединения к БД
func Update_Objoid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Classoid == 0) ||  (m.Objoid == 0) ||  (m.Objsubid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Objoid() error: Classoid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Objoid
	tx = db.Model(&m).Update("Objoid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Objoid() id: %v, error: %w", m.Classoid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_description.StringIdentifier(m.Classoid, m.Objoid, m.Objsubid, m.VersionID))

	return err
}

// Update_Objsubid - изменяет 1 поле Objsubid в базе данных
func (crud Crud_DB) Update_Objsubid(m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Objsubid_ctx(ctx, db, m)
	return err
}

// Update_Objsubid_ctx - изменяет 1 поле Objsubid в базе данных
// с учётом контекста и соединения к БД
func Update_Objsubid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Classoid == 0) ||  (m.Objoid == 0) ||  (m.Objsubid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Objsubid() error: Classoid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Objsubid
	tx = db.Model(&m).Update("Objsubid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Objsubid() id: %v, error: %w", m.Classoid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_description.StringIdentifier(m.Classoid, m.Objoid, m.Objsubid, m.VersionID))

	return err
}

// Update_VersionID - изменяет 1 поле VersionID в базе данных
func (crud Crud_DB) Update_VersionID(m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
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
func Update_VersionID_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_description.PostgresMigratePgDescription) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Classoid == 0) ||  (m.Objoid == 0) ||  (m.Objsubid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_VersionID() error: Classoid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.VersionID
	tx = db.Model(&m).Update("VersionID", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_VersionID() id: %v, error: %w", m.Classoid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_description.StringIdentifier(m.Classoid, m.Objoid, m.Objsubid, m.VersionID))

	return err
}
