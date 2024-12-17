//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_class"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud pg_class.ICrud_PgClass) {
	pg_class.Crud_PgClass = crud

	return
}
