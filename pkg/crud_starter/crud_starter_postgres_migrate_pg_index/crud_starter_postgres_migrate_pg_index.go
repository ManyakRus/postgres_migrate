//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud postgres_migrate_pg_index.ICrud_PostgresMigratePgIndex) {
	postgres_migrate_pg_index.Crud_PostgresMigratePgIndex = crud

	return
}