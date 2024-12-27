package object_postgres_migrate_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
)

// ObjectPostgresMigratePgClass - объект для таблицы postgres_migrate_pg_class: В каталоге postgres_migrate_pg_class описываются таблицы и практически всё, что имеет столбцы или каким-то образом подобно таблице. Сюда входят индексы (но смотрите также postgres_migrate_pg_index), последовательности, представления, материализованные представления, составные типы и таблицы TOAST; см. relkind. Далее, подразумевая все эти типы объектов, мы будем говорить об «отношениях». Не все столбцы здесь имеют смысл для всех типов отношений.
type ObjectPostgresMigratePgClass struct {
	postgres_migrate_pg_class.PostgresMigratePgClass
	Version postgres_migrate_version.PostgresMigrateVersion	`json:"version"	gorm:"-:all"`
}
