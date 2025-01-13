package sql_attributes

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

// Start_Attributes_delete - добавляет текст SQL в Text
func Start_Attributes_delete(Settings *config.SettingsINI, VersionID int64) (string, error) {
	var err error
	Otvet := ""

	//найдём массив новых
	MassAttribute, err := Find_MassAttribute_Delete(Settings, VersionID)
	if err != nil {
		err = fmt.Errorf("Find_MassClass_Create() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//создадим файлы
	Otvet1, err := TextSQL_Delete(Settings, MassAttribute)
	if err != nil {
		err = fmt.Errorf("Create_files() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//
	Otvet = Otvet + Otvet1

	return Otvet, err
}

// Find_MassAttribute_Delete - возвращает массив postgres_migrate_pg_class
func Find_MassAttribute_Delete(Settings *config.SettingsINI, VersionID int64) ([]AttributeAlter, error) {
	Otvet := make([]AttributeAlter, 0)
	var err error

	//соединение
	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Minute*10)
	defer ctxCancelFunc()

	db := postgres_gorm.GetConnection()
	tx := db.WithContext(ctx)

	//
	TextSQL := `
------------------------------- temp_pg_attribute_max --------------------------- 
drop table if exists temp_pg_attribute_max; 
CREATE TEMPORARY TABLE temp_pg_attribute_max ("attrelid" oid, attname name, version_id bigint);
INSERT into temp_pg_attribute_max
SELECT
	pmpa."attrelid",
	pmpa."attname",
	max(pmpa.version_id) as version_id
FROM
    SCHEMA_PM.postgres_migrate_pg_attribute as pmpa


JOIN
	SCHEMA_PM.postgres_migrate_pg_class as pmpc
ON 
	pmpc.oid = pmpa.attrelid


JOIN
	SCHEMA_PM.postgres_migrate_pg_namespace as pmpn
ON 
	pmpn.oid = pmpc.relnamespace


WHERE 1=1
	and pmpn.nspname = 'SCHEMA_DB'
	--and pmpa.is_deleted = false
	--and pmpc.is_deleted = false
	--and pmpn.is_deleted = false

GROUP BY
	pmpa."attrelid",
	pmpa."attname"
;


------------------------------- temp_pm_pg_attribute --------------------------- 
drop table if exists temp_pm_pg_attribute; 
CREATE TEMPORARY TABLE temp_pm_pg_attribute (
	attrelid oid,
	attname name,
	atttypid oid,
	attstattarget int4,
	attlen int2,
	attnum int2,
	attndims int4,
	attcacheoff int4,
	atttypmod int4,
	attbyval bool,
	attstorage char,
	attalign char,
	attnotnull bool,
	atthasdef bool,
	atthasmissing bool,
	attidentity char,
	attgenerated char,
	attisdropped bool,
	attislocal bool,
	attinhcount int4,
	attcollation oid,
	is_deleted bool,
	attmissingval Text
);
INSERT into temp_pm_pg_attribute
SELECT
	pmpa.attrelid,
	pmpa.attname,
	pmpa.atttypid,
	pmpa.attstattarget,
	pmpa.attlen,
	pmpa.attnum,
	pmpa.attndims,
	pmpa.attcacheoff,
	pmpa.atttypmod,
	pmpa.attbyval,
	pmpa.attstorage,
	pmpa.attalign,
	pmpa.attnotnull,
	pmpa.atthasdef,
	pmpa.atthasmissing,
	pmpa.attidentity,
	pmpa.attgenerated,
	pmpa.attisdropped,
	pmpa.attislocal,
	pmpa.attinhcount,
	pmpa.attcollation,
	pmpa.is_deleted,
	pmpa.attmissingval
		
FROM
    SCHEMA_PM.postgres_migrate_pg_attribute as pmpa
	
JOIN
	temp_pg_attribute_max
ON 
	temp_pg_attribute_max.attrelid = pmpa.attrelid
	and temp_pg_attribute_max.attname = pmpa.attname
	and temp_pg_attribute_max.version_id = pmpa.version_id

WHERE 1=1

;

------------------------------- temp_pg_attribute --------------------------- 
drop table if exists temp_pg_attribute; 
CREATE TEMPORARY TABLE temp_pg_attribute (
	attrelid oid,
	attname name,
	atttypid oid,
	attstattarget int4,
	attlen int2,
	attnum int2,
	attndims int4,
	attcacheoff int4,
	atttypmod int4,
	attbyval bool,
	attstorage char,
	attalign char,
	attnotnull bool,
	atthasdef bool,
	atthasmissing bool,
	attidentity char,
	attgenerated char,
	attisdropped bool,
	attislocal bool,
	attinhcount int4,
	attcollation oid,
	attmissingval Text
);
INSERT into temp_pg_attribute as tc
SELECT
	a.attrelid,
	a.attname,
	a.atttypid,
	a.attstattarget,
	a.attlen,
	a.attnum,
	a.attndims,
	a.attcacheoff,
	a.atttypmod,
	a.attbyval,
	a.attstorage,
	a.attalign,
	a.attnotnull,
	a.atthasdef,
	a.atthasmissing,
	a.attidentity,
	a.attgenerated,
	a.attisdropped,
	a.attislocal,
	a.attinhcount,
	a.attcollation,
	a.attmissingval::Text as attmissingval
		
FROM
    pg_catalog.pg_attribute as a


JOIN
	pg_catalog.pg_class as pc
ON 
	pc.oid = a.attrelid


JOIN
	pg_catalog.pg_namespace as pn
ON 
	pn.oid = pc.relnamespace


WHERE 1=1
	and pn.nspname = 'SCHEMA_DB'
;

------------------------------ сравнение -------------------------------------------
SELECT
	a.attname as attname,
	c.relname as TableName
FROM
	temp_pm_pg_attribute as pa

LEFT JOIN
	temp_pg_attribute as a
ON 
	a.attrelid = pa.attrelid
	and a.attname = pa.attname
	

LEFT JOIN
	pg_catalog.pg_class as c
ON 
	c.oid = pa.attrelid


WHERE 1=1
	AND a.attname IS NULL


`

	TextSQL = strings.ReplaceAll(TextSQL, "SCHEMA_DB", Settings.DB_SCHEME_DATABASE)
	TextSQL = strings.ReplaceAll(TextSQL, "SCHEMA_PM", postgres_gorm.Settings.DB_SCHEMA)
	TextSQL = strings.ReplaceAll(TextSQL, ":version_id", micro.StringFromInt64(VersionID))

	tx = postgres_gorm.RawMultipleSQL(tx, TextSQL)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf("db.Raw() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//
	//MassClass := make([]postgres_migrate_pg_class.PostgresMigratePgClass, 0)
	tx = tx.Scan(&Otvet)
	err = tx.Error
	if err != nil {
		err = fmt.Errorf("tx.Scan() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//Otvet = len(MassNames)

	return Otvet, err
}

// TextSQL_Delete - возвращает текст SQL для создания таблиц
func TextSQL_Delete(Settings *config.SettingsINI, MassClass []AttributeAlter) (string, error) {
	Otvet := ""
	var err error

	for _, v := range MassClass {
		Otvet1 := `ALTER TABLE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.TableName + `" DROP COLUMN "` + v.Attname + `";` + "\n"

		Otvet = Otvet + Otvet1
	}

	return Otvet, err
}
