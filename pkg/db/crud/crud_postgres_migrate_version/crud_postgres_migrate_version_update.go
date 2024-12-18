//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_version

import (
	"context"
	"errors"
	"fmt"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gorm.io/gorm"
	"time"
)

// UpdateManyFields - изменяет несколько полей в базе данных
// MassNeedUpdateFields - список полей структуры golang для обновления
func (crud Crud_DB) UpdateManyFields(m *postgres_migrate_version.PostgresMigrateVersion, MassNeedUpdateFields []string) error {
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
func UpdateManyFields_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_version.PostgresMigrateVersion, MassNeedUpdateFields []string) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	if m.ID == 0 {
		err = errors.New(m.TableNameDB() + ` UpdateManyFields() error: ID=0`)
		return err
	}

	err = create_update_ctx(ctx, db, m, MassNeedUpdateFields)
	return err
}

// Update_Description - изменяет 1 поле Description в базе данных
func (crud Crud_DB) Update_Description(m *postgres_migrate_version.PostgresMigrateVersion) error {
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
func Update_Description_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_version.PostgresMigrateVersion) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if m.ID == 0 {
		err = errors.New(m.TableNameDB() + ` Update_Description() error: ID=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Description
	tx = db.Model(&m).Update("Description", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Description() id: %v, error: %w", m.ID, err)
		return err
	}

	//удалим из кэша
	cache.Remove(m.ID)

	return err
}

// Update_Name - изменяет 1 поле Name в базе данных
func (crud Crud_DB) Update_Name(m *postgres_migrate_version.PostgresMigrateVersion) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Name_ctx(ctx, db, m)
	return err
}

// Update_Name_ctx - изменяет 1 поле Name в базе данных
// с учётом контекста и соединения к БД
func Update_Name_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_version.PostgresMigrateVersion) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if m.ID == 0 {
		err = errors.New(m.TableNameDB() + ` Update_Name() error: ID=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Name
	tx = db.Model(&m).Update("Name", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Name() id: %v, error: %w", m.ID, err)
		return err
	}

	//удалим из кэша
	cache.Remove(m.ID)

	return err
}
