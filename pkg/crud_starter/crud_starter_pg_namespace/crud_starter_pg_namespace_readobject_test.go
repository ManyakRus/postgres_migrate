//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_pg_namespace"
	"testing"
)

func TestSetCrudInterface(t *testing.T) {
	crud := crud_object_pg_namespace.Crud_DB{}

	SetCrudInterface(crud)
}
