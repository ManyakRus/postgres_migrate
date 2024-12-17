package crud_starter_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_namespace"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud pg_namespace.ICrud_manual_PgNamespace) {
	pg_namespace.Crud_manual_PgNamespace = crud

	return
}
