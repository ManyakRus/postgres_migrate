package crud_starter_postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud postgres_migrate_version.ICrud_manual_PostgresMigrateVersion) {
	postgres_migrate_version.Crud_manual_PostgresMigrateVersion = crud

	return
}
