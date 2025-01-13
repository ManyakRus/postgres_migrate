package table_postgres_migrate_pg_sequence

// Table_PostgresMigratePgSequence - модель для таблицы postgres_migrate_pg_sequence:
type Table_PostgresMigratePgSequence struct {
	IsDeleted    bool  `json:"is_deleted" gorm:"column:is_deleted"`                               //Признак что оригинальная запись удалена
	Seqcache     int64 `json:"seqcache" gorm:"column:seqcache;default:0"`                         //Размер кеша последовательности
	Seqcycle     bool  `json:"seqcycle" gorm:"column:seqcycle"`                                   //Зацикливается ли последовательность
	Seqincrement int64 `json:"seqincrement" gorm:"column:seqincrement;default:0"`                 //Шаг увеличения последовательности
	Seqmax       int64 `json:"seqmax" gorm:"column:seqmax;default:0"`                             //Максимальное значение последовательности
	Seqmin       int64 `json:"seqmin" gorm:"column:seqmin;default:0"`                             //Минимальное значение последовательности
	Seqrelid     int64 `json:"seqrelid" gorm:"column:seqrelid;primaryKey;autoIncrement:true"`     //OID записи в pg_class для этой последовательности
	Seqstart     int64 `json:"seqstart" gorm:"column:seqstart;default:0"`                         //Начальное значение последовательности
	Seqtypid     int64 `json:"seqtypid" gorm:"column:seqtypid;default:0"`                         //Тип данных последовательности
	VersionID    int64 `json:"version_id" gorm:"column:version_id;primaryKey;autoIncrement:true"` //Версия изменений (ИД)

}
