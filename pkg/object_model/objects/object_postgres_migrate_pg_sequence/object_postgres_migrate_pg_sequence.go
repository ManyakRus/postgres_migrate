package object_postgres_migrate_pg_sequence

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
)

// ObjectPostgresMigratePgSequence - объект для таблицы postgres_migrate_pg_sequence:
type ObjectPostgresMigratePgSequence struct {
	postgres_migrate_pg_sequence.PostgresMigratePgSequence
	Version postgres_migrate_version.PostgresMigrateVersion `json:"version"	gorm:"-:all"`
}
