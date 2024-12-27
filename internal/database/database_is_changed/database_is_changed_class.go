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

// IsChanged_Class - проверка изменения метаданных
func IsChanged_Class() (int, error) {
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
------------------------------- temp_pg_class_max --------------------------- 
drop table if exists temp_pg_class_max; 
CREATE TEMPORARY TABLE temp_pg_class_max ("oid" oid, version_id bigint);
INSERT into temp_pg_class_max
SELECT
	pmpc."oid" ,
	max(pmpc.version_id) as version_id
FROM
    SCHEMA_PM.postgres_migrate_pg_class as pmpc

JOIN
	SCHEMA_PM.postgres_migrate_pg_namespace as pmpn
ON 
	pmpn.oid = pmpc.relnamespace

WHERE 1=1
	and pmpn.nspname = 'SCHEMA_DB'
	--and pmpc.is_deleted = false
	--and pmpn.is_deleted = false

GROUP BY
	pmpc."oid"
;


------------------------------- temp_pm_pg_class --------------------------- 
drop table if exists temp_pm_pg_class; 
CREATE TEMPORARY TABLE temp_pm_pg_class (
	"oid" oid,	
	relname name,
	relnamespace oid,
	reltype oid,
	reloftype oid,
	relowner oid,
	relam oid,
	relfilenode oid,
	reltablespace oid,
	relpages int4,
	reltuples float4,
	relallvisible int4,
	reltoastrelid oid,
	relhasindex bool,
	relisshared bool,
	relpersistence char,
	relkind char,
	relnatts int2,
	relchecks int2,
	relhasrules bool,
	relhastriggers bool,
	relhassubclass bool,
	relrowsecurity bool,
	relforcerowsecurity bool,
	relispopulated bool,
	relreplident char,
	relispartition bool,
	relrewrite oid,
	relfrozenxid xid,
	relminmxid xid,
	is_deleted bool
);
INSERT into temp_pm_pg_class
SELECT
	pmpc."oid",
	pmpc.relname,
	pmpc.relnamespace,
	pmpc.reltype,
	pmpc.reloftype,
	pmpc.relowner,
	pmpc.relam,
	pmpc.relfilenode,
	pmpc.reltablespace,
	pmpc.relpages,
	pmpc.reltuples,
	pmpc.relallvisible,
	pmpc.reltoastrelid,
	pmpc.relhasindex,
	pmpc.relisshared,
	pmpc.relpersistence,
	pmpc.relkind,
	pmpc.relnatts,
	pmpc.relchecks,
	pmpc.relhasrules,
	pmpc.relhastriggers,
	pmpc.relhassubclass,
	pmpc.relrowsecurity,
	pmpc.relforcerowsecurity,
	pmpc.relispopulated,
	pmpc.relreplident,
	pmpc.relispartition,
	pmpc.relrewrite,
	pmpc.relfrozenxid,
	pmpc.relminmxid,
	pmpc.is_deleted
		
FROM
    SCHEMA_PM.postgres_migrate_pg_class as pmpc
	
JOIN
	temp_pg_class_max
ON 
	temp_pg_class_max.oid = pmpc.oid
	and temp_pg_class_max.version_id = pmpc.version_id

WHERE 1=1
	--and pmpn.nspname = 'SCHEMA_DB'

;

------------------------------- temp_pg_class --------------------------- 
drop table if exists temp_pg_class; 
CREATE TEMPORARY TABLE temp_pg_class (
	"oid" oid,	
	relname name,
	relnamespace oid,
	reltype oid,
	reloftype oid,
	relowner oid,
	relam oid,
	relfilenode oid,
	reltablespace oid,
	relpages int4,
	reltuples float4,
	relallvisible int4,
	reltoastrelid oid,
	relhasindex bool,
	relisshared bool,
	relpersistence char,
	relkind char,
	relnatts int2,
	relchecks int2,
	relhasrules bool,
	relhastriggers bool,
	relhassubclass bool,
	relrowsecurity bool,
	relforcerowsecurity bool,
	relispopulated bool,
	relreplident char,
	relispartition bool,
	relrewrite oid,
	relfrozenxid xid,
	relminmxid xid
);
INSERT into temp_pg_class as tc
SELECT
	pc."oid",
	pc.relname,
	pc.relnamespace,
	pc.reltype,
	pc.reloftype,
	pc.relowner,
	pc.relam,
	pc.relfilenode,
	pc.reltablespace,
	pc.relpages,
	pc.reltuples,
	pc.relallvisible,
	pc.reltoastrelid,
	pc.relhasindex,
	pc.relisshared,
	pc.relpersistence,
	pc.relkind,
	pc.relnatts,
	pc.relchecks,
	pc.relhasrules,
	pc.relhastriggers,
	pc.relhassubclass,
	pc.relrowsecurity,
	pc.relforcerowsecurity,
	pc.relispopulated,
	pc.relreplident,
	pc.relispartition,
	pc.relrewrite,
	pc.relfrozenxid,
	pc.relminmxid
		
FROM
    pg_catalog.pg_class as pc


JOIN
	pg_catalog.pg_namespace as pn
ON 
	pn.oid = pc.relnamespace


WHERE 1=1
	and pn.nspname = 'SCHEMA_DB'
;

------------------------------ сравнение -------------------------------------------
SELECT
	temp_pg_class.relname as name
FROM
	temp_pm_pg_class

FULL JOIN
	temp_pg_class
ON 
	temp_pg_class.oid = temp_pm_pg_class.oid

WHERE 
	(temp_pg_class.oid IS NULL
	OR
	temp_pm_pg_class.oid IS NULL
	)
	and COALESCE(temp_pm_pg_class.is_deleted, false) = false

UNION

SELECT
	c.relname
FROM
	temp_pm_pg_class as pc

FULL JOIN
	temp_pg_class as c
ON 
	c.oid = pc.oid

WHERE 0=1
	OR pc."oid" <> c."oid"
	OR pc.relname <> c.relname
	OR pc.relnamespace <> c.relnamespace
	OR pc.reltype <> c.reltype
	OR pc.reloftype <> c.reloftype
	OR pc.relowner <> c.relowner
	OR pc.relam <> c.relam
	--OR pc.relfilenode <> c.relfilenode
	OR pc.reltablespace <> c.reltablespace
	--OR pc.relpages <> c.relpages
	--OR pc.reltuples <> c.reltuples
	--OR pc.relallvisible <> c.relallvisible
	OR pc.reltoastrelid <> c.reltoastrelid
	OR pc.relhasindex <> c.relhasindex
	OR pc.relisshared <> c.relisshared
	OR pc.relpersistence <> c.relpersistence
	OR pc.relkind <> c.relkind
	OR pc.relnatts <> c.relnatts
	OR pc.relchecks <> c.relchecks
	OR pc.relhasrules <> c.relhasrules
	OR pc.relhastriggers <> c.relhastriggers
	OR pc.relhassubclass <> c.relhassubclass
	OR pc.relrowsecurity <> c.relrowsecurity
	OR pc.relforcerowsecurity <> c.relforcerowsecurity
	OR pc.relispopulated <> c.relispopulated
	OR pc.relreplident <> c.relreplident
	OR pc.relispartition <> c.relispartition
	OR pc.relrewrite <> c.relrewrite
	--OR pc.relfrozenxid <> c.relfrozenxid
	--OR pc.relminmxid <> c.relminmxid
	OR pc.is_deleted = true

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
