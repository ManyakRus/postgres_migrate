package table_pg_namespace


// Table_PgNamespace - модель для таблицы pg_namespace: В pg_namespace сохраняются пространства имён. Пространство имён представляет собой структуру, на которой основываются схемы SQL: в каждом пространстве имён без конфликтов может существовать отдельный набор отношений, типов и т. д.
type Table_PgNamespace struct {
	Nspacl string	`json:"nspacl" gorm:"column:nspacl;default:null"`	//Права доступа; за подробностями обратитесь к описанию GRANT и REVOKE
	Nspname string	`json:"nspname" gorm:"column:nspname;default:\"\""`	//Имя пространства имён
	Nspowner int64	`json:"nspowner" gorm:"column:nspowner;default:0"`	//Владелец пространства имён
	Oid int64	`json:"oid" gorm:"column:oid;primaryKey;autoIncrement:true"`	//Идентификатор строки (скрытый атрибут; должен выбираться явно)
	VersionID int64	`json:"version_id" gorm:"column:version_id;primaryKey;autoIncrement:true"`	//Версия изменений (ИД)

}
