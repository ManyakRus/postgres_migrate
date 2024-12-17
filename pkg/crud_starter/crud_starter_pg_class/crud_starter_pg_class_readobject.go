//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_class"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_pg_class.ICrud_ObjectPgClass) {
	object_pg_class.Crud_ObjectPgClass = crud

	return
}
