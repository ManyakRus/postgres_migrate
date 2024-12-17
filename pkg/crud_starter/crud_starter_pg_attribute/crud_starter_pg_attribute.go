//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_attribute"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud pg_attribute.ICrud_PgAttribute) {
	pg_attribute.Crud_PgAttribute = crud

	return
}
