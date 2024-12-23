package table_postgres_migrate_pg_attribute

// Table_PostgresMigratePgAttribute - модель для таблицы postgres_migrate_pg_attribute: В каталоге postgres_migrate_pg_attribute хранится информация о столбцах таблицы. Для каждого столбца каждой таблицы в postgres_migrate_pg_attribute существует ровно одна строка.
type Table_PostgresMigratePgAttribute struct {
	Attalign      string `json:"attalign" gorm:"column:attalign;default:\"\""`                      //Копия pg_type.typalign из записи типа столбца
	Attbyval      bool   `json:"attbyval" gorm:"column:attbyval"`                                   //Копия pg_type.typbyval из записи типа столбца
	Attcacheoff   int32  `json:"attcacheoff" gorm:"column:attcacheoff;default:0"`                   //Всегда -1 в постоянном хранилище, но когда запись загружается в память, в этом поле может кешироваться смещение атрибута в строке
	Attcollation  int64  `json:"attcollation" gorm:"column:attcollation;default:0"`                 //Заданное для столбца правило сортировки, либо ноль, если тип столбца не сортируемый.
	Attgenerated  string `json:"attgenerated" gorm:"column:attgenerated;default:\"\""`              //Если нулевой байт (''), это не генерируемый столбец. В противном случае, s означает, что генерируемое значение хранится (stored). (Другие значения могут быть добавлены в будущем.)
	Atthasdef     bool   `json:"atthasdef" gorm:"column:atthasdef"`                                 //Столбец имеет значение по умолчанию, в этом случае в каталоге pg_attrdef будет соответствующая запись, определяющая это значение.
	Atthasmissing bool   `json:"atthasmissing" gorm:"column:atthasmissing"`                         //Столбец имеет значение, которое используется, когда он полностью отсутствует в строке. Это имеет место, когда столбец добавляется с неизменчивым значением DEFAULT после создания строки. Фактическое значение хранится в attmissingval.
	Attidentity   string `json:"attidentity" gorm:"column:attidentity;default:\"\""`                //Пустой символ ('') указывает, что это не столбец идентификации. Символ a указывает, что значение генерируется всегда, а d — что значение генерируется по умолчанию.
	Attinhcount   int32  `json:"attinhcount" gorm:"column:attinhcount;default:0"`                   //Число прямых предков этого столбца. Столбец с ненулевым числом предков нельзя удалить или переименовать.
	Attisdropped  bool   `json:"attisdropped" gorm:"column:attisdropped"`                           //Столбец был удалён и теперь не является рабочим. Удалённый столбец может по-прежнему физически присутствовать в таблице, но анализатор запросов его игнорирует, так что обратиться к нему из SQL нельзя.
	Attislocal    bool   `json:"attislocal" gorm:"column:attislocal"`                               //Столбец определён локально в данном отношении. Заметьте, что столбец может быть определён локально и при этом наследоваться.
	Attlen        int32  `json:"attlen" gorm:"column:attlen;default:0"`                             //Копия pg_type.typlen из записи типа столбца
	Attname       string `json:"attname" gorm:"column:attname;primaryKey;autoIncrement:true"`       //Имя столбца
	Attndims      int32  `json:"attndims" gorm:"column:attndims;default:0"`                         //Число размерностей, если столбец имеет тип массива; ноль в противном случае. (В настоящее время число размерностей массива не контролируется, поэтому любое ненулевое значение по сути означает «это массив».)
	Attnotnull    bool   `json:"attnotnull" gorm:"column:attnotnull"`                               //Представляет ограничение NOT NULL.
	Attnum        int32  `json:"attnum" gorm:"column:attnum;default:0"`                             //Порядковый номер столбца. Обычные столбцы нумеруются по возрастанию, начиная с 1. Системные столбцы, такие как oid, имеют (обычно) отрицательные номера.
	Attrelid      int64  `json:"attrelid" gorm:"column:attrelid;primaryKey;autoIncrement:true"`     //Таблица, к которой принадлежит столбец
	Attstattarget int32  `json:"attstattarget" gorm:"column:attstattarget;default:0"`               //Столбец attstattarget управляет детализацией статистики, собираемой по этому столбцу командой ANALYZE. Нулевое значение указывает, что статистика не собирается. При отрицательном значении используется системное ограничение статистики по умолчанию. Точное значение положительных величин определяется типом данных. Для скалярных типов данных, attstattarget задаёт и целевое число собираемых «самых частых значений», и целевое число создаваемых групп гистограммы.
	Attstorage    string `json:"attstorage" gorm:"column:attstorage;default:\"\""`                  //Обычно копия pg_type.typstorage из записи типа столбца. Для типов, поддерживающих TOAST, можно изменять это значение после создания столбца и таким образом управлять политикой хранения.
	Atttypid      int64  `json:"atttypid" gorm:"column:atttypid;default:0"`                         //Тип данных этого столбца
	Atttypmod     int32  `json:"atttypmod" gorm:"column:atttypmod;default:0"`                       //В поле atttypmod записывается дополнительное число, связанное с определённым типом данных, указываемое при создании таблицы (например, максимальный размер столбца varchar). Это значение передаётся функциям ввода и преобразования длины конкретного типа. Для типов, которым не нужен atttypmod, это обычно -1.
	IsDeleted     bool   `json:"is_deleted" gorm:"column:is_deleted"`                               //Признак что оригинальная запись удалена
	VersionID     int64  `json:"version_id" gorm:"column:version_id;primaryKey;autoIncrement:true"` //Версия изменений (ИД)

}
