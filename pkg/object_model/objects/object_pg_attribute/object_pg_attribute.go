package object_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
)

// ObjectPgAttribute - объект для таблицы pg_attribute: В каталоге pg_attribute хранится информация о столбцах таблицы. Для каждого столбца каждой таблицы в pg_attribute существует ровно одна строка.
type ObjectPgAttribute struct {
	pg_attribute.PgAttribute
	Version version.Version	`json:"version"	gorm:"-:all"`
}
