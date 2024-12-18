package postgres_migrate_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_postgres_migrate_pg_description"
)

// PostgresMigratePgDescription - модель для таблицы postgres_migrate_pg_description: В каталоге postgres_migrate_pg_description хранятся дополнительные описания (комментарии) для объектов баз данных. Описания можно задавать с помощью команды COMMENT и просматривать в psql, используя команды \d. В начальном содержимом postgres_migrate_pg_description представлены описания многих встроенных системных объектов.Также смотрите каталог pg_shdescription, который играет подобную роль в отношении совместно используемых объектов в кластере баз данных.
type PostgresMigratePgDescription struct {
	table_postgres_migrate_pg_description.Table_PostgresMigratePgDescription
}
