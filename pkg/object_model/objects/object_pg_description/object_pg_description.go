package object_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_description"
)

// ObjectPgDescription - объект для таблицы pg_description: В каталоге pg_description хранятся дополнительные описания (комментарии) для объектов баз данных. Описания можно задавать с помощью команды COMMENT и просматривать в psql, используя команды \d. В начальном содержимом pg_description представлены описания многих встроенных системных объектов.Также смотрите каталог pg_shdescription, который играет подобную роль в отношении совместно используемых объектов в кластере баз данных.
type ObjectPgDescription struct {
	pg_description.PgDescription
}
