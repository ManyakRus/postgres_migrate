package database_is_changed

import (
	"context"
	"fmt"
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/postgres_gorm"
	"strings"
	"time"
)

// IsChanged_Description - проверка изменения метаданных
func IsChanged_Description() (int, error) {
	Otvet := 0
	var err error

	//соединение
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Minute*10)
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	tx := db.WithContext(ctx)

	//
	TextSQL := `
	-- таблицы
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
	SCHEMA_PM.postgres_migrate_pg_namespace as pmpn
ON 
	pmpn.oid = pmpd.classoid

WHERE 1=1
	and pmpn.nspname = 'SCHEMA_BD'

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
	description text
);
INSERT into temp_pm_pg_description
SELECT
	pmpd.objoid,
	pmpd.classoid,
	pmpd.objsubid,
	pmpd.description
		
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
	--and pmpn.nspname = 'SCHEMA_BD'

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
	pg_catalog.pg_namespace as pn
ON 
	pn.oid = pd.classoid


WHERE 1=1
	and pn.nspname = 'SCHEMA_BD'
;

------------------------------ сравнение -------------------------------------------
SELECT
	d.description as name
FROM
	temp_pm_pg_description as pd

FULL JOIN
	temp_pg_description as d
ON 
	d.objoid = pd.objoid
	and d.classoid = pd.classoid
	and d.objsubid = pd.objsubid

WHERE 
	(d.objoid IS NULL
	OR
	pd.objoid IS NULL
	)

UNION

SELECT
	d.description
FROM
	temp_pm_pg_description as pd

FULL JOIN
	temp_pg_description as d
ON 
	d.objoid = pd.objoid
	and d.classoid = pd.classoid
	and d.objsubid = pd.objsubid

WHERE 0=1
	OR pd.objoid <> d.objoid
	OR pd.classoid <> d.classoid
	OR pd.objsubid <> d.objsubid
	OR pd.description <> d.description

`

	TextSQL = strings.ReplaceAll(TextSQL, "SCHEMA_BD", config.Settings.DB_SCHEME_DATABASE)
	TextSQL = strings.ReplaceAll(TextSQL, "SCHEMA_PM", postgres_gorm.Settings.DB_SCHEMA)

	tx = postgres_gorm.RawMultipleSQL(tx, TextSQL)
	//tx = db.Raw(TextSQL)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf("db.Raw() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//
	MassNames := make([]string, 0)
	tx = tx.Scan(&MassNames)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf("tx.Scan() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	Otvet = len(MassNames)

	return Otvet, err
}
