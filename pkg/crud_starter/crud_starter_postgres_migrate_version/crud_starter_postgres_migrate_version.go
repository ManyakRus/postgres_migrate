//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter_postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
)

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudInterface(crud postgres_migrate_version.ICrud_PostgresMigrateVersion) {
	postgres_migrate_version.Crud_PostgresMigrateVersion = crud

	return
}
