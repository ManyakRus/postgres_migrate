package crud_starter_postgres_migrate_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud postgres_migrate_pg_attribute.ICrud_manual_PostgresMigratePgAttribute) {
	postgres_migrate_pg_attribute.Crud_manual_PostgresMigratePgAttribute = crud

	return
}
