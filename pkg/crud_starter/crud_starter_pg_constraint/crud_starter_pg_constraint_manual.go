package crud_starter_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_constraint"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud pg_constraint.ICrud_manual_PgConstraint) {
	pg_constraint.Crud_manual_PgConstraint = crud

	return
}
