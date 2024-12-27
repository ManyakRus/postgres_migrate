//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_attribute

import (
	"fmt"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
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
func (crud Crud_DB) UpdateManyFields(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute, MassNeedUpdateFields []string) error {
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
func UpdateManyFields_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute, MassNeedUpdateFields []string) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` UpdateManyFields() error: ID=0`)
		return err
	}

	err = create_update_ctx(ctx, db, m, MassNeedUpdateFields)
	return err
}

// Update_Attalign - изменяет 1 поле Attalign в базе данных
func (crud Crud_DB) Update_Attalign(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attalign_ctx(ctx, db, m)
	return err
}

// Update_Attalign_ctx - изменяет 1 поле Attalign в базе данных
// с учётом контекста и соединения к БД
func Update_Attalign_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attalign() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attalign
	tx = db.Model(&m).Update("Attalign", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attalign() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attbyval - изменяет 1 поле Attbyval в базе данных
func (crud Crud_DB) Update_Attbyval(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attbyval_ctx(ctx, db, m)
	return err
}

// Update_Attbyval_ctx - изменяет 1 поле Attbyval в базе данных
// с учётом контекста и соединения к БД
func Update_Attbyval_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attbyval() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attbyval
	tx = db.Model(&m).Update("Attbyval", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attbyval() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attcacheoff - изменяет 1 поле Attcacheoff в базе данных
func (crud Crud_DB) Update_Attcacheoff(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attcacheoff_ctx(ctx, db, m)
	return err
}

// Update_Attcacheoff_ctx - изменяет 1 поле Attcacheoff в базе данных
// с учётом контекста и соединения к БД
func Update_Attcacheoff_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attcacheoff() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attcacheoff
	tx = db.Model(&m).Update("Attcacheoff", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attcacheoff() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attcollation - изменяет 1 поле Attcollation в базе данных
func (crud Crud_DB) Update_Attcollation(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attcollation_ctx(ctx, db, m)
	return err
}

// Update_Attcollation_ctx - изменяет 1 поле Attcollation в базе данных
// с учётом контекста и соединения к БД
func Update_Attcollation_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attcollation() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attcollation
	tx = db.Model(&m).Update("Attcollation", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attcollation() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attgenerated - изменяет 1 поле Attgenerated в базе данных
func (crud Crud_DB) Update_Attgenerated(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attgenerated_ctx(ctx, db, m)
	return err
}

// Update_Attgenerated_ctx - изменяет 1 поле Attgenerated в базе данных
// с учётом контекста и соединения к БД
func Update_Attgenerated_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attgenerated() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attgenerated
	tx = db.Model(&m).Update("Attgenerated", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attgenerated() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Atthasdef - изменяет 1 поле Atthasdef в базе данных
func (crud Crud_DB) Update_Atthasdef(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Atthasdef_ctx(ctx, db, m)
	return err
}

// Update_Atthasdef_ctx - изменяет 1 поле Atthasdef в базе данных
// с учётом контекста и соединения к БД
func Update_Atthasdef_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Atthasdef() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Atthasdef
	tx = db.Model(&m).Update("Atthasdef", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Atthasdef() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Atthasmissing - изменяет 1 поле Atthasmissing в базе данных
func (crud Crud_DB) Update_Atthasmissing(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Atthasmissing_ctx(ctx, db, m)
	return err
}

// Update_Atthasmissing_ctx - изменяет 1 поле Atthasmissing в базе данных
// с учётом контекста и соединения к БД
func Update_Atthasmissing_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Atthasmissing() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Atthasmissing
	tx = db.Model(&m).Update("Atthasmissing", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Atthasmissing() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attidentity - изменяет 1 поле Attidentity в базе данных
func (crud Crud_DB) Update_Attidentity(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attidentity_ctx(ctx, db, m)
	return err
}

// Update_Attidentity_ctx - изменяет 1 поле Attidentity в базе данных
// с учётом контекста и соединения к БД
func Update_Attidentity_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attidentity() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attidentity
	tx = db.Model(&m).Update("Attidentity", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attidentity() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attinhcount - изменяет 1 поле Attinhcount в базе данных
func (crud Crud_DB) Update_Attinhcount(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attinhcount_ctx(ctx, db, m)
	return err
}

// Update_Attinhcount_ctx - изменяет 1 поле Attinhcount в базе данных
// с учётом контекста и соединения к БД
func Update_Attinhcount_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attinhcount() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attinhcount
	tx = db.Model(&m).Update("Attinhcount", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attinhcount() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attisdropped - изменяет 1 поле Attisdropped в базе данных
func (crud Crud_DB) Update_Attisdropped(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attisdropped_ctx(ctx, db, m)
	return err
}

// Update_Attisdropped_ctx - изменяет 1 поле Attisdropped в базе данных
// с учётом контекста и соединения к БД
func Update_Attisdropped_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attisdropped() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attisdropped
	tx = db.Model(&m).Update("Attisdropped", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attisdropped() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attislocal - изменяет 1 поле Attislocal в базе данных
func (crud Crud_DB) Update_Attislocal(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attislocal_ctx(ctx, db, m)
	return err
}

// Update_Attislocal_ctx - изменяет 1 поле Attislocal в базе данных
// с учётом контекста и соединения к БД
func Update_Attislocal_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attislocal() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attislocal
	tx = db.Model(&m).Update("Attislocal", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attislocal() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attlen - изменяет 1 поле Attlen в базе данных
func (crud Crud_DB) Update_Attlen(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attlen_ctx(ctx, db, m)
	return err
}

// Update_Attlen_ctx - изменяет 1 поле Attlen в базе данных
// с учётом контекста и соединения к БД
func Update_Attlen_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attlen() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attlen
	tx = db.Model(&m).Update("Attlen", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attlen() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attmissingval - изменяет 1 поле Attmissingval в базе данных
func (crud Crud_DB) Update_Attmissingval(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attmissingval_ctx(ctx, db, m)
	return err
}

// Update_Attmissingval_ctx - изменяет 1 поле Attmissingval в базе данных
// с учётом контекста и соединения к БД
func Update_Attmissingval_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attmissingval() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attmissingval
	tx = db.Model(&m).Update("Attmissingval", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attmissingval() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attname - изменяет 1 поле Attname в базе данных
func (crud Crud_DB) Update_Attname(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attname_ctx(ctx, db, m)
	return err
}

// Update_Attname_ctx - изменяет 1 поле Attname в базе данных
// с учётом контекста и соединения к БД
func Update_Attname_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attname() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attname
	tx = db.Model(&m).Update("Attname", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attname() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attndims - изменяет 1 поле Attndims в базе данных
func (crud Crud_DB) Update_Attndims(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attndims_ctx(ctx, db, m)
	return err
}

// Update_Attndims_ctx - изменяет 1 поле Attndims в базе данных
// с учётом контекста и соединения к БД
func Update_Attndims_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attndims() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attndims
	tx = db.Model(&m).Update("Attndims", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attndims() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attnotnull - изменяет 1 поле Attnotnull в базе данных
func (crud Crud_DB) Update_Attnotnull(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attnotnull_ctx(ctx, db, m)
	return err
}

// Update_Attnotnull_ctx - изменяет 1 поле Attnotnull в базе данных
// с учётом контекста и соединения к БД
func Update_Attnotnull_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attnotnull() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attnotnull
	tx = db.Model(&m).Update("Attnotnull", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attnotnull() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attnum - изменяет 1 поле Attnum в базе данных
func (crud Crud_DB) Update_Attnum(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attnum_ctx(ctx, db, m)
	return err
}

// Update_Attnum_ctx - изменяет 1 поле Attnum в базе данных
// с учётом контекста и соединения к БД
func Update_Attnum_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attnum() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attnum
	tx = db.Model(&m).Update("Attnum", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attnum() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attrelid - изменяет 1 поле Attrelid в базе данных
func (crud Crud_DB) Update_Attrelid(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attrelid_ctx(ctx, db, m)
	return err
}

// Update_Attrelid_ctx - изменяет 1 поле Attrelid в базе данных
// с учётом контекста и соединения к БД
func Update_Attrelid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attrelid() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attrelid
	tx = db.Model(&m).Update("Attrelid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attrelid() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attstattarget - изменяет 1 поле Attstattarget в базе данных
func (crud Crud_DB) Update_Attstattarget(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attstattarget_ctx(ctx, db, m)
	return err
}

// Update_Attstattarget_ctx - изменяет 1 поле Attstattarget в базе данных
// с учётом контекста и соединения к БД
func Update_Attstattarget_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attstattarget() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attstattarget
	tx = db.Model(&m).Update("Attstattarget", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attstattarget() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Attstorage - изменяет 1 поле Attstorage в базе данных
func (crud Crud_DB) Update_Attstorage(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Attstorage_ctx(ctx, db, m)
	return err
}

// Update_Attstorage_ctx - изменяет 1 поле Attstorage в базе данных
// с учётом контекста и соединения к БД
func Update_Attstorage_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Attstorage() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Attstorage
	tx = db.Model(&m).Update("Attstorage", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Attstorage() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Atttypid - изменяет 1 поле Atttypid в базе данных
func (crud Crud_DB) Update_Atttypid(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Atttypid_ctx(ctx, db, m)
	return err
}

// Update_Atttypid_ctx - изменяет 1 поле Atttypid в базе данных
// с учётом контекста и соединения к БД
func Update_Atttypid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Atttypid() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Atttypid
	tx = db.Model(&m).Update("Atttypid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Atttypid() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_Atttypmod - изменяет 1 поле Atttypmod в базе данных
func (crud Crud_DB) Update_Atttypmod(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Atttypmod_ctx(ctx, db, m)
	return err
}

// Update_Atttypmod_ctx - изменяет 1 поле Atttypmod в базе данных
// с учётом контекста и соединения к БД
func Update_Atttypmod_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Atttypmod() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Atttypmod
	tx = db.Model(&m).Update("Atttypmod", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Atttypmod() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}

// Update_VersionID - изменяет 1 поле VersionID в базе данных
func (crud Crud_DB) Update_VersionID(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
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
func Update_VersionID_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Attname == "") ||  (m.Attrelid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_VersionID() error: Attname=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.VersionID
	tx = db.Model(&m).Update("VersionID", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_VersionID() id: %v, error: %w", m.Attname, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_attribute.StringIdentifier(m.Attname, m.Attrelid, m.VersionID))

	return err
}
