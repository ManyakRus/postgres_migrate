//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_namespace"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_pg_namespace.ICrud_ObjectPgNamespace) {
	object_pg_namespace.Crud_ObjectPgNamespace = crud

	return
}
