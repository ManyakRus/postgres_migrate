package object_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
)

// ObjectPgIndex - объект для таблицы pg_index: В каталоге pg_index содержится часть информации об индексах. Остальная информация в основном находится в pg_class.
type ObjectPgIndex struct {
	pg_index.PgIndex
	Version version.Version	`json:"version"	gorm:"-:all"`
}
