package table_postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/tables"
	"time"
)

// Table_PostgresMigrateVersion - модель для таблицы postgres_migrate_version:
type Table_PostgresMigrateVersion struct {
	tables.NameStruct
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create;"`     //Дата создания элемента
	ID         int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`                 //Уникальный технический идентификатор
	ModifiedAt time.Time `json:"modified_at" gorm:"column:modified_at;default:null;autoUpdateTime"` //Дата последнего изменения элемента

}
