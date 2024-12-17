package object_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
)

// ObjectPgConstraint - объект для таблицы pg_constraint: В каталоге pg_constraint хранятся ограничения-проверки, ограничения-исключения, а также ограничения первичного ключа, уникальности и внешних ключей, определённые для таблиц. (Ограничения столбцов описываются как и все остальные. Любое ограничение столбца равнозначно некоторому ограничению таблицы.) Ограничения на NULL представляются не здесь, а в каталоге pg_attribute.Для пользовательских триггеров ограничений (создаваемых командой CREATE CONSTRAINT TRIGGER) в этой таблице также создаётся запись.Здесь также хранятся ограничения доменов.
type ObjectPgConstraint struct {
	pg_constraint.PgConstraint
	Version version.Version	`json:"version"	gorm:"-:all"`
}
