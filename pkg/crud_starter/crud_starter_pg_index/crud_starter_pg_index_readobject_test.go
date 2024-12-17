//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_pg_index"
	"testing"
)

func TestSetCrudInterface(t *testing.T) {
	crud := crud_object_pg_index.Crud_DB{}

	SetCrudInterface(crud)
}
