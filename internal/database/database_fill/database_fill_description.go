package database_fill

import (
	"context"
	"fmt"
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"strings"
	"time"
)

// Fill_description - проверка изменения метаданных
func Fill_description(VersionID int64) error {
	//Otvet := 0
	var err error

	//соединение
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Minute*10)
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	tx := db.WithContext(ctx)

	//
	TextSQL := `
------------------------------- temp_pg_description_max --------------------------- 
drop table if exists temp_pg_description_max; 
CREATE TEMPORARY TABLE temp_pg_description_max ("objoid" oid, classoid oid, objsubid int4, version_id bigint);
INSERT into temp_pg_description_max
SELECT
	pmpd.objoid,
	pmpd.classoid,
	pmpd.objsubid,
	max(pmpd.version_id) as version_id
FROM
    SCHEMA_PM.postgres_migrate_pg_description as pmpd

JOIN
	SCHEMA_PM.postgres_migrate_pg_class AS pg_class_objoid   
ON 
	pmpd.objoid   = pg_class_objoid.oid


JOIN
	pg_catalog.pg_class AS pg_class_classoid 
ON 
	pmpd.classoid = pg_class_classoid.oid


JOIN
	SCHEMA_PM.postgres_migrate_pg_namespace as pmpn
ON 
	pmpn.oid = pg_class_objoid.relnamespace

WHERE 1=1
	and pg_class_classoid.relname = 'pg_class'
	and pmpn.nspname = 'SCHEMA_DB'
	--and pmpd.is_deleted = false
	--and pmpn.is_deleted = false

GROUP BY
	pmpd.objoid,
	pmpd.classoid,
	pmpd.objsubid
;


------------------------------- temp_pm_pg_description --------------------------- 
drop table if exists temp_pm_pg_description; 
CREATE TEMPORARY TABLE temp_pm_pg_description (
	objoid oid,
	classoid oid,
	objsubid int4,
	description text,
	is_deleted bool
);
INSERT into temp_pm_pg_description
SELECT
	pmpd.objoid,
	pmpd.classoid,
	pmpd.objsubid,
	pmpd.description,
	pmpd.is_deleted
		
FROM
    SCHEMA_PM.postgres_migrate_pg_description as pmpd
	
JOIN
	temp_pg_description_max
ON 
	temp_pg_description_max.objoid = pmpd.objoid
	and temp_pg_description_max.classoid = pmpd.classoid
	and temp_pg_description_max.objsubid = pmpd.objsubid
	and temp_pg_description_max.version_id = pmpd.version_id

WHERE 1=1
	--and pmpn.nspname = 'SCHEMA_DB'

;

------------------------------- temp_pg_description --------------------------- 
drop table if exists temp_pg_description; 
CREATE TEMPORARY TABLE temp_pg_description (
	objoid oid,
	classoid oid,
	objsubid int4,
	description text
);
INSERT into temp_pg_description as tc
SELECT
	pd.objoid,
	pd.classoid,
	pd.objsubid,
	pd.description
		
FROM
    pg_catalog.pg_description as pd


JOIN
	pg_catalog.pg_class AS pg_class_objoid   
ON 
	pd.objoid   = pg_class_objoid.oid


JOIN
	pg_catalog.pg_class AS pg_class_classoid 
ON 
	pd.classoid = pg_class_classoid.oid


JOIN
	pg_catalog.pg_namespace as pn
ON 
	pn.oid = pg_class_objoid.relnamespace

WHERE 1=1
	and pg_class_classoid.relname = 'pg_class'
	and pn.nspname = 'SCHEMA_DB'
;


------------------------------ сравнение -------------------------------------------
INSERT INTO SCHEMA_PM.postgres_migrate_pg_description
(
SELECT --новые строки
	:version_id as version_id,
	d.objoid,
	d.classoid,
	d.objsubid,
	d.description,
	false as is_deleted
FROM
	temp_pm_pg_description as pd

RIGHT JOIN
	temp_pg_description as d
ON 
	d.objoid = pd.objoid
	and d.classoid = pd.classoid
	and d.objsubid = pd.objsubid

WHERE 1=1
	AND (pd.objoid IS NULL
		OR COALESCE(pd.is_deleted, false) = true
		)


UNION ALL


SELECT --изменённые строки
	:version_id as version_id,
	d.objoid,
	d.classoid,
	d.objsubid,
	d.description,
	false as is_deleted
FROM
	temp_pm_pg_description as pd

JOIN
	temp_pg_description as d
ON 
	d.objoid = pd.objoid
	and d.classoid = pd.classoid
	and d.objsubid = pd.objsubid
	and pd.is_deleted = false

WHERE 0=1
	OR pd.objoid <> d.objoid
	OR pd.classoid <> d.classoid
	OR pd.objsubid <> d.objsubid
	OR pd.description <> d.description


UNION ALL


SELECT --удалённые строки
	:version_id as version_id,
	pd.objoid,
	pd.classoid,
	pd.objsubid,
	pd.description,
	false as is_deleted

FROM
	temp_pm_pg_description as pd

LEFT JOIN
	temp_pg_description as d
ON 
	d.objoid = pd.objoid
	and d.classoid = pd.classoid
	and d.objsubid = pd.objsubid

WHERE 1=1
	AND d.objoid IS NULL
	AND pd.is_deleted = false
)

`

	TextSQL = strings.ReplaceAll(TextSQL, "SCHEMA_DB", config.Settings.DB_SCHEME_DATABASE)
	TextSQL = strings.ReplaceAll(TextSQL, "SCHEMA_PM", postgres_gorm.Settings.DB_SCHEMA)
	TextSQL = strings.ReplaceAll(TextSQL, ":version_id", micro.StringFromInt64(VersionID))

	//tx = postgres_gorm.RawMultipleSQL(tx, TextSQL)
	//tx = db.Raw(TextSQL)
	tx = db.Exec(TextSQL)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf("db.Raw() error: %w", err)
		log.Error(err)
		return err
	}

	////
	//MassNames := make([]string, 0)
	//tx = tx.Scan(&MassNames)
	//err = tx.Error
	//if err != nil {
	//	err = fmt.Errorf("tx.Scan() error: %w", err)
	//	log.Error(err)
	//	return err
	//}

	//Otvet = len(MassNames)

	return err
}
