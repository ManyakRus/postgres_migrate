package object_postgres_migrate_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
)

// ObjectPostgresMigratePgAttribute - объект для таблицы postgres_migrate_pg_attribute: В каталоге postgres_migrate_pg_attribute хранится информация о столбцах таблицы. Для каждого столбца каждой таблицы в postgres_migrate_pg_attribute существует ровно одна строка.
type ObjectPostgresMigratePgAttribute struct {
	postgres_migrate_pg_attribute.PostgresMigratePgAttribute
	Version postgres_migrate_version.PostgresMigrateVersion `json:"version"	gorm:"-:all"`
}
