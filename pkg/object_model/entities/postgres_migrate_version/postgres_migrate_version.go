package postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_postgres_migrate_version"
)

// PostgresMigrateVersion - модель для таблицы postgres_migrate_version:
type PostgresMigrateVersion struct {
	table_postgres_migrate_version.Table_PostgresMigrateVersion
}
