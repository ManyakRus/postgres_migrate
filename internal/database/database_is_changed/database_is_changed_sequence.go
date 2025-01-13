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

// IsChanged_Sequence - проверка изменения метаданных
func IsChanged_Sequence() (int, error) {
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
SELECT
	COALESCE(s.seqrelid, ps.seqrelid) as id
FROM
	temp_pm_pg_sequence as ps

FULL JOIN
	temp_pg_sequence as s
ON 
	s.seqrelid = ps.seqrelid

WHERE 
	(s.seqrelid IS NULL
	OR
	ps.seqrelid IS NULL
	)
	and COALESCE(ps.is_deleted, false) = false


UNION


SELECT
	COALESCE(s.seqrelid, ps.seqrelid) as id
FROM
	temp_pm_pg_sequence as ps

FULL JOIN
	temp_pg_sequence as s
ON 
	s.seqrelid = ps.seqrelid

WHERE 0=1
	OR ps.seqrelid <> s.seqrelid
	OR ps.seqtypid <> s.seqtypid
	OR ps.seqstart <> s.seqstart
	OR ps.seqincrement <> s.seqincrement
	OR ps.seqmax <> s.seqmax
	OR ps.seqmin <> s.seqmin
	OR ps.seqcache <> s.seqcache
	OR ps.seqcycle <> s.seqcycle
	OR ps.is_deleted = true

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
	MassID := make([]string, 0)
	tx = tx.Scan(&MassID)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf("tx.Scan() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	Otvet = len(MassID)

	return Otvet, err
}
