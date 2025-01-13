package postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_postgres_migrate_pg_index"
)

// PostgresMigratePgIndex - модель для таблицы postgres_migrate_pg_index: В каталоге postgres_migrate_pg_index содержится часть информации об индексах. Остальная информация в основном находится в postgres_migrate_pg_class.
type PostgresMigratePgIndex struct {
	table_postgres_migrate_pg_index.Table_PostgresMigratePgIndex
}
