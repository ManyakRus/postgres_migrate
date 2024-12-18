//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_class"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_postgres_migrate_pg_class.ICrud_ObjectPostgresMigratePgClass) {
	object_postgres_migrate_pg_class.Crud_ObjectPostgresMigratePgClass = crud

	return
}
