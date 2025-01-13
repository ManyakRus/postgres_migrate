//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_sequence

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud postgres_migrate_pg_sequence.ICrud_PostgresMigratePgSequence) {
	postgres_migrate_pg_sequence.Crud_PostgresMigratePgSequence = crud

	return
}
