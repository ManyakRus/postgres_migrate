//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_version"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_version.ICrud_ObjectVersion) {
	object_version.Crud_ObjectVersion = crud

	return
}
