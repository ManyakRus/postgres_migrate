//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_namespace"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud pg_namespace.ICrud_PgNamespace) {
	pg_namespace.Crud_PgNamespace = crud

	return
}
