//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud version.ICrud_Version) {
	version.Crud_Version = crud

	return
}
