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

// IsChanged_Namespace - проверка изменения метаданных
func IsChanged_Namespace() (int, error) {
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
------------------------------- temp_pg_namespace_max --------------------------- 
drop table if exists temp_pg_namespace_max; 
CREATE TEMPORARY TABLE temp_pg_namespace_max ("oid" oid, version_id bigint);
INSERT into temp_pg_namespace_max
SELECT
	pmpn."oid" ,
	max(pmpn.version_id) as version_id
FROM
	SCHEMA_PM.postgres_migrate_pg_namespace as pmpn

WHERE 1=1
	and pmpn.nspname = 'SCHEMA_DB'

GROUP BY
	pmpn."oid"
;


------------------------------- temp_pm_pg_namespace --------------------------- 
drop table if exists temp_pm_pg_namespace; 
CREATE TEMPORARY TABLE temp_pm_pg_namespace (
	"oid" oid,
	nspname name,
	nspowner oid,
	nspacl _aclitem
);
INSERT into temp_pm_pg_namespace
SELECT
	pmpn."oid",
	pmpn.nspname,
	pmpn.nspowner,
	pmpn.nspacl
		
FROM
    SCHEMA_PM.postgres_migrate_pg_namespace as pmpn
	
JOIN
	temp_pg_namespace_max
ON 
	temp_pg_namespace_max.oid = pmpn.oid
	and temp_pg_namespace_max.version_id = pmpn.version_id

WHERE 1=1

;

------------------------------- temp_pg_namespace --------------------------- 
drop table if exists temp_pg_namespace; 
CREATE TEMPORARY TABLE temp_pg_namespace (
	"oid" oid,
	nspname name,
	nspowner oid,
	nspacl _aclitem
);
INSERT into temp_pg_namespace
SELECT
	pn."oid",
	pn.nspname,
	pn.nspowner,
	pn.nspacl
		
FROM
    pg_catalog.pg_namespace as pn

WHERE 1=1
	and pn.nspname = 'SCHEMA_DB'
;

------------------------------ сравнение -------------------------------------------
SELECT
	temp_pg_namespace.nspname as name
FROM
	temp_pm_pg_namespace

FULL JOIN
	temp_pg_namespace
ON 
	temp_pg_namespace.oid = temp_pm_pg_namespace.oid

WHERE 
	(temp_pg_namespace.oid IS NULL
	OR
	temp_pm_pg_namespace.oid IS NULL
	)

UNION

SELECT
	pn.nspname
FROM
	temp_pm_pg_namespace as pn

FULL JOIN
	temp_pg_namespace as n
ON 
	n.oid = pn.oid

WHERE 0=1
	OR pn."oid" <> n."oid" 
	OR pn.nspname <> n.nspname
	--OR pn.nspowner <> n.nspowner
	--OR pn.nspacl <> n.nspacl

`

	TextSQL = strings.ReplaceAll(TextSQL, "SCHEMA_DB", config.Settings.DB_SCHEME_DATABASE)
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
