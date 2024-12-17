package object_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
)

// ObjectPgClass - объект для таблицы pg_class: В каталоге pg_class описываются таблицы и практически всё, что имеет столбцы или каким-то образом подобно таблице. Сюда входят индексы (но смотрите также pg_index), последовательности, представления, материализованные представления, составные типы и таблицы TOAST; см. relkind. Далее, подразумевая все эти типы объектов, мы будем говорить об «отношениях». Не все столбцы здесь имеют смысл для всех типов отношений.
type ObjectPgClass struct {
	pg_class.PgClass
	Version version.Version	`json:"version"	gorm:"-:all"`
}
