//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_object_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_version"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_namespace"
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
func (crud Crud_DB) ReadObject(m *object_pg_namespace.ObjectPgNamespace) error {
	var err error

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()

	err = ReadObject_ctx(ctx, db, m)
	return err
}

// ReadObject_ctx - находит запись в БД по ID, также заполняет внешние поля
func ReadObject_ctx(ctx context.Context, db *gorm.DB, m *object_pg_namespace.ObjectPgNamespace) error {
	var err error

	if micro.ContextDone(ctx) == true {
		err = context.Canceled
		return err
	}

	//заполним model
	Model := m.PgNamespace
	err = crud_pg_namespace.Read_ctx(ctx, db, &Model)
	if err != nil {
		err = fmt.Errorf(m.TableNameDB()+" Read() Oid: %v, VersionID: %v, error: %w", m.Oid, m.VersionID, err)
	}
	m.PgNamespace = Model

	//VersionID
	Version:= version.Version{}
	VersionID := m.VersionID
	Version.ID = VersionID
	err = crud_version.Read_ctx(ctx, db, &Version)
	if err != nil {
		err = fmt.Errorf(Version.TableNameDB()+" Read() Oid: %v, VersionID: %v, error: %w", m.Oid, m.VersionID, err)
	}
	m.Version = Version

	return err
}
