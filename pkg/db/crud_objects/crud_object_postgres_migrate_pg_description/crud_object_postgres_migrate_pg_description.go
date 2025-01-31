//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_object_postgres_migrate_pg_description

import (
	"context"
	"fmt"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_description"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"gorm.io/gorm"
	"time"
)

// Crud_DB - объект для CRUD операций через БД
type Crud_DB struct {
}

// ReadObject - находит запись в БД по ID, также заполняет внешние поля
func (crud Crud_DB) ReadObject(m *object_postgres_migrate_pg_description.ObjectPostgresMigratePgDescription) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = ReadObject_ctx(ctx, db, m)
	return err
}

// ReadObject_ctx - находит запись в БД по ID, также заполняет внешние поля
func ReadObject_ctx(ctx context.Context, db *gorm.DB, m *object_postgres_migrate_pg_description.ObjectPostgresMigratePgDescription) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//заполним model
	Model := m.PostgresMigratePgDescription
	err = crud_postgres_migrate_pg_description.Read_ctx(ctx, db, &Model)
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Read() Classoid: %v, Objoid: %v, Objsubid: %v, VersionID: %v, error: %w", m.Classoid, m.Objoid, m.Objsubid, m.VersionID, err)
	}
	m.PostgresMigratePgDescription = Model

	return err
}
