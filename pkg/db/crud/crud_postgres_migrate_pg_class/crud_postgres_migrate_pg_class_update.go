//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_class

import (
	"context"
	"errors"
	"fmt"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gorm.io/gorm"
	"time"
)

// UpdateManyFields - изменяет несколько полей в базе данных
// MassNeedUpdateFields - список полей структуры golang для обновления
func (crud Crud_DB) UpdateManyFields(m *postgres_migrate_pg_class.PostgresMigratePgClass, MassNeedUpdateFields []string) error {
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
func UpdateManyFields_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass, MassNeedUpdateFields []string) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` UpdateManyFields() error: ID=0`)
		return err
	}

	err = create_update_ctx(ctx, db, m, MassNeedUpdateFields)
	return err
}

// Update_Oid - изменяет 1 поле Oid в базе данных
func (crud Crud_DB) Update_Oid(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
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
func Update_Oid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Oid() error: Oid=0`)
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
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relallvisible - изменяет 1 поле Relallvisible в базе данных
func (crud Crud_DB) Update_Relallvisible(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relallvisible_ctx(ctx, db, m)
	return err
}

// Update_Relallvisible_ctx - изменяет 1 поле Relallvisible в базе данных
// с учётом контекста и соединения к БД
func Update_Relallvisible_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relallvisible() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relallvisible
	tx = db.Model(&m).Update("Relallvisible", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relallvisible() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relam - изменяет 1 поле Relam в базе данных
func (crud Crud_DB) Update_Relam(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relam_ctx(ctx, db, m)
	return err
}

// Update_Relam_ctx - изменяет 1 поле Relam в базе данных
// с учётом контекста и соединения к БД
func Update_Relam_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relam() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relam
	tx = db.Model(&m).Update("Relam", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relam() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relchecks - изменяет 1 поле Relchecks в базе данных
func (crud Crud_DB) Update_Relchecks(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relchecks_ctx(ctx, db, m)
	return err
}

// Update_Relchecks_ctx - изменяет 1 поле Relchecks в базе данных
// с учётом контекста и соединения к БД
func Update_Relchecks_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relchecks() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relchecks
	tx = db.Model(&m).Update("Relchecks", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relchecks() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relfilenode - изменяет 1 поле Relfilenode в базе данных
func (crud Crud_DB) Update_Relfilenode(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relfilenode_ctx(ctx, db, m)
	return err
}

// Update_Relfilenode_ctx - изменяет 1 поле Relfilenode в базе данных
// с учётом контекста и соединения к БД
func Update_Relfilenode_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relfilenode() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relfilenode
	tx = db.Model(&m).Update("Relfilenode", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relfilenode() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relforcerowsecurity - изменяет 1 поле Relforcerowsecurity в базе данных
func (crud Crud_DB) Update_Relforcerowsecurity(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relforcerowsecurity_ctx(ctx, db, m)
	return err
}

// Update_Relforcerowsecurity_ctx - изменяет 1 поле Relforcerowsecurity в базе данных
// с учётом контекста и соединения к БД
func Update_Relforcerowsecurity_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relforcerowsecurity() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relforcerowsecurity
	tx = db.Model(&m).Update("Relforcerowsecurity", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relforcerowsecurity() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relfrozenxid - изменяет 1 поле Relfrozenxid в базе данных
func (crud Crud_DB) Update_Relfrozenxid(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relfrozenxid_ctx(ctx, db, m)
	return err
}

// Update_Relfrozenxid_ctx - изменяет 1 поле Relfrozenxid в базе данных
// с учётом контекста и соединения к БД
func Update_Relfrozenxid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relfrozenxid() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relfrozenxid
	tx = db.Model(&m).Update("Relfrozenxid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relfrozenxid() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relhasindex - изменяет 1 поле Relhasindex в базе данных
func (crud Crud_DB) Update_Relhasindex(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relhasindex_ctx(ctx, db, m)
	return err
}

// Update_Relhasindex_ctx - изменяет 1 поле Relhasindex в базе данных
// с учётом контекста и соединения к БД
func Update_Relhasindex_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relhasindex() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relhasindex
	tx = db.Model(&m).Update("Relhasindex", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relhasindex() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relhasrules - изменяет 1 поле Relhasrules в базе данных
func (crud Crud_DB) Update_Relhasrules(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relhasrules_ctx(ctx, db, m)
	return err
}

// Update_Relhasrules_ctx - изменяет 1 поле Relhasrules в базе данных
// с учётом контекста и соединения к БД
func Update_Relhasrules_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relhasrules() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relhasrules
	tx = db.Model(&m).Update("Relhasrules", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relhasrules() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relhassubclass - изменяет 1 поле Relhassubclass в базе данных
func (crud Crud_DB) Update_Relhassubclass(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relhassubclass_ctx(ctx, db, m)
	return err
}

// Update_Relhassubclass_ctx - изменяет 1 поле Relhassubclass в базе данных
// с учётом контекста и соединения к БД
func Update_Relhassubclass_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relhassubclass() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relhassubclass
	tx = db.Model(&m).Update("Relhassubclass", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relhassubclass() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relhastriggers - изменяет 1 поле Relhastriggers в базе данных
func (crud Crud_DB) Update_Relhastriggers(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relhastriggers_ctx(ctx, db, m)
	return err
}

// Update_Relhastriggers_ctx - изменяет 1 поле Relhastriggers в базе данных
// с учётом контекста и соединения к БД
func Update_Relhastriggers_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relhastriggers() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relhastriggers
	tx = db.Model(&m).Update("Relhastriggers", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relhastriggers() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relispartition - изменяет 1 поле Relispartition в базе данных
func (crud Crud_DB) Update_Relispartition(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relispartition_ctx(ctx, db, m)
	return err
}

// Update_Relispartition_ctx - изменяет 1 поле Relispartition в базе данных
// с учётом контекста и соединения к БД
func Update_Relispartition_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relispartition() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relispartition
	tx = db.Model(&m).Update("Relispartition", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relispartition() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relispopulated - изменяет 1 поле Relispopulated в базе данных
func (crud Crud_DB) Update_Relispopulated(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relispopulated_ctx(ctx, db, m)
	return err
}

// Update_Relispopulated_ctx - изменяет 1 поле Relispopulated в базе данных
// с учётом контекста и соединения к БД
func Update_Relispopulated_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relispopulated() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relispopulated
	tx = db.Model(&m).Update("Relispopulated", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relispopulated() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relisshared - изменяет 1 поле Relisshared в базе данных
func (crud Crud_DB) Update_Relisshared(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relisshared_ctx(ctx, db, m)
	return err
}

// Update_Relisshared_ctx - изменяет 1 поле Relisshared в базе данных
// с учётом контекста и соединения к БД
func Update_Relisshared_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relisshared() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relisshared
	tx = db.Model(&m).Update("Relisshared", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relisshared() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relkind - изменяет 1 поле Relkind в базе данных
func (crud Crud_DB) Update_Relkind(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relkind_ctx(ctx, db, m)
	return err
}

// Update_Relkind_ctx - изменяет 1 поле Relkind в базе данных
// с учётом контекста и соединения к БД
func Update_Relkind_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relkind() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relkind
	tx = db.Model(&m).Update("Relkind", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relkind() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relminmxid - изменяет 1 поле Relminmxid в базе данных
func (crud Crud_DB) Update_Relminmxid(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relminmxid_ctx(ctx, db, m)
	return err
}

// Update_Relminmxid_ctx - изменяет 1 поле Relminmxid в базе данных
// с учётом контекста и соединения к БД
func Update_Relminmxid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relminmxid() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relminmxid
	tx = db.Model(&m).Update("Relminmxid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relminmxid() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relname - изменяет 1 поле Relname в базе данных
func (crud Crud_DB) Update_Relname(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relname_ctx(ctx, db, m)
	return err
}

// Update_Relname_ctx - изменяет 1 поле Relname в базе данных
// с учётом контекста и соединения к БД
func Update_Relname_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relname() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relname
	tx = db.Model(&m).Update("Relname", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relname() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relnamespace - изменяет 1 поле Relnamespace в базе данных
func (crud Crud_DB) Update_Relnamespace(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relnamespace_ctx(ctx, db, m)
	return err
}

// Update_Relnamespace_ctx - изменяет 1 поле Relnamespace в базе данных
// с учётом контекста и соединения к БД
func Update_Relnamespace_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relnamespace() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relnamespace
	tx = db.Model(&m).Update("Relnamespace", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relnamespace() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relnatts - изменяет 1 поле Relnatts в базе данных
func (crud Crud_DB) Update_Relnatts(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relnatts_ctx(ctx, db, m)
	return err
}

// Update_Relnatts_ctx - изменяет 1 поле Relnatts в базе данных
// с учётом контекста и соединения к БД
func Update_Relnatts_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relnatts() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relnatts
	tx = db.Model(&m).Update("Relnatts", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relnatts() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Reloftype - изменяет 1 поле Reloftype в базе данных
func (crud Crud_DB) Update_Reloftype(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Reloftype_ctx(ctx, db, m)
	return err
}

// Update_Reloftype_ctx - изменяет 1 поле Reloftype в базе данных
// с учётом контекста и соединения к БД
func Update_Reloftype_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Reloftype() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Reloftype
	tx = db.Model(&m).Update("Reloftype", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Reloftype() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relowner - изменяет 1 поле Relowner в базе данных
func (crud Crud_DB) Update_Relowner(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relowner_ctx(ctx, db, m)
	return err
}

// Update_Relowner_ctx - изменяет 1 поле Relowner в базе данных
// с учётом контекста и соединения к БД
func Update_Relowner_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relowner() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relowner
	tx = db.Model(&m).Update("Relowner", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relowner() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relpages - изменяет 1 поле Relpages в базе данных
func (crud Crud_DB) Update_Relpages(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relpages_ctx(ctx, db, m)
	return err
}

// Update_Relpages_ctx - изменяет 1 поле Relpages в базе данных
// с учётом контекста и соединения к БД
func Update_Relpages_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relpages() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relpages
	tx = db.Model(&m).Update("Relpages", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relpages() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relpersistence - изменяет 1 поле Relpersistence в базе данных
func (crud Crud_DB) Update_Relpersistence(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relpersistence_ctx(ctx, db, m)
	return err
}

// Update_Relpersistence_ctx - изменяет 1 поле Relpersistence в базе данных
// с учётом контекста и соединения к БД
func Update_Relpersistence_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relpersistence() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relpersistence
	tx = db.Model(&m).Update("Relpersistence", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relpersistence() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relreplident - изменяет 1 поле Relreplident в базе данных
func (crud Crud_DB) Update_Relreplident(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relreplident_ctx(ctx, db, m)
	return err
}

// Update_Relreplident_ctx - изменяет 1 поле Relreplident в базе данных
// с учётом контекста и соединения к БД
func Update_Relreplident_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relreplident() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relreplident
	tx = db.Model(&m).Update("Relreplident", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relreplident() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relrewrite - изменяет 1 поле Relrewrite в базе данных
func (crud Crud_DB) Update_Relrewrite(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relrewrite_ctx(ctx, db, m)
	return err
}

// Update_Relrewrite_ctx - изменяет 1 поле Relrewrite в базе данных
// с учётом контекста и соединения к БД
func Update_Relrewrite_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relrewrite() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relrewrite
	tx = db.Model(&m).Update("Relrewrite", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relrewrite() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Relrowsecurity - изменяет 1 поле Relrowsecurity в базе данных
func (crud Crud_DB) Update_Relrowsecurity(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Relrowsecurity_ctx(ctx, db, m)
	return err
}

// Update_Relrowsecurity_ctx - изменяет 1 поле Relrowsecurity в базе данных
// с учётом контекста и соединения к БД
func Update_Relrowsecurity_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Relrowsecurity() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Relrowsecurity
	tx = db.Model(&m).Update("Relrowsecurity", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Relrowsecurity() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Reltablespace - изменяет 1 поле Reltablespace в базе данных
func (crud Crud_DB) Update_Reltablespace(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Reltablespace_ctx(ctx, db, m)
	return err
}

// Update_Reltablespace_ctx - изменяет 1 поле Reltablespace в базе данных
// с учётом контекста и соединения к БД
func Update_Reltablespace_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Reltablespace() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Reltablespace
	tx = db.Model(&m).Update("Reltablespace", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Reltablespace() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Reltoastrelid - изменяет 1 поле Reltoastrelid в базе данных
func (crud Crud_DB) Update_Reltoastrelid(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Reltoastrelid_ctx(ctx, db, m)
	return err
}

// Update_Reltoastrelid_ctx - изменяет 1 поле Reltoastrelid в базе данных
// с учётом контекста и соединения к БД
func Update_Reltoastrelid_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Reltoastrelid() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Reltoastrelid
	tx = db.Model(&m).Update("Reltoastrelid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Reltoastrelid() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Reltuples - изменяет 1 поле Reltuples в базе данных
func (crud Crud_DB) Update_Reltuples(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Reltuples_ctx(ctx, db, m)
	return err
}

// Update_Reltuples_ctx - изменяет 1 поле Reltuples в базе данных
// с учётом контекста и соединения к БД
func Update_Reltuples_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Reltuples() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Reltuples
	tx = db.Model(&m).Update("Reltuples", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Reltuples() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Reltype - изменяет 1 поле Reltype в базе данных
func (crud Crud_DB) Update_Reltype(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Reltype_ctx(ctx, db, m)
	return err
}

// Update_Reltype_ctx - изменяет 1 поле Reltype в базе данных
// с учётом контекста и соединения к БД
func Update_Reltype_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_Reltype() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Reltype
	tx = db.Model(&m).Update("Reltype", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Reltype() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_VersionID - изменяет 1 поле VersionID в базе данных
func (crud Crud_DB) Update_VersionID(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
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
func Update_VersionID_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) || (m.VersionID == 0) {
		err = errors.New(m.TableNameDB() + ` Update_VersionID() error: Oid=0`)
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
	cache.Remove(postgres_migrate_pg_class.StringIdentifier(m.Oid, m.VersionID))

	return err
}
