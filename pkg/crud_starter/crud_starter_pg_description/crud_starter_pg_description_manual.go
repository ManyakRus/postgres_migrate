package crud_starter_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_description"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud pg_description.ICrud_manual_PgDescription) {
	pg_description.Crud_manual_PgDescription = crud

	return
}
