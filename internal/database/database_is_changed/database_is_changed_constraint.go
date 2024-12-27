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

// IsChanged_Constraint - проверка изменения метаданных
func IsChanged_Constraint() (int, error) {
	Otvet := 0
	var err error

	//соединение
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Minute*10)
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	tx := db.WithContext(ctx)

	//
	TextSQL :=
		`
------------------------------- temp_pg_constraint_max --------------------------- 
drop table if exists temp_pg_constraint_max; 
CREATE TEMPORARY TABLE temp_pg_constraint_max ("oid" oid, version_id bigint);
INSERT into temp_pg_constraint_max
SELECT
	pmpc."oid" ,
	max(pmpc.version_id) as version_id
FROM
    SCHEMA_PM.postgres_migrate_pg_constraint as pmpc

JOIN
	SCHEMA_PM.postgres_migrate_pg_namespace as pmpn
ON 
	pmpn.oid = pmpc.connamespace

WHERE 1=1
	and pmpn.nspname = 'SCHEMA_DB'
	--and pmpc.is_deleted = false
	--and pmpn.is_deleted = false

GROUP BY
	pmpc."oid"
;


------------------------------- temp_pm_pg_constraint --------------------------- 
drop table if exists temp_pm_pg_constraint; 
CREATE TEMPORARY TABLE temp_pm_pg_constraint (
	"oid" oid,
	conname name,
	connamespace oid,
	contype char,
	condeferrable bool,
	condeferred bool,
	convalidated bool,
	conrelid oid,
	contypid oid,
	conindid oid,
	conparentid oid,
	confrelid oid,
	confupdtype char,
	confdeltype char,
	confmatchtype char,
	conislocal bool,
	coninhcount int4,
	connoinherit bool,
	conkey _int2,
	confkey _int2,
	conpfeqop _oid,
	conppeqop _oid,
	conffeqop _oid,
	conexclop _oid,
	is_deleted bool
);
INSERT into temp_pm_pg_constraint
SELECT
	pmpc."oid",
	pmpc.conname,
	pmpc.connamespace,
	pmpc.contype,
	pmpc.condeferrable,
	pmpc.condeferred,
	pmpc.convalidated,
	pmpc.conrelid,
	pmpc.contypid,
	pmpc.conindid,
	pmpc.conparentid,
	pmpc.confrelid,
	pmpc.confupdtype,
	pmpc.confdeltype,
	pmpc.confmatchtype,
	pmpc.conislocal,
	pmpc.coninhcount,
	pmpc.connoinherit,
	pmpc.conkey,
	pmpc.confkey,
	pmpc.conpfeqop,
	pmpc.conppeqop,
	pmpc.conffeqop,
	pmpc.conexclop,
	pmpc.is_deleted
		
FROM
    SCHEMA_PM.postgres_migrate_pg_constraint as pmpc
	
JOIN
	temp_pg_constraint_max
ON 
	temp_pg_constraint_max.oid = pmpc.oid
	and temp_pg_constraint_max.version_id = pmpc.version_id

WHERE 1=1
	--and pmpn.nspname = 'SCHEMA_DB'

;

------------------------------- temp_pg_constraint --------------------------- 
drop table if exists temp_pg_constraint; 
CREATE TEMPORARY TABLE temp_pg_constraint (
	"oid" oid,
	conname name,
	connamespace oid,
	contype char,
	condeferrable bool,
	condeferred bool,
	convalidated bool,
	conrelid oid,
	contypid oid,
	conindid oid,
	conparentid oid,
	confrelid oid,
	confupdtype char,
	confdeltype char,
	confmatchtype char,
	conislocal bool,
	coninhcount int4,
	connoinherit bool,
	conkey _int2,
	confkey _int2,
	conpfeqop _oid,
	conppeqop _oid,
	conffeqop _oid,
	conexclop _oid
);
INSERT into temp_pg_constraint as tc
SELECT
	pc."oid",
	pc.conname,
	pc.connamespace,
	pc.contype,
	pc.condeferrable,
	pc.condeferred,
	pc.convalidated,
	pc.conrelid,
	pc.contypid,
	pc.conindid,
	pc.conparentid,
	pc.confrelid,
	pc.confupdtype,
	pc.confdeltype,
	pc.confmatchtype,
	pc.conislocal,
	pc.coninhcount,
	pc.connoinherit,
	pc.conkey,
	pc.confkey,
	pc.conpfeqop,
	pc.conppeqop,
	pc.conffeqop,
	pc.conexclop
		
FROM
    pg_catalog.pg_constraint as pc


JOIN
	pg_catalog.pg_namespace as pn
ON 
	pn.oid = pc.connamespace


WHERE 1=1
	and pn.nspname = 'SCHEMA_DB'
;

------------------------------ сравнение -------------------------------------------
SELECT
	temp_pg_constraint.conname as name
FROM
	temp_pm_pg_constraint

FULL JOIN
	temp_pg_constraint
ON 
	temp_pg_constraint.oid = temp_pm_pg_constraint.oid

WHERE 
	(temp_pg_constraint.oid IS NULL
	OR
	temp_pm_pg_constraint.oid IS NULL
	)
	and COALESCE(temp_pm_pg_constraint.is_deleted, false) = false

UNION

SELECT
	c.conname
FROM
	temp_pm_pg_constraint as pc

FULL JOIN
	temp_pg_constraint as c
ON 
	c.oid = pc.oid

WHERE 0=1
	OR pc."oid" <> c."oid"
	OR pc.conname <> c.conname
	OR pc.connamespace <> c.connamespace
	OR pc.contype <> c.contype
	OR pc.condeferrable <> c.condeferrable
	OR pc.condeferred <> c.condeferred
	OR pc.convalidated <> c.convalidated
	OR pc.conrelid <> c.conrelid
	OR pc.contypid <> c.contypid
	OR pc.conindid <> c.conindid
	OR pc.conparentid <> c.conparentid
	OR pc.confrelid <> c.confrelid
	OR pc.confupdtype <> c.confupdtype
	OR pc.confdeltype <> c.confdeltype
	OR pc.confmatchtype <> c.confmatchtype
	OR pc.conislocal <> c.conislocal
	OR pc.coninhcount <> c.coninhcount
	OR pc.connoinherit <> c.connoinherit
	OR pc.conkey <> c.conkey
	OR pc.confkey <> c.confkey
	OR pc.conpfeqop <> c.conpfeqop
	OR pc.conppeqop <> c.conppeqop
	OR pc.conffeqop <> c.conffeqop
	OR pc.conexclop <> c.conexclop
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
