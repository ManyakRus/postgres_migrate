package postgres_migrate_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_postgres_migrate_pg_attribute"
)

// PostgresMigratePgAttribute - модель для таблицы postgres_migrate_pg_attribute: В каталоге postgres_migrate_pg_attribute хранится информация о столбцах таблицы. Для каждого столбца каждой таблицы в postgres_migrate_pg_attribute существует ровно одна строка.
type PostgresMigratePgAttribute struct {
	table_postgres_migrate_pg_attribute.Table_PostgresMigratePgAttribute
}
