package crud_starter_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_class"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud pg_class.ICrud_manual_PgClass) {
	pg_class.Crud_manual_PgClass = crud

	return
}
