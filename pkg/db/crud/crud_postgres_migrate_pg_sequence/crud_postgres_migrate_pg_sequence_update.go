//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_sequence

import (
	"context"
	"errors"
	"fmt"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gorm.io/gorm"
	"time"
)

// UpdateManyFields - изменяет несколько полей в базе данных
// MassNeedUpdateFields - список полей структуры golang для обновления
func (crud Crud_DB) UpdateManyFields(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence, MassNeedUpdateFields []string) error {
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
func UpdateManyFields_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence, MassNeedUpdateFields []string) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` UpdateManyFields() error: ID=0`)
		return err
	}

	err = create_update_ctx(ctx, db, m, MassNeedUpdateFields)
	return err
}

// Update_Seqcache - изменяет 1 поле Seqcache в базе данных
func (crud Crud_DB) Update_Seqcache(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Seqcache_ctx(ctx, db, m)
	return err
}

// Update_Seqcache_ctx - изменяет 1 поле Seqcache в базе данных
// с учётом контекста и соединения к БД
func Update_Seqcache_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Seqcache() error: Seqrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Seqcache
	tx = db.Model(&m).Update("Seqcache", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Seqcache() id: %v, error: %w", m.Seqrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_sequence.StringIdentifier(m.Seqrelid, m.VersionID))

	return err
}

// Update_Seqcycle - изменяет 1 поле Seqcycle в базе данных
func (crud Crud_DB) Update_Seqcycle(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Seqcycle_ctx(ctx, db, m)
	return err
}

// Update_Seqcycle_ctx - изменяет 1 поле Seqcycle в базе данных
// с учётом контекста и соединения к БД
func Update_Seqcycle_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Seqcycle() error: Seqrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Seqcycle
	tx = db.Model(&m).Update("Seqcycle", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Seqcycle() id: %v, error: %w", m.Seqrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_sequence.StringIdentifier(m.Seqrelid, m.VersionID))

	return err
}

// Update_Seqincrement - изменяет 1 поле Seqincrement в базе данных
func (crud Crud_DB) Update_Seqincrement(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Seqincrement_ctx(ctx, db, m)
	return err
}

// Update_Seqincrement_ctx - изменяет 1 поле Seqincrement в базе данных
// с учётом контекста и соединения к БД
func Update_Seqincrement_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Seqincrement() error: Seqrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Seqincrement
	tx = db.Model(&m).Update("Seqincrement", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Seqincrement() id: %v, error: %w", m.Seqrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_sequence.StringIdentifier(m.Seqrelid, m.VersionID))

	return err
}

// Update_Seqmax - изменяет 1 поле Seqmax в базе данных
func (crud Crud_DB) Update_Seqmax(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Seqmax_ctx(ctx, db, m)
	return err
}

// Update_Seqmax_ctx - изменяет 1 поле Seqmax в базе данных
// с учётом контекста и соединения к БД
func Update_Seqmax_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Seqmax() error: Seqrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Seqmax
	tx = db.Model(&m).Update("Seqmax", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Seqmax() id: %v, error: %w", m.Seqrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_sequence.StringIdentifier(m.Seqrelid, m.VersionID))

	return err
}

// Update_Seqmin - изменяет 1 поле Seqmin в базе данных
func (crud Crud_DB) Update_Seqmin(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Seqmin_ctx(ctx, db, m)
	return err
}

// Update_Seqmin_ctx - изменяет 1 поле Seqmin в базе данных
// с учётом контекста и соединения к БД
func Update_Seqmin_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Seqmin() error: Seqrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Seqmin
	tx = db.Model(&m).Update("Seqmin", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Seqmin() id: %v, error: %w", m.Seqrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_sequence.StringIdentifier(m.Seqrelid, m.VersionID))

	return err
}

// Update_Seqrelid - изменяет 1 поле Seqrelid в базе данных
func (crud Crud_DB) Update_Seqrelid(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Seqrelid_ctx(ctx, db, m)
	return err
}

// Update_Seqrelid_ctx - изменяет 1 поле Seqrelid в базе данных
// с учётом контекста и соединения к БД
func Update_Seqrelid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Seqrelid() error: Seqrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Seqrelid
	tx = db.Model(&m).Update("Seqrelid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Seqrelid() id: %v, error: %w", m.Seqrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_sequence.StringIdentifier(m.Seqrelid, m.VersionID))

	return err
}

// Update_Seqstart - изменяет 1 поле Seqstart в базе данных
func (crud Crud_DB) Update_Seqstart(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Seqstart_ctx(ctx, db, m)
	return err
}

// Update_Seqstart_ctx - изменяет 1 поле Seqstart в базе данных
// с учётом контекста и соединения к БД
func Update_Seqstart_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Seqstart() error: Seqrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Seqstart
	tx = db.Model(&m).Update("Seqstart", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Seqstart() id: %v, error: %w", m.Seqrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_sequence.StringIdentifier(m.Seqrelid, m.VersionID))

	return err
}

// Update_Seqtypid - изменяет 1 поле Seqtypid в базе данных
func (crud Crud_DB) Update_Seqtypid(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Seqtypid_ctx(ctx, db, m)
	return err
}

// Update_Seqtypid_ctx - изменяет 1 поле Seqtypid в базе данных
// с учётом контекста и соединения к БД
func Update_Seqtypid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Seqtypid() error: Seqrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Seqtypid
	tx = db.Model(&m).Update("Seqtypid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Seqtypid() id: %v, error: %w", m.Seqrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_sequence.StringIdentifier(m.Seqrelid, m.VersionID))

	return err
}

// Update_VersionID - изменяет 1 поле VersionID в базе данных
func (crud Crud_DB) Update_VersionID(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
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
func Update_VersionID_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Seqrelid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_VersionID() error: Seqrelid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.VersionID
	tx = db.Model(&m).Update("VersionID", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_VersionID() id: %v, error: %w", m.Seqrelid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_sequence.StringIdentifier(m.Seqrelid, m.VersionID))

	return err
}
