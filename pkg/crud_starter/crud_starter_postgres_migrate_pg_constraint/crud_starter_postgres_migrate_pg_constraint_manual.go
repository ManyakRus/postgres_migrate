package crud_starter_postgres_migrate_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_constraint"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud postgres_migrate_pg_constraint.ICrud_manual_PostgresMigratePgConstraint) {
	postgres_migrate_pg_constraint.Crud_manual_PostgresMigratePgConstraint = crud

	return
}
