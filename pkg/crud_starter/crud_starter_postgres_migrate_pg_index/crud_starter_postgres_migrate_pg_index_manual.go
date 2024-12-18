package crud_starter_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud postgres_migrate_pg_index.ICrud_manual_PostgresMigratePgIndex) {
	postgres_migrate_pg_index.Crud_manual_PostgresMigratePgIndex = crud

	return
}
