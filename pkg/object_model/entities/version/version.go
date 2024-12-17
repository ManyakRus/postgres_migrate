package version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables/table_version"
)

// Version - модель для таблицы version: 
type Version struct {
	table_version.Table_Version	
}
