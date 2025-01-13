//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_sequence

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_sequence"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_postgres_migrate_pg_sequence.ICrud_ObjectPostgresMigratePgSequence) {
	object_postgres_migrate_pg_sequence.Crud_ObjectPostgresMigratePgSequence = crud

	return
}
