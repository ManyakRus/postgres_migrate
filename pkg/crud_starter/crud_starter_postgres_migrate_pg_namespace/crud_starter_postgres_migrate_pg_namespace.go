//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_namespace"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud postgres_migrate_pg_namespace.ICrud_PostgresMigratePgNamespace) {
	postgres_migrate_pg_namespace.Crud_PostgresMigratePgNamespace = crud

	return
}
