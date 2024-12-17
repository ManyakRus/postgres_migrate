//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_description"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_pg_description.ICrud_ObjectPgDescription) {
	object_pg_description.Crud_ObjectPgDescription = crud

	return
}
