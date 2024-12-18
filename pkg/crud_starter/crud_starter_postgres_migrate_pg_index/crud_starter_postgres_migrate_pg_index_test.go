//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_index"
	"testing"
)

func TestSetCrudInterface(t *testing.T) {
	crud := crud_postgres_migrate_pg_index.Crud_DB{}

	SetCrudInterface(crud)
}
