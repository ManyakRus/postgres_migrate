//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_index"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_pg_index.ICrud_ObjectPgIndex) {
	object_pg_index.Crud_ObjectPgIndex = crud

	return
}
