package crud_starter_postgres_migrate_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud postgres_migrate_pg_class.ICrud_manual_PostgresMigratePgClass) {
	postgres_migrate_pg_class.Crud_manual_PostgresMigratePgClass = crud

	return
}
