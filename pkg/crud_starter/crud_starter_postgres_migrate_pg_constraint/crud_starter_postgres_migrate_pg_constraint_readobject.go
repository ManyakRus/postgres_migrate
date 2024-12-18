//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_constraint"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_postgres_migrate_pg_constraint.ICrud_ObjectPostgresMigratePgConstraint) {
	object_postgres_migrate_pg_constraint.Crud_ObjectPostgresMigratePgConstraint = crud

	return
}
