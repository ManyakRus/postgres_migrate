//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_pg_constraint

import (
	"fmt"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_constraint"
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
func (crud Crud_DB) UpdateManyFields(m *pg_constraint.PgConstraint, MassNeedUpdateFields []string) error {
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
func UpdateManyFields_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint, MassNeedUpdateFields []string) error {
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

// Update_Condeferrable - изменяет 1 поле Condeferrable в базе данных
func (crud Crud_DB) Update_Condeferrable(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Condeferrable_ctx(ctx, db, m)
	return err
}

// Update_Condeferrable_ctx - изменяет 1 поле Condeferrable в базе данных
// с учётом контекста и соединения к БД
func Update_Condeferrable_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Condeferrable() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Condeferrable
	tx = db.Model(&m).Update("Condeferrable", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Condeferrable() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Condeferred - изменяет 1 поле Condeferred в базе данных
func (crud Crud_DB) Update_Condeferred(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Condeferred_ctx(ctx, db, m)
	return err
}

// Update_Condeferred_ctx - изменяет 1 поле Condeferred в базе данных
// с учётом контекста и соединения к БД
func Update_Condeferred_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Condeferred() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Condeferred
	tx = db.Model(&m).Update("Condeferred", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Condeferred() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conexclop - изменяет 1 поле Conexclop в базе данных
func (crud Crud_DB) Update_Conexclop(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conexclop_ctx(ctx, db, m)
	return err
}

// Update_Conexclop_ctx - изменяет 1 поле Conexclop в базе данных
// с учётом контекста и соединения к БД
func Update_Conexclop_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conexclop() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conexclop
	tx = db.Model(&m).Update("Conexclop", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conexclop() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Confdeltype - изменяет 1 поле Confdeltype в базе данных
func (crud Crud_DB) Update_Confdeltype(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Confdeltype_ctx(ctx, db, m)
	return err
}

// Update_Confdeltype_ctx - изменяет 1 поле Confdeltype в базе данных
// с учётом контекста и соединения к БД
func Update_Confdeltype_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Confdeltype() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Confdeltype
	tx = db.Model(&m).Update("Confdeltype", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Confdeltype() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conffeqop - изменяет 1 поле Conffeqop в базе данных
func (crud Crud_DB) Update_Conffeqop(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conffeqop_ctx(ctx, db, m)
	return err
}

// Update_Conffeqop_ctx - изменяет 1 поле Conffeqop в базе данных
// с учётом контекста и соединения к БД
func Update_Conffeqop_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conffeqop() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conffeqop
	tx = db.Model(&m).Update("Conffeqop", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conffeqop() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Confkey - изменяет 1 поле Confkey в базе данных
func (crud Crud_DB) Update_Confkey(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Confkey_ctx(ctx, db, m)
	return err
}

// Update_Confkey_ctx - изменяет 1 поле Confkey в базе данных
// с учётом контекста и соединения к БД
func Update_Confkey_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Confkey() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Confkey
	tx = db.Model(&m).Update("Confkey", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Confkey() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Confmatchtype - изменяет 1 поле Confmatchtype в базе данных
func (crud Crud_DB) Update_Confmatchtype(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Confmatchtype_ctx(ctx, db, m)
	return err
}

// Update_Confmatchtype_ctx - изменяет 1 поле Confmatchtype в базе данных
// с учётом контекста и соединения к БД
func Update_Confmatchtype_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Confmatchtype() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Confmatchtype
	tx = db.Model(&m).Update("Confmatchtype", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Confmatchtype() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Confrelid - изменяет 1 поле Confrelid в базе данных
func (crud Crud_DB) Update_Confrelid(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Confrelid_ctx(ctx, db, m)
	return err
}

// Update_Confrelid_ctx - изменяет 1 поле Confrelid в базе данных
// с учётом контекста и соединения к БД
func Update_Confrelid_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Confrelid() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Confrelid
	tx = db.Model(&m).Update("Confrelid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Confrelid() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Confupdtype - изменяет 1 поле Confupdtype в базе данных
func (crud Crud_DB) Update_Confupdtype(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Confupdtype_ctx(ctx, db, m)
	return err
}

// Update_Confupdtype_ctx - изменяет 1 поле Confupdtype в базе данных
// с учётом контекста и соединения к БД
func Update_Confupdtype_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Confupdtype() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Confupdtype
	tx = db.Model(&m).Update("Confupdtype", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Confupdtype() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conindid - изменяет 1 поле Conindid в базе данных
func (crud Crud_DB) Update_Conindid(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conindid_ctx(ctx, db, m)
	return err
}

// Update_Conindid_ctx - изменяет 1 поле Conindid в базе данных
// с учётом контекста и соединения к БД
func Update_Conindid_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conindid() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conindid
	tx = db.Model(&m).Update("Conindid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conindid() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Coninhcount - изменяет 1 поле Coninhcount в базе данных
func (crud Crud_DB) Update_Coninhcount(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Coninhcount_ctx(ctx, db, m)
	return err
}

// Update_Coninhcount_ctx - изменяет 1 поле Coninhcount в базе данных
// с учётом контекста и соединения к БД
func Update_Coninhcount_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Coninhcount() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Coninhcount
	tx = db.Model(&m).Update("Coninhcount", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Coninhcount() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conislocal - изменяет 1 поле Conislocal в базе данных
func (crud Crud_DB) Update_Conislocal(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conislocal_ctx(ctx, db, m)
	return err
}

// Update_Conislocal_ctx - изменяет 1 поле Conislocal в базе данных
// с учётом контекста и соединения к БД
func Update_Conislocal_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conislocal() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conislocal
	tx = db.Model(&m).Update("Conislocal", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conislocal() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conkey - изменяет 1 поле Conkey в базе данных
func (crud Crud_DB) Update_Conkey(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conkey_ctx(ctx, db, m)
	return err
}

// Update_Conkey_ctx - изменяет 1 поле Conkey в базе данных
// с учётом контекста и соединения к БД
func Update_Conkey_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conkey() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conkey
	tx = db.Model(&m).Update("Conkey", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conkey() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conname - изменяет 1 поле Conname в базе данных
func (crud Crud_DB) Update_Conname(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conname_ctx(ctx, db, m)
	return err
}

// Update_Conname_ctx - изменяет 1 поле Conname в базе данных
// с учётом контекста и соединения к БД
func Update_Conname_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conname() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conname
	tx = db.Model(&m).Update("Conname", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conname() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Connamespace - изменяет 1 поле Connamespace в базе данных
func (crud Crud_DB) Update_Connamespace(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Connamespace_ctx(ctx, db, m)
	return err
}

// Update_Connamespace_ctx - изменяет 1 поле Connamespace в базе данных
// с учётом контекста и соединения к БД
func Update_Connamespace_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Connamespace() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Connamespace
	tx = db.Model(&m).Update("Connamespace", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Connamespace() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Connoinherit - изменяет 1 поле Connoinherit в базе данных
func (crud Crud_DB) Update_Connoinherit(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Connoinherit_ctx(ctx, db, m)
	return err
}

// Update_Connoinherit_ctx - изменяет 1 поле Connoinherit в базе данных
// с учётом контекста и соединения к БД
func Update_Connoinherit_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Connoinherit() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Connoinherit
	tx = db.Model(&m).Update("Connoinherit", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Connoinherit() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conparentid - изменяет 1 поле Conparentid в базе данных
func (crud Crud_DB) Update_Conparentid(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conparentid_ctx(ctx, db, m)
	return err
}

// Update_Conparentid_ctx - изменяет 1 поле Conparentid в базе данных
// с учётом контекста и соединения к БД
func Update_Conparentid_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conparentid() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conparentid
	tx = db.Model(&m).Update("Conparentid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conparentid() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conpfeqop - изменяет 1 поле Conpfeqop в базе данных
func (crud Crud_DB) Update_Conpfeqop(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conpfeqop_ctx(ctx, db, m)
	return err
}

// Update_Conpfeqop_ctx - изменяет 1 поле Conpfeqop в базе данных
// с учётом контекста и соединения к БД
func Update_Conpfeqop_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conpfeqop() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conpfeqop
	tx = db.Model(&m).Update("Conpfeqop", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conpfeqop() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conppeqop - изменяет 1 поле Conppeqop в базе данных
func (crud Crud_DB) Update_Conppeqop(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conppeqop_ctx(ctx, db, m)
	return err
}

// Update_Conppeqop_ctx - изменяет 1 поле Conppeqop в базе данных
// с учётом контекста и соединения к БД
func Update_Conppeqop_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conppeqop() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conppeqop
	tx = db.Model(&m).Update("Conppeqop", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conppeqop() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Conrelid - изменяет 1 поле Conrelid в базе данных
func (crud Crud_DB) Update_Conrelid(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Conrelid_ctx(ctx, db, m)
	return err
}

// Update_Conrelid_ctx - изменяет 1 поле Conrelid в базе данных
// с учётом контекста и соединения к БД
func Update_Conrelid_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Conrelid() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Conrelid
	tx = db.Model(&m).Update("Conrelid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Conrelid() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Contype - изменяет 1 поле Contype в базе данных
func (crud Crud_DB) Update_Contype(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Contype_ctx(ctx, db, m)
	return err
}

// Update_Contype_ctx - изменяет 1 поле Contype в базе данных
// с учётом контекста и соединения к БД
func Update_Contype_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Contype() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Contype
	tx = db.Model(&m).Update("Contype", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Contype() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Contypid - изменяет 1 поле Contypid в базе данных
func (crud Crud_DB) Update_Contypid(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Contypid_ctx(ctx, db, m)
	return err
}

// Update_Contypid_ctx - изменяет 1 поле Contypid в базе данных
// с учётом контекста и соединения к БД
func Update_Contypid_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Contypid() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Contypid
	tx = db.Model(&m).Update("Contypid", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Contypid() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Convalidated - изменяет 1 поле Convalidated в базе данных
func (crud Crud_DB) Update_Convalidated(m *pg_constraint.PgConstraint) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Update_Convalidated_ctx(ctx, db, m)
	return err
}

// Update_Convalidated_ctx - изменяет 1 поле Convalidated в базе данных
// с учётом контекста и соединения к БД
func Update_Convalidated_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//ID не должен быть =0
	if (m.Oid == 0) ||  (m.VersionID == 0) {
		err = errors.New(m.TableNameDB()+` Update_Convalidated() error: Oid=0`)
		return err
	}

	//
	tx := db.WithContext(ctx)

	//
	Value := m.Convalidated
	tx = db.Model(&m).Update("Convalidated", Value)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Update_Convalidated() id: %v, error: %w", m.Oid, err)
		return err
	}

	//удалим из кэша
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_Oid - изменяет 1 поле Oid в базе данных
func (crud Crud_DB) Update_Oid(m *pg_constraint.PgConstraint) error {
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
func Update_Oid_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
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
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}

// Update_VersionID - изменяет 1 поле VersionID в базе данных
func (crud Crud_DB) Update_VersionID(m *pg_constraint.PgConstraint) error {
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
func Update_VersionID_ctx(ctx context.Context, db *gorm.DB, m *pg_constraint.PgConstraint) error {
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
	cache.Remove(pg_constraint.StringIdentifier(m.Oid, m.VersionID))

	return err
}
