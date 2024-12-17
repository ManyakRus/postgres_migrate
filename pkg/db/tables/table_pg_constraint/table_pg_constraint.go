package table_pg_constraint


// Table_PgConstraint - модель для таблицы pg_constraint: В каталоге pg_constraint хранятся ограничения-проверки, ограничения-исключения, а также ограничения первичного ключа, уникальности и внешних ключей, определённые для таблиц. (Ограничения столбцов описываются как и все остальные. Любое ограничение столбца равнозначно некоторому ограничению таблицы.) Ограничения на NULL представляются не здесь, а в каталоге pg_attribute.Для пользовательских триггеров ограничений (создаваемых командой CREATE CONSTRAINT TRIGGER) в этой таблице также создаётся запись.Здесь также хранятся ограничения доменов.
type Table_PgConstraint struct {
	Condeferrable bool	`json:"condeferrable" gorm:"column:condeferrable"`	//Является ли ограничение откладываемым?
	Condeferred bool	`json:"condeferred" gorm:"column:condeferred"`	//Является ли ограничение отложенным по умолчанию?
	Conexclop string	`json:"conexclop" gorm:"column:conexclop;default:null"`	//Для ограничения-исключения — список операторов исключения по столбцам
	Confdeltype string	`json:"confdeltype" gorm:"column:confdeltype;default:\"\""`	//Код действия при удалении внешнего ключа: a = нет действия, r = ограничить (restrict), c = каскадное действие (cascade), n = присвоить NULL, d = поведение по умолчанию
	Conffeqop string	`json:"conffeqop" gorm:"column:conffeqop;default:null"`	//Для внешнего ключа — список операторов равенства для сравнений FK = FK
	Confkey string	`json:"confkey" gorm:"column:confkey;default:null"`	//Для внешнего ключа определяет список столбцов, на которые он ссылается
	Confmatchtype string	`json:"confmatchtype" gorm:"column:confmatchtype;default:\"\""`	//Тип сопоставления внешнего ключа: f = полное (full), p = частичное (partial), s = простое (simple)
	Confrelid int64	`json:"confrelid" gorm:"column:confrelid;default:0"`	//Если это внешний ключ, таблица, на которую он ссылается; иначе 0
	Confupdtype string	`json:"confupdtype" gorm:"column:confupdtype;default:\"\""`	//Код действия при изменении внешнего ключа: a = нет действия, r = ограничить (restrict), c = каскадное действие (cascade), n = присвоить NULL, d = поведение по умолчанию
	Conindid int64	`json:"conindid" gorm:"column:conindid;default:0"`	//Индекс, поддерживающий это ограничение, если это ограничение уникальности, первичного или внешнего ключа, либо ограничение-исключение; в противном случае — 0
	Coninhcount int32	`json:"coninhcount" gorm:"column:coninhcount;default:0"`	//Число прямых предков этого ограничения. Ограничение с ненулевым числом предков нельзя удалить или переименовать.
	Conislocal bool	`json:"conislocal" gorm:"column:conislocal"`	//Ограничение определено локально в данном отношении. Заметьте, что ограничение может быть определено локально и при этом наследоваться.
	Conkey string	`json:"conkey" gorm:"column:conkey;default:null"`	//Для ограничений таблицы (включая внешние ключи, но не триггеры ограничений), определяет список столбцов, образующих ограничение
	Conname string	`json:"conname" gorm:"column:conname;default:\"\""`	//Имя ограничения (не обязательно уникальное!)
	Connamespace int64	`json:"connamespace" gorm:"column:connamespace;default:0"`	//OID пространства имён, содержащего это ограничение
	Connoinherit bool	`json:"connoinherit" gorm:"column:connoinherit"`	//Ограничение определено локально для данного отношения и является ненаследуемым.
	Conparentid int64	`json:"conparentid" gorm:"column:conparentid;default:0"`	//Соответствующее ограничение в родительской секционированной таблице, если это ограничение в секции; иначе ноль
	Conpfeqop string	`json:"conpfeqop" gorm:"column:conpfeqop;default:null"`	//Для внешнего ключа — список операторов равенства для сравнений PK = FK
	Conppeqop string	`json:"conppeqop" gorm:"column:conppeqop;default:null"`	//Для внешнего ключа — список операторов равенства для сравнений PK = PK
	Conrelid int64	`json:"conrelid" gorm:"column:conrelid;default:0"`	//Таблица, для которой установлено это ограничение; 0, если это не ограничение таблицы
	Contype string	`json:"contype" gorm:"column:contype;default:\"\""`	//c = ограничение-проверка (check), f = внешний ключ (foreign key), p = первичный ключ (primary key), u = ограничение уникальности (unique), t = триггер ограничения (trigger), x = ограничение-исключение (exclusion)
	Contypid int64	`json:"contypid" gorm:"column:contypid;default:0"`	//Домен, к которому относится это ограничение; 0, если это не ограничение домена
	Convalidated bool	`json:"convalidated" gorm:"column:convalidated"`	//Было ли ограничение проверено? В настоящее время значение false возможно только для внешних ключей и ограничений CHECK
	Oid int64	`json:"oid" gorm:"column:oid;primaryKey;autoIncrement:true"`	//Идентификатор строки (скрытый атрибут; должен выбираться явно)
	VersionID int64	`json:"version_id" gorm:"column:version_id;primaryKey;autoIncrement:true"`	//Версия изменений (ИД)

}
