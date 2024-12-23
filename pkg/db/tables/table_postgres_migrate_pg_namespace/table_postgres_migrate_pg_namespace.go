package table_postgres_migrate_pg_namespace

// Table_PostgresMigratePgNamespace - модель для таблицы postgres_migrate_pg_namespace: В postgres_migrate_pg_namespace сохраняются пространства имён. Пространство имён представляет собой структуру, на которой основываются схемы SQL: в каждом пространстве имён без конфликтов может существовать отдельный набор отношений, типов и т. д.
type Table_PostgresMigratePgNamespace struct {
	IsDeleted bool   `json:"is_deleted" gorm:"column:is_deleted"`                               //Признак что оригинальная запись удалена
	Nspacl    string `json:"nspacl" gorm:"column:nspacl;default:null"`                          //Права доступа; за подробностями обратитесь к описанию GRANT и REVOKE
	Nspname   string `json:"nspname" gorm:"column:nspname;default:\"\""`                        //Имя пространства имён
	Nspowner  int64  `json:"nspowner" gorm:"column:nspowner;default:0"`                         //Владелец пространства имён
	Oid       int64  `json:"oid" gorm:"column:oid;primaryKey;autoIncrement:true"`               //Идентификатор строки (скрытый атрибут; должен выбираться явно)
	VersionID int64  `json:"version_id" gorm:"column:version_id;primaryKey;autoIncrement:true"` //Версия изменений (ИД)

}
