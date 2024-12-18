//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud postgres_migrate_pg_class.ICrud_PostgresMigratePgClass) {
	postgres_migrate_pg_class.Crud_PostgresMigratePgClass = crud

	return
}
