package postgres_migrate_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_postgres_migrate_pg_constraint"
)

// PostgresMigratePgConstraint - модель для таблицы postgres_migrate_pg_constraint: В каталоге postgres_migrate_pg_constraint хранятся ограничения-проверки, ограничения-исключения, а также ограничения первичного ключа, уникальности и внешних ключей, определённые для таблиц. (Ограничения столбцов описываются как и все остальные. Любое ограничение столбца равнозначно некоторому ограничению таблицы.) Ограничения на NULL представляются не здесь, а в каталоге postgres_migrate_pg_attribute.Для пользовательских триггеров ограничений (создаваемых командой CREATE CONSTRAINT TRIGGER) в этой таблице также создаётся запись.Здесь также хранятся ограничения доменов.
type PostgresMigratePgConstraint struct {
	table_postgres_migrate_pg_constraint.Table_PostgresMigratePgConstraint	
}
