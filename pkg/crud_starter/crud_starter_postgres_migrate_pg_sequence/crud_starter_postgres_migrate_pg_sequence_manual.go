package crud_starter_postgres_migrate_pg_sequence

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud postgres_migrate_pg_sequence.ICrud_manual_PostgresMigratePgSequence) {
	postgres_migrate_pg_sequence.Crud_manual_PostgresMigratePgSequence = crud

	return
}
