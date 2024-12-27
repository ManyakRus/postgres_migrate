package object_postgres_migrate_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
)

// ObjectPostgresMigratePgNamespace - объект для таблицы postgres_migrate_pg_namespace: В postgres_migrate_pg_namespace сохраняются пространства имён. Пространство имён представляет собой структуру, на которой основываются схемы SQL: в каждом пространстве имён без конфликтов может существовать отдельный набор отношений, типов и т. д.
type ObjectPostgresMigratePgNamespace struct {
	postgres_migrate_pg_namespace.PostgresMigratePgNamespace
	Version postgres_migrate_version.PostgresMigrateVersion	`json:"version"	gorm:"-:all"`
}
