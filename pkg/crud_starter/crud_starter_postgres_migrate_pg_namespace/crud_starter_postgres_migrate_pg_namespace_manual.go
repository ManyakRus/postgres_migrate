package crud_starter_postgres_migrate_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_namespace"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud postgres_migrate_pg_namespace.ICrud_manual_PostgresMigratePgNamespace) {
	postgres_migrate_pg_namespace.Crud_manual_PostgresMigratePgNamespace = crud

	return
}
