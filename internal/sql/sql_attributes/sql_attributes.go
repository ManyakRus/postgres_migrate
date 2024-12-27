package sql_attributes

import "github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"

// AttributeSQL - структура для атрибутов и дополнительных данных
type AttributeSQL struct {
	TableName string
	TypeName  string
	postgres_migrate_pg_attribute.PostgresMigratePgAttribute
}
