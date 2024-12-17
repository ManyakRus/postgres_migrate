package object_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
)

// ObjectVersion - объект для таблицы version: 
type ObjectVersion struct {
	version.Version
}
