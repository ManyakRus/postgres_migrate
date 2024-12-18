//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_constraint"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud postgres_migrate_pg_constraint.ICrud_PostgresMigratePgConstraint) {
	postgres_migrate_pg_constraint.Crud_PostgresMigratePgConstraint = crud

	return
}
