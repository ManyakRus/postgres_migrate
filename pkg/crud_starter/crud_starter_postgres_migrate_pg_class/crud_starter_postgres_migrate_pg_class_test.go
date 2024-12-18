//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_class"
	"testing"
)

func TestSetCrudInterface(t *testing.T) {
	crud := crud_postgres_migrate_pg_class.Crud_DB{}

	SetCrudInterface(crud)
}
