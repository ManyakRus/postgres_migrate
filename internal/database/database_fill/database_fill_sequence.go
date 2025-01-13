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

// Fill_sequence - проверка изменения метаданных
func Fill_sequence(VersionID int64) error {
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
------------------------------- temp_pg_sequence_max --------------------------- 
drop table if exists temp_pg_sequence_max; 
CREATE TEMPORARY TABLE temp_pg_sequence_max (seqrelid oid, version_id bigint);
INSERT into temp_pg_sequence_max
SELECT
	ps.seqrelid,
	max(ps.version_id) as version_id
FROM
    SCHEMA_PM.postgres_migrate_pg_sequence as ps

JOIN
    SCHEMA_PM.postgres_migrate_pg_class as pmpc
ON
	pmpc.oid = ps.seqrelid


JOIN
	SCHEMA_PM.postgres_migrate_pg_namespace as pmpn
ON 
	pmpn.oid = pmpc.relnamespace


WHERE 1=1
	and pmpn.nspname = 'SCHEMA_DB'
	--and pmpc.is_deleted = false
	--and pmpn.is_deleted = false

GROUP BY
	ps.seqrelid
;


------------------------------- temp_pm_pg_sequence --------------------------- 
drop table if exists temp_pm_pg_sequence; 
CREATE TEMPORARY TABLE temp_pm_pg_sequence (
	seqrelid oid,
	seqtypid oid,
	seqstart int8,
	seqincrement int8,
	seqmax int8,
	seqmin int8,
	seqcache int8,
	seqcycle bool,
	is_deleted bool
);
INSERT into temp_pm_pg_sequence
SELECT
	ps.seqrelid,
	ps.seqtypid,
	ps.seqstart,
	ps.seqincrement,
	ps.seqmax,
	ps.seqmin,
	ps.seqcache,
	ps.seqcycle,
	ps.is_deleted
		
FROM
    SCHEMA_PM.postgres_migrate_pg_sequence as ps
	
JOIN
	temp_pg_sequence_max
ON 
	temp_pg_sequence_max.seqrelid = ps.seqrelid
	and temp_pg_sequence_max.version_id = ps.version_id

WHERE 1=1

;

------------------------------- temp_pg_sequence --------------------------- 
drop table if exists temp_pg_sequence; 
CREATE TEMPORARY TABLE temp_pg_sequence (
	seqrelid oid,
	seqtypid oid,
	seqstart int8,
	seqincrement int8,
	seqmax int8,
	seqmin int8,
	seqcache int8,
	seqcycle bool
);

INSERT into temp_pg_sequence as tc
SELECT
	ps.seqrelid,
	ps.seqtypid,
	ps.seqstart,
	ps.seqincrement,
	ps.seqmax,
	ps.seqmin,
	ps.seqcache,
	ps.seqcycle
		
FROM
    pg_catalog.pg_sequence as ps

JOIN
    pg_catalog.pg_class as pc
ON
	pc.oid = ps.seqrelid


JOIN
	pg_catalog.pg_namespace as pn
ON 
	pn.oid = pc.relnamespace


WHERE 1=1
	and pn.nspname = 'SCHEMA_DB'
;

------------------------------ сравнение -------------------------------------------
INSERT INTO SCHEMA_PM.postgres_migrate_pg_sequence
(
SELECT --новые строки
	:version_id as version_id,
	s.seqrelid,
	s.seqtypid,
	s.seqstart,
	s.seqincrement,
	s.seqmax,
	s.seqmin,
	s.seqcache,
	s.seqcycle,
	false as is_deleted
FROM
	temp_pm_pg_sequence as ps

RIGHT JOIN
	temp_pg_sequence as s
ON 
	s.seqrelid = ps.seqrelid

WHERE 1=1
	AND (ps.seqrelid IS NULL
		OR COALESCE(ps.is_deleted, false) = true
		)


UNION ALL


SELECT --изменённые строки
	:version_id as version_id,
	s.seqrelid,
	s.seqtypid,
	s.seqstart,
	s.seqincrement,
	s.seqmax,
	s.seqmin,
	s.seqcache,
	s.seqcycle,
	false as is_deleted
FROM
	temp_pm_pg_sequence as ps

JOIN
	temp_pg_sequence as s
ON 
	s.seqrelid = ps.seqrelid
	and ps.is_deleted = false

WHERE 0=1
	OR ps.seqrelid <> s.seqrelid
	OR ps.seqtypid <> s.seqtypid
	OR ps.seqstart <> s.seqstart
	OR ps.seqincrement <> s.seqincrement
	OR ps.seqmax <> s.seqmax
	OR ps.seqmin <> s.seqmin
	OR ps.seqcache <> s.seqcache
	OR ps.seqcycle <> s.seqcycle


UNION ALL


SELECT --удалённые строки
	:version_id as version_id,
	ps.seqrelid,
	ps.seqtypid,
	ps.seqstart,
	ps.seqincrement,
	ps.seqmax,
	ps.seqmin,
	ps.seqcache,
	ps.seqcycle,
	false as is_deleted

FROM
	temp_pm_pg_sequence as ps

LEFT JOIN
	temp_pg_sequence as s
ON 
	s.seqrelid = ps.seqrelid

WHERE 1=1
	AND s.seqrelid IS NULL
	AND ps.is_deleted = false
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
