//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_description"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud pg_description.ICrud_PgDescription) {
	pg_description.Crud_PgDescription = crud

	return
}
