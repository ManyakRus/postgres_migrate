package table_postgres_migrate_pg_index


// Table_PostgresMigratePgIndex - модель для таблицы postgres_migrate_pg_index: В каталоге postgres_migrate_pg_index содержится часть информации об индексах. Остальная информация в основном находится в postgres_migrate_pg_class.
type Table_PostgresMigratePgIndex struct {
	Indcheckxmin bool	`json:"indcheckxmin" gorm:"column:indcheckxmin"`	//Если true, запросы не должны использовать этот индекс, пока поле xmin данной записи в postgres_migrate_pg_index не окажется ниже их горизонта событий TransactionXmin, так как таблица может содержать оборванные цепочки HOT, с видимыми несовместимыми строками
	Indclass string	`json:"indclass" gorm:"column:indclass;default:\"\""`	//Для каждого столбца в ключе индекса этот массив содержит OID применяемых классов операторов. Подробнее это рассматривается в описании pg_opclass.
	Indcollation string	`json:"indcollation" gorm:"column:indcollation;default:\"\""`	//Для каждого столбца в ключе индекса этот массив содержит OID правила сортировки для применения в этом индексе.
	Indexprs string	`json:"indexprs" gorm:"column:indexprs;default:null"`	//Деревья выражений (в представлении nodeToString()) для атрибутов индекса, не являющихся простыми ссылками на столбцы. Этот список содержит один элемент для каждого нулевого значения в indkey. Значением может быть NULL, если все атрибуты индекса представляют собой простые ссылки.
	Indexrelid int64	`json:"indexrelid" gorm:"column:indexrelid;primaryKey;autoIncrement:true"`	//OID записи в postgres_migrate_pg_class для этого индекса
	Indimmediate bool	`json:"indimmediate" gorm:"column:indimmediate"`	//Если true, проверка уникальности осуществляется непосредственно при добавлении данных (неприменимо, если значение indisunique не true)
	Indisclustered bool	`json:"indisclustered" gorm:"column:indisclustered"`	//Если true, таблица в последний раз кластеризовалась по этому индексу
	Indisexclusion bool	`json:"indisexclusion" gorm:"column:indisexclusion"`	//Если true, этот индекс поддерживает ограничение-исключение
	Indislive bool	`json:"indislive" gorm:"column:indislive"`	//Если false, индекс находится в процессе удаления и его следует игнорировать для любых целей (включая вопрос применимости HOT)
	Indisprimary bool	`json:"indisprimary" gorm:"column:indisprimary"`	//Если true, этот индекс представляет первичный ключ таблицы (в этом случае и в поле indisunique должно быть значение true)
	Indisready bool	`json:"indisready" gorm:"column:indisready"`	//Если true, индекс готов к добавлению данных. Значение false означает, что индекс игнорируется операциями INSERT/UPDATE.
	Indisreplident bool	`json:"indisreplident" gorm:"column:indisreplident"`	//Если true, этот индекс выбран в качестве «идентификатора реплики» командой ALTER TABLE ... REPLICA IDENTITY USING INDEX ...
	Indisunique bool	`json:"indisunique" gorm:"column:indisunique"`	//Если true, это уникальный индекс
	Indisvalid bool	`json:"indisvalid" gorm:"column:indisvalid"`	//Если true, индекс можно применять в запросах. Значение false означает, что индекс, возможно, неполный: он будет тем не менее изменяться командами INSERT/UPDATE, но безопасно применять его в запросах нельзя. Если он уникальный, свойство уникальности также не гарантируется.
	Indkey string	`json:"indkey" gorm:"column:indkey;default:\"\""`	//Это массив из indnatts значений, указывающих, какие столбцы таблицы индексирует этот индекс. Например, значения 1 3 будут означать, что ключ индекса составляют первый и третий столбцы таблицы. Ноль в этом массиве означает, что соответствующий атрибут индекса определяется выражением со столбцами таблицы, а не просто ссылкой на столбец.
	Indnatts int32	`json:"indnatts" gorm:"column:indnatts;default:0"`	//Число столбцов в индексе (повторяет значение postgres_migrate_pg_class.relnatts)
	Indnkeyatts int32	`json:"indnkeyatts" gorm:"column:indnkeyatts;default:0"`	//Количество ключевых столбцов в индексе. «Ключевые столбцы» это обычные столбцы, в отличие от «включённых» столбцов.
	Indoption string	`json:"indoption" gorm:"column:indoption;default:\"\""`	//Это массив из indnatts значений, в которых хранятся битовые флаги для отдельных столбцов. Значение этих флагов определяется методом доступа конкретного индекса.
	Indpred string	`json:"indpred" gorm:"column:indpred;default:null"`	//Дерево выражения (в представлении nodeToString()) для предиката частичного индекса, либо NULL, если это не частичный индекс.
	Indrelid int64	`json:"indrelid" gorm:"column:indrelid;default:0"`	//OID записи в postgres_migrate_pg_class для таблицы, к которой относится этот индекс
	IsDeleted bool	`json:"is_deleted" gorm:"column:is_deleted"`	//Признак что оригинальная запись удалена
	VersionID int64	`json:"version_id" gorm:"column:version_id;primaryKey;autoIncrement:true"`	//Версия изменений (ИД)

}
