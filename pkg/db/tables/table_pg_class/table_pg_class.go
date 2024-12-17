package table_pg_class


// Table_PgClass - модель для таблицы pg_class: В каталоге pg_class описываются таблицы и практически всё, что имеет столбцы или каким-то образом подобно таблице. Сюда входят индексы (но смотрите также pg_index), последовательности, представления, материализованные представления, составные типы и таблицы TOAST; см. relkind. Далее, подразумевая все эти типы объектов, мы будем говорить об «отношениях». Не все столбцы здесь имеют смысл для всех типов отношений.
type Table_PgClass struct {
	Oid int64	`json:"oid" gorm:"column:oid;primaryKey;autoIncrement:true"`	//Идентификатор строки (скрытый атрибут; должен выбираться явно)
	Relallvisible int32	`json:"relallvisible" gorm:"column:relallvisible;default:0"`	//Число страниц, помеченных как «полностью видимые» в карте видимости таблицы. Это лишь примерная оценка, используемая планировщиком. Она обновляется командами VACUUM, ANALYZE и несколькими командами DDL, например, CREATE INDEX.
	Relam int64	`json:"relam" gorm:"column:relam;default:0"`	//Если это индекс, применяемый метод доступа (B-дерево, хеш и т. д.)
	Relchecks int32	`json:"relchecks" gorm:"column:relchecks;default:0"`	//Число ограничений CHECK в таблице; см. каталог pg_constraint
	Relfilenode int64	`json:"relfilenode" gorm:"column:relfilenode;default:0"`	//Имя файла на диске с этим отношением; ноль означает, что это «отображённое» представление, имя файла для которого определяется состоянием на нижнем уровне
	Relforcerowsecurity bool	`json:"relforcerowsecurity" gorm:"column:relforcerowsecurity"`	//True, если защита на уровне строк (когда она включена) также применяется к владельцу таблицы; см. каталог pg_policy
	Relfrozenxid int32	`json:"relfrozenxid" gorm:"column:relfrozenxid;default:0"`	//Идентификаторы транзакций, предшествующие данному, в этой таблице заменены постоянным («замороженным») идентификатором транзакции. Это нужно для определения, когда требуется очищать таблицу для предотвращения зацикливания идентификаторов или для сокращения объёма pg_clog. Если это отношение — не таблица, значение равно нулю (InvalidTransactionId).
	Relhasindex bool	`json:"relhasindex" gorm:"column:relhasindex"`	//True, если это таблица и она имеет (или недавно имела) индексы
	Relhasrules bool	`json:"relhasrules" gorm:"column:relhasrules"`	//True, если для таблицы определены (или были определены) правила; см. каталог pg_rewrite
	Relhassubclass bool	`json:"relhassubclass" gorm:"column:relhassubclass"`	//True, если у таблицы есть (или были) потомки в иерархии наследования
	Relhastriggers bool	`json:"relhastriggers" gorm:"column:relhastriggers"`	//True, если для таблицы определены (или были определены) триггеры; см. каталог pg_trigger
	Relispartition bool	`json:"relispartition" gorm:"column:relispartition"`	//
	Relispopulated bool	`json:"relispopulated" gorm:"column:relispopulated"`	//True, если отношение наполнено данными (это истинно для всех отношений, кроме некоторых материализованных представлений)
	Relisshared bool	`json:"relisshared" gorm:"column:relisshared"`	//True, если эта таблица разделяется всеми базами данных в кластере. Разделяемыми являются только некоторые системные каталоги (как например, pg_database).
	Relkind string	`json:"relkind" gorm:"column:relkind;default:\"\""`	//r = обычная таблица, i = индекс (index), S = последовательность (sequence), v = представление (view), m = материализованное представление (materialized view), c = составной тип (composite), t = таблица TOAST, f = сторонняя таблица (foreign)
	Relminmxid int32	`json:"relminmxid" gorm:"column:relminmxid;default:0"`	//Идентификаторы мультитранзакций, предшествующие данному, в этой таблице заменены другим идентификатором транзакции. Это нужно для определения, когда требуется очищать таблицу для предотвращения зацикливания идентификаторов мультитранзакций или для сокращения объёма pg_multixact. Если это отношение — не таблица, значение равно нулю (InvalidMultiXactId).
	Relname string	`json:"relname" gorm:"column:relname;default:\"\""`	//Имя таблицы, индекса, представления и т. п.
	Relnamespace int64	`json:"relnamespace" gorm:"column:relnamespace;default:0"`	//OID пространства имён, содержащего это отношение
	Relnatts int32	`json:"relnatts" gorm:"column:relnatts;default:0"`	//Число пользовательских столбцов в отношении (системные столбцы не считаются). Столько же соответствующих строк должно быть в pg_attribute. См. также pg_attribute.attnum.
	Reloftype int64	`json:"reloftype" gorm:"column:reloftype;default:0"`	//Для типизированных таблиц, OID нижележащего составного типа, или ноль для всех других отношений
	Relowner int64	`json:"relowner" gorm:"column:relowner;default:0"`	//Владелец отношения
	Relpages int32	`json:"relpages" gorm:"column:relpages;default:0"`	//Размер представления этой таблицы на диске (в страницах размера BLCKSZ). Это лишь примерная оценка, используемая планировщиком. Она обновляется командами VACUUM, ANALYZE и несколькими командами DDL, например, CREATE INDEX.
	Relpersistence string	`json:"relpersistence" gorm:"column:relpersistence;default:\"\""`	//p = постоянная таблица (permanent), u = нежурналируемая таблица (unlogged), t = временная таблица (temporary)
	Relreplident string	`json:"relreplident" gorm:"column:relreplident;default:\"\""`	//Столбцы, формирующие «идентификатор реплики» для строк: d = по умолчанию (первичный ключ, если есть), n = никакие (nothing), f = все столбцы, i = индекс со свойством indisreplident (если ранее использованный индекс удалён, действует так же, как n)
	Relrewrite int64	`json:"relrewrite" gorm:"column:relrewrite;default:0"`	//
	Relrowsecurity bool	`json:"relrowsecurity" gorm:"column:relrowsecurity"`	//True, если для таблицы включена защита на уровне строк; см. каталог pg_policy
	Reltablespace int64	`json:"reltablespace" gorm:"column:reltablespace;default:0"`	//Табличное пространство, в котором хранится это отношение. Если ноль, подразумевается пространство базы данных по умолчанию. (Не имеет значения, если с отношением не связан файл на диске.)
	Reltoastrelid int64	`json:"reltoastrelid" gorm:"column:reltoastrelid;default:0"`	//OID таблицы TOAST, связанной с данной таблицей, или 0, если таковой нет. В таблицу TOAST, как во вторичную, «выносятся» большие атрибуты.
	Reltuples float32	`json:"reltuples" gorm:"column:reltuples;default:0"`	//Число строк в таблице. Это лишь примерная оценка, используемая планировщиком. Она обновляется командами VACUUM, ANALYZE и несколькими командами DDL, например, CREATE INDEX.
	Reltype int64	`json:"reltype" gorm:"column:reltype;default:0"`	//OID типа данных, соответствующего типу строки этой таблицы, если таковой есть (ноль для индексов, так как они не имеют записи в pg_type)
	VersionID int64	`json:"version_id" gorm:"column:version_id;primaryKey;autoIncrement:true"`	//Версия изменений (ИД)

}
