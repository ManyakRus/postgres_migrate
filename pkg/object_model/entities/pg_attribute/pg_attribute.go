package pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_pg_attribute"
)

// PgAttribute - модель для таблицы pg_attribute: В каталоге pg_attribute хранится информация о столбцах таблицы. Для каждого столбца каждой таблицы в pg_attribute существует ровно одна строка.
type PgAttribute struct {
	table_pg_attribute.Table_PgAttribute	
}
