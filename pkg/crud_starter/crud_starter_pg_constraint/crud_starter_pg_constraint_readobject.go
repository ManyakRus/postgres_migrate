//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_constraint"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_pg_constraint.ICrud_ObjectPgConstraint) {
	object_pg_constraint.Crud_ObjectPgConstraint = crud

	return
}
