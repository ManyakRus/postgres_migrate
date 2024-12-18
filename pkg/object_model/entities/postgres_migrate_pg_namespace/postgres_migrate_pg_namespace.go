package postgres_migrate_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_postgres_migrate_pg_namespace"
)

// PostgresMigratePgNamespace - модель для таблицы postgres_migrate_pg_namespace: В postgres_migrate_pg_namespace сохраняются пространства имён. Пространство имён представляет собой структуру, на которой основываются схемы SQL: в каждом пространстве имён без конфликтов может существовать отдельный набор отношений, типов и т. д.
type PostgresMigratePgNamespace struct {
	table_postgres_migrate_pg_namespace.Table_PostgresMigratePgNamespace
}
