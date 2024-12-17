package crud_starter_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_attribute"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud pg_attribute.ICrud_manual_PgAttribute) {
	pg_attribute.Crud_manual_PgAttribute = crud

	return
}
