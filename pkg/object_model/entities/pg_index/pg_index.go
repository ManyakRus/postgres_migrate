package pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_pg_index"
)

// PgIndex - модель для таблицы pg_index: В каталоге pg_index содержится часть информации об индексах. Остальная информация в основном находится в pg_class.
type PgIndex struct {
	table_pg_index.Table_PgIndex	
}
