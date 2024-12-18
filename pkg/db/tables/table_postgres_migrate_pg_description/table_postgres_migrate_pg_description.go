package table_postgres_migrate_pg_description

// Table_PostgresMigratePgDescription - модель для таблицы postgres_migrate_pg_description: В каталоге postgres_migrate_pg_description хранятся дополнительные описания (комментарии) для объектов баз данных. Описания можно задавать с помощью команды COMMENT и просматривать в psql, используя команды \d. В начальном содержимом postgres_migrate_pg_description представлены описания многих встроенных системных объектов.Также смотрите каталог pg_shdescription, который играет подобную роль в отношении совместно используемых объектов в кластере баз данных.
type Table_PostgresMigratePgDescription struct {
	Classoid    int64  `json:"classoid" gorm:"column:classoid;primaryKey;autoIncrement:true"`     //OID системного каталога, к которому относится этот объект
	Description string `json:"description" gorm:"column:description;default:\"\""`                //Произвольный текст, служащий описанием данного объекта
	Objoid      int64  `json:"objoid" gorm:"column:objoid;primaryKey;autoIncrement:true"`         //OID объекта, к которому относится это описание
	Objsubid    int32  `json:"objsubid" gorm:"column:objsubid;primaryKey;autoIncrement:true"`     //Для комментария к столбцу таблицы это номер столбца (objoid и classoid указывают на саму таблицу). Для всех других типов объектов это поле содержит ноль.
	VersionID   int64  `json:"version_id" gorm:"column:version_id;primaryKey;autoIncrement:true"` //Версия изменений (ИД)

}
