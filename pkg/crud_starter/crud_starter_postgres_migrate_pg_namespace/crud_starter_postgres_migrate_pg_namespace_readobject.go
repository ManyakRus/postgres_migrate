//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_namespace"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_postgres_migrate_pg_namespace.ICrud_ObjectPostgresMigratePgNamespace) {
	object_postgres_migrate_pg_namespace.Crud_ObjectPostgresMigratePgNamespace = crud

	return
}
