//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_attribute"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudReadObjectInterface(crud object_postgres_migrate_pg_attribute.ICrud_ObjectPostgresMigratePgAttribute) {
	object_postgres_migrate_pg_attribute.Crud_ObjectPostgresMigratePgAttribute = crud

	return
}
