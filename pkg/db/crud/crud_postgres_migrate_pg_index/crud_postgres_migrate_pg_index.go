//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_functions"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"context"
	"errors"
	"fmt"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"time"
	"github.com/ManyakRus/starter/contextmain"
	"gorm.io/gorm"
)

// TableName - имя таблицы в БД Postgres
const TableName string = "postgres_migrate_pg_index"

// Crud_DB - объект для CRUD операций через БД
type Crud_DB struct {
}

// Read - находит запись в БД по ID
func (crud Crud_DB) Read(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Read_ctx(ctx, db, m)
	return err
}

// Read_ctx - находит запись в БД по ID
func Read_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	tx := db.WithContext(ctx)

	tx = tx.Where(m, "Indexrelid", "VersionID").Take(m)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Read() Indexrelid: %v, error: %w", m.Indexrelid, err)
	}

	return err
}

// Save - записывает новый или существующий объект в базу данных
func (crud Crud_DB) Save(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err := Save_ctx(ctx, db, m)
	return err
}

// Save_ctx - записывает новый или существующий объект в базу данных
func Save_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	err = create_update_ctx(ctx, db, m, nil)
	return err
}

// Update - записывает существующий объект в базу данных
func (crud Crud_DB) Update(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err := Update_ctx(ctx, db, m)
	return err
}

// Update_ctx - записывает существующий объект в базу данных
func Update_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	// проверка ID
	if m.Indexrelid == 0 || m.VersionID == 0 {
		TextError := fmt.Sprint(m.TableNameDB()+" Update() ", TableName, " error: Indexrelid =0")
		err = errors.New(TextError)
		return err
	}

	err = create_update_ctx(ctx, db, m, nil)
	return err
}

// Create - записывает новый объект в базу данных
func (crud Crud_DB) Create(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err := Create_ctx(ctx, db, m)
	return err
}

// Create_ctx - записывает новый объект в базу данных
func Create_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	// проверка ID
	if m.Indexrelid != 0 || m.VersionID != 0 {
		TextError := fmt.Sprint(m.TableNameDB()+" Create() ", TableName, " error: Indexrelid !=0")
		err = errors.New(TextError)
		return err
	}

	err = create_update_ctx(ctx, db, m, nil)
	return err
}

// create_update - записывает объект в базу данных
func (crud Crud_DB) create_update(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = create_update_ctx(ctx, db, m, nil)
	return err
}

// create_update_ctx - записывает объект в базу данных
func create_update_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex, MassNeedUpdateFields []string) error {
	var err error

	// log.Trace("start Save() ", TableName, " id: ", int64(m.ID))

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//
	tx := db.WithContext(ctx).Model(&m)

	//колонки для обновления
	MassNeedFields := crud_functions.MassNeedFields_from_MassNeedUpdateFields(MassNeedUpdateFields)
	if len(MassNeedUpdateFields) > 0 {
		tx = tx.Select(MassNeedFields)
	}

	//колонки с null
	MassOmit := make([]string, 0)

	//игнор пустых колонок
	tx = tx.Omit(MassOmit...)

	// запись
	tx = tx.Save(&m)
	err = tx.Error
	if err != nil {
		return err
	}

	//удалим из кэша
	cache.Remove(postgres_migrate_pg_index.StringIdentifier(m.Indexrelid, m.VersionID))

	//запишем NULL в пустые колонки
	MapOmit := crud_functions.MapOmit_from_MassOmit(MassOmit)
	tx = db.Model(&m)
	if len(MassNeedUpdateFields) > 0 {
		tx = tx.Select(MassNeedFields)
	}
	tx = tx.Updates(MapOmit)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Create_Update() Indexrelid: %v, error: %w", m.Indexrelid, err)
	}

	return err
}

// Delete - записывает is_deleted = true
func (crud Crud_DB) Delete(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Delete_ctx(ctx, db, m)
	return err
}

// Delete_ctx - записывает is_deleted = true
func Delete_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	m2 := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m2.Indexrelid = m.Indexrelid
	m2.VersionID = m.VersionID

	err = Read_ctx(ctx, db, &m2)
	if err != nil {
		return err
	}

	m2.IsDeleted = true
	m.IsDeleted = true

	err = Save_ctx(ctx, db, &m2)
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Delete() Indexrelid: %v, error: %w", m.Indexrelid, err)
	}

	return err
}

// Restore - записывает is_deleted = true
func (crud Crud_DB) Restore(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Restore_ctx(ctx, db, m)
	return err
}

// Restore_ctx - записывает is_deleted = true
func Restore_ctx(ctx context.Context, db *gorm.DB, m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	m2 := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m2.Indexrelid = m.Indexrelid
	m2.VersionID = m.VersionID

	err = Read_ctx(ctx, db, &m2)
	if err != nil {
		return err
	}

	m2.IsDeleted = false
	m.IsDeleted = false

	err = Save_ctx(ctx, db, &m2)
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Restore() Indexrelid: %v, error: %w", m.Indexrelid, err)
	}

	return err
}
