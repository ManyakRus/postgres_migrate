//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_index"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_postgres_migrate_pg_index.ICrud_ObjectPostgresMigratePgIndex) {
	object_postgres_migrate_pg_index.Crud_ObjectPostgresMigratePgIndex = crud

	return
}
