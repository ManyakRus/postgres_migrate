//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_attribute"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_pg_attribute.ICrud_ObjectPgAttribute) {
	object_pg_attribute.Crud_ObjectPgAttribute = crud

	return
}
