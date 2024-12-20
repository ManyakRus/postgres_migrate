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

// IsChanged_Index - проверка изменения метаданных
func IsChanged_Index() (int, error) {
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
------------------------------- temp_pg_index_max --------------------------- 
drop table if exists temp_pg_index_max; 
CREATE TEMPORARY TABLE temp_pg_index_max ("indexrelid" oid, version_id bigint);
INSERT into temp_pg_index_max
SELECT
	pmpi."indexrelid" ,
	max(pmpi.version_id) as version_id
FROM
    SCHEMA_PM.postgres_migrate_pg_index as pmpi


JOIN
	SCHEMA_PM.postgres_migrate_pg_class as pmpc
ON 
	pmpc.oid = pmpi.indexrelid


JOIN
	SCHEMA_PM.postgres_migrate_pg_namespace as pmpn
ON 
	pmpn.oid = pmpc.relnamespace

WHERE 1=1
	and pmpn.nspname = 'SCHEMA_DB'

GROUP BY
	pmpi."indexrelid"
;


------------------------------- temp_pm_pg_index --------------------------- 
drop table if exists temp_pm_pg_index; 
CREATE TEMPORARY TABLE temp_pm_pg_index (
	indexrelid oid,
	indrelid oid,
	indnatts int2,
	indnkeyatts int2,
	indisunique bool,
	indisprimary bool,
	indisexclusion bool,
	indimmediate bool,
	indisclustered bool,
	indisvalid bool,
	indcheckxmin bool,
	indisready bool,
	indislive bool,
	indisreplident bool,
	indkey int2vector,
	indcollation oidvector,
	indclass oidvector,
	indoption int2vector,
	indexprs pg_node_tree,
	indpred pg_node_tree
);
INSERT into temp_pm_pg_index
SELECT
	pmpi.indexrelid,
	pmpi.indrelid,
	pmpi.indnatts,
	pmpi.indnkeyatts,
	pmpi.indisunique,
	pmpi.indisprimary,
	pmpi.indisexclusion,
	pmpi.indimmediate,
	pmpi.indisclustered,
	pmpi.indisvalid,
	pmpi.indcheckxmin,
	pmpi.indisready,
	pmpi.indislive,
	pmpi.indisreplident,
	pmpi.indkey,
	pmpi.indcollation,
	pmpi.indclass,
	pmpi.indoption,
	pmpi.indexprs,
	pmpi.indpred
		
FROM
    SCHEMA_PM.postgres_migrate_pg_index as pmpi
	
JOIN
	temp_pg_index_max
ON 
	temp_pg_index_max.indexrelid = pmpi.indexrelid
	and temp_pg_index_max.version_id = pmpi.version_id

WHERE 1=1
;

------------------------------- temp_pg_index --------------------------- 
drop table if exists temp_pg_index; 
CREATE TEMPORARY TABLE temp_pg_index (
	indexrelid oid,
	indrelid oid,
	indnatts int2,
	indnkeyatts int2,
	indisunique bool,
	indisprimary bool,
	indisexclusion bool,
	indimmediate bool,
	indisclustered bool,
	indisvalid bool,
	indcheckxmin bool,
	indisready bool,
	indislive bool,
	indisreplident bool,
	indkey int2vector,
	indcollation oidvector,
	indclass oidvector,
	indoption int2vector,
	indexprs pg_node_tree,
	indpred pg_node_tree
);
INSERT into temp_pg_index as tc
SELECT
	pi.indexrelid,
	pi.indrelid,
	pi.indnatts,
	pi.indnkeyatts,
	pi.indisunique,
	pi.indisprimary,
	pi.indisexclusion,
	pi.indimmediate,
	pi.indisclustered,
	pi.indisvalid,
	pi.indcheckxmin,
	pi.indisready,
	pi.indislive,
	pi.indisreplident,
	pi.indkey,
	pi.indcollation,
	pi.indclass,
	pi.indoption,
	pi.indexprs,
	pi.indpred
		
FROM
    pg_catalog.pg_index as pi



JOIN
	pg_catalog.pg_class as pc
ON 
	pc.oid = pi.indexrelid



JOIN
	pg_catalog.pg_namespace as pn
ON 
	pn.oid = pc.relnamespace


WHERE 1=1
	and pn.nspname = 'SCHEMA_DB'
;

------------------------------ сравнение -------------------------------------------
SELECT
	temp_pg_index.indexrelid as id
FROM
	temp_pm_pg_index

FULL JOIN
	temp_pg_index
ON 
	temp_pg_index.indexrelid = temp_pm_pg_index.indexrelid

WHERE 
	(temp_pg_index.indexrelid IS NULL
	OR
	temp_pm_pg_index.indexrelid IS NULL
	)

UNION

SELECT
	i.indexrelid
FROM
	temp_pm_pg_index as pi

FULL JOIN
	temp_pg_index as i
ON 
	i.indexrelid = pi.indexrelid

WHERE 0=1
	OR pi.indexrelid <> i.indexrelid
	--OR pi.indrelid <> i.indrelid
	--OR pi.indnatts <> i.indnatts
	--OR pi.indnkeyatts <> i.indnkeyatts
	OR pi.indisunique <> i.indisunique
	--OR pi.indisprimary <> i.indisprimary 
	--OR pi.indisexclusion <> i.indisexclusion
	--OR pi.indimmediate <> i.indimmediate
	--OR pi.indisclustered <> i.indisclustered
	--OR pi.indisvalid <> i.indisvalid
	--OR pi.indcheckxmin <> i.indcheckxmin
	--OR pi.indisready <> i.indisready
	--OR pi.indislive <> i.indislive
	--OR pi.indisreplident <> i.indisreplident
	--OR pi.indkey <> i.indkey
	--OR pi.indcollation <> i.indcollation
	--OR pi.indclass <> i.indclass
	--OR pi.indoption <> i.indoption
	--OR pi.indexprs <> i.indexprs
	--OR pi.indpred <> i.indpred

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
	MassNames := make([]int, 0)
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