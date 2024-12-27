package object_postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
)

// ObjectPostgresMigrateVersion - объект для таблицы postgres_migrate_version: 
type ObjectPostgresMigrateVersion struct {
	postgres_migrate_version.PostgresMigrateVersion
}
