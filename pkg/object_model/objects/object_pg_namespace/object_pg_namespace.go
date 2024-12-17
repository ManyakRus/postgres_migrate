package object_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
)

// ObjectPgNamespace - объект для таблицы pg_namespace: В pg_namespace сохраняются пространства имён. Пространство имён представляет собой структуру, на которой основываются схемы SQL: в каждом пространстве имён без конфликтов может существовать отдельный набор отношений, типов и т. д.
type ObjectPgNamespace struct {
	pg_namespace.PgNamespace
	Version version.Version	`json:"version"	gorm:"-:all"`
}
