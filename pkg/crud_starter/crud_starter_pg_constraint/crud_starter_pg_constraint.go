//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_constraint"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud pg_constraint.ICrud_PgConstraint) {
	pg_constraint.Crud_PgConstraint = crud

	return
}
