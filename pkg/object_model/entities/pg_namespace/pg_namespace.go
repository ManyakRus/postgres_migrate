package pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_pg_namespace"
)

// PgNamespace - модель для таблицы pg_namespace: В pg_namespace сохраняются пространства имён. Пространство имён представляет собой структуру, на которой основываются схемы SQL: в каждом пространстве имён без конфликтов может существовать отдельный набор отношений, типов и т. д.
type PgNamespace struct {
	table_pg_namespace.Table_PgNamespace	
}
