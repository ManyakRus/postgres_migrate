//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_object_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_version"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_index"
	"context"
	"fmt"
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
func (crud Crud_DB) ReadObject(m *object_pg_index.ObjectPgIndex) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = ReadObject_ctx(ctx, db, m)
	return err
}

// ReadObject_ctx - находит запись в БД по ID, также заполняет внешние поля
func ReadObject_ctx(ctx context.Context, db *gorm.DB, m *object_pg_index.ObjectPgIndex) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//заполним model
	Model := m.PgIndex
	err = crud_pg_index.Read_ctx(ctx, db, &Model)
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Read() Indexrelid: %v, Indrelid: %v, VersionID: %v, error: %w", m.Indexrelid, m.Indrelid, m.VersionID, err)
	}
	m.PgIndex = Model

	//VersionID
	Version:= version.Version{}
	VersionID := m.VersionID
	Version.ID = VersionID
	err = crud_version.Read_ctx(ctx, db, &Version)
	if err != nil {
		err = fmt.Errorf(Version.TableNameDB()+" Read() Indexrelid: %v, Indrelid: %v, VersionID: %v, error: %w", m.Indexrelid, m.Indrelid, m.VersionID, err)
	}
	m.Version = Version

	return err
}
