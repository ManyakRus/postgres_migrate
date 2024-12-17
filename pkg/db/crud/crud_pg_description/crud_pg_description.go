//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_functions"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_description"
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
const TableName string = "pg_description"

// Crud_DB - объект для CRUD операций через БД
type Crud_DB struct {
}

// Read - находит запись в БД по ID
func (crud Crud_DB) Read(m *pg_description.PgDescription) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = Read_ctx(ctx, db, m)
	return err
}

// Read_ctx - находит запись в БД по ID
func Read_ctx(ctx context.Context, db *gorm.DB, m *pg_description.PgDescription) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	tx := db.WithContext(ctx)

	tx = tx.Where(m, "Classoid", "Objoid", "Objsubid", "VersionID").Take(m)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Read() Classoid: %v, error: %w", m.Classoid, err)
	}

	return err
}

// Save - записывает новый или существующий объект в базу данных
func (crud Crud_DB) Save(m *pg_description.PgDescription) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err := Save_ctx(ctx, db, m)
	return err
}

// Save_ctx - записывает новый или существующий объект в базу данных
func Save_ctx(ctx context.Context, db *gorm.DB, m *pg_description.PgDescription) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	err = create_update_ctx(ctx, db, m, nil)
	return err
}

// Update - записывает существующий объект в базу данных
func (crud Crud_DB) Update(m *pg_description.PgDescription) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err := Update_ctx(ctx, db, m)
	return err
}

// Update_ctx - записывает существующий объект в базу данных
func Update_ctx(ctx context.Context, db *gorm.DB, m *pg_description.PgDescription) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	// проверка ID
	if m.Classoid == 0 || m.Objoid == 0 || m.Objsubid == 0 || m.VersionID == 0 {
		TextError := fmt.Sprint(m.TableNameDB()+" Update() ", TableName, " error: Classoid =0")
		err = errors.New(TextError)
		return err
	}

	err = create_update_ctx(ctx, db, m, nil)
	return err
}

// Create - записывает новый объект в базу данных
func (crud Crud_DB) Create(m *pg_description.PgDescription) error {
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err := Create_ctx(ctx, db, m)
	return err
}

// Create_ctx - записывает новый объект в базу данных
func Create_ctx(ctx context.Context, db *gorm.DB, m *pg_description.PgDescription) error {
	var err error
	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	// проверка ID
	if m.Classoid != 0 || m.Objoid != 0 || m.Objsubid != 0 || m.VersionID != 0 {
		TextError := fmt.Sprint(m.TableNameDB()+" Create() ", TableName, " error: Classoid !=0")
		err = errors.New(TextError)
		return err
	}

	err = create_update_ctx(ctx, db, m, nil)
	return err
}

// create_update - записывает объект в базу данных
func (crud Crud_DB) create_update(m *pg_description.PgDescription) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = create_update_ctx(ctx, db, m, nil)
	return err
}

// create_update_ctx - записывает объект в базу данных
func create_update_ctx(ctx context.Context, db *gorm.DB, m *pg_description.PgDescription, MassNeedUpdateFields []string) error {
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
	cache.Remove(pg_description.StringIdentifier(m.Classoid, m.Objoid, m.Objsubid, m.VersionID))

	//запишем NULL в пустые колонки
	MapOmit := crud_functions.MapOmit_from_MassOmit(MassOmit)
	tx = db.Model(&m)
	if len(MassNeedUpdateFields) > 0 {
		tx = tx.Select(MassNeedFields)
	}
	tx = tx.Updates(MapOmit)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Create_Update() Classoid: %v, error: %w", m.Classoid, err)
	}

	return err
}
