//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_version"
	"testing"
)

func TestSetCrudInterface(t *testing.T) {
	crud := crud_version.Crud_DB{}

	SetCrudInterface(crud)
}
