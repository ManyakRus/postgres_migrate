package object_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
)

// ObjectPostgresMigratePgIndex - объект для таблицы postgres_migrate_pg_index: В каталоге postgres_migrate_pg_index содержится часть информации об индексах. Остальная информация в основном находится в postgres_migrate_pg_class.
type ObjectPostgresMigratePgIndex struct {
	postgres_migrate_pg_index.PostgresMigratePgIndex
	Version postgres_migrate_version.PostgresMigrateVersion	`json:"version"	gorm:"-:all"`
}
