//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_index"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud pg_index.ICrud_PgIndex) {
	pg_index.Crud_PgIndex = crud

	return
}
