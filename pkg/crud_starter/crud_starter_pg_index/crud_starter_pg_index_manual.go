package crud_starter_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_index"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud pg_index.ICrud_manual_PgIndex) {
	pg_index.Crud_manual_PgIndex = crud

	return
}
