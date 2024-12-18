package crud_starter_postgres_migrate_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud postgres_migrate_pg_description.ICrud_manual_PostgresMigratePgDescription) {
	postgres_migrate_pg_description.Crud_manual_PostgresMigratePgDescription = crud

	return
}
