package sql_attributes

import (
	"context"
	"fmt"
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"strings"
	"time"
)

// AttributeSQL - структура для атрибутов и дополнительных данных
type AttributeSQL struct {
	TableName     string `json:"TableName" gorm:"column:TableName"`
	TypeName      string `json:"TypeName" gorm:"column:TypeName"`
	CollationName string `json:"CollationName" gorm:"column:CollationName"`
	postgres_migrate_pg_attribute.PostgresMigratePgAttribute
}

// Start_Attributes_create - добавляет текст SQL в Text
func Start_Attributes_create(Settings *config.SettingsINI, VersionID int64) (string, error) {
	var err error
	Otvet := ""

	//найдём массив новых
	MassClass, err := Find_MassAttributeSQL_Create(Settings, VersionID)
	if err != nil {
		err = fmt.Errorf("Find_MassAttributeSQL_Create() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//создадим файлы
	Otvet1, err := TextSQL_Create(Settings, MassClass)
	if err != nil {
		err = fmt.Errorf("Create_files() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//
	Otvet = Otvet + Otvet1

	return Otvet, err
}

// Find_MassAttributeSQL_Create - возвращает массив postgres_migrate_pg_class
func Find_MassAttributeSQL_Create(Settings *config.SettingsINI, VersionID int64) ([]AttributeSQL, error) {
	Otvet := make([]AttributeSQL, 0)
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
	pa.attrelid,
	pa.attname,
	pa.atttypid,
	pa.attstattarget,
	pa.attlen,
	pa.attnum,
	pa.attndims,
	pa.attcacheoff,
	pa.atttypmod,
	pa.attbyval,
	pa.attstorage,
	pa.attalign,
	pa.attnotnull,
	pa.atthasdef,
	pa.atthasmissing,
	pa.attidentity,
	pa.attgenerated,
	pa.attisdropped,
	pa.attislocal,
	pa.attinhcount,
	pa.attcollation,
	pa.is_deleted,
	pa.attmissingval
		
FROM
    SCHEMA_PM.postgres_migrate_pg_attribute as pa
	
JOIN
	temp_pg_attribute_max
ON 
	temp_pg_attribute_max.attrelid = pa.attrelid
	and temp_pg_attribute_max.attname = pa.attname
	and temp_pg_attribute_max.version_id = pa.version_id

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
	a.attmissingval::Text
		
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
	a.attmissingval::Text as attmissingval,

	c.relname as TableName,
	t.typname as TypeName
FROM
	temp_pm_pg_attribute as pa

FULL JOIN
	temp_pg_attribute as a
ON 
	a.attrelid = pa.attrelid
	and a.attname = pa.attname
	

LEFT JOIN
	pg_catalog.pg_class as c
ON 
	pc.oid = pa.attrelid


LEFT JOIN
	pg_catalog.pg_type as t
ON
	t.oid = pa.atttypid


WHERE 
	(a.attrelid IS NULL
	OR
	pa.attrelid IS NULL
	)
	and COALESCE(pa.is_deleted, false) = false


UNION


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
	a.attmissingval::Text as attmissingval,

	c.relname as TableName,
	t.typname as TypeName,
	pc.callname as CollationName
FROM
	temp_pm_pg_attribute as pa

FULL JOIN
	temp_pg_attribute as a
ON 
	a.attrelid = pa.attrelid
	and a.attname = pa.attname


LEFT JOIN
	pg_catalog.pg_class as c
ON 
	pc.oid = pa.attrelid


LEFT JOIN
	pg_catalog.pg_type as t
ON
	t.oid = pa.atttypid


LEFT JOIN
	pg_catalog.pg_collation as pc
ON
	pc.oid = pa.attcollation


WHERE 0=1
	OR pa.attrelid <> a.attrelid
	OR pa.attname <> a.attname
	OR pa.atttypid <> a.atttypid
	--OR pa.attstattarget <> a.attstattarget
	OR pa.attlen <> a.attlen
	OR pa.attnum <> a.attnum
	OR pa.attndims <> a.attndims
	--OR pa.attcacheoff <> a.attcacheoff
	OR pa.atttypmod <> a.atttypmod
	OR pa.attbyval <> a.attbyval
	OR pa.attstorage <> a.attstorage
	OR pa.attalign <> a.attalign
	OR pa.attnotnull <> a.attnotnull
	OR pa.atthasdef <> a.atthasdef
	OR pa.atthasmissing <> a.atthasmissing
	OR pa.attidentity <> a.attidentity
	OR pa.attgenerated <> a.attgenerated
	OR pa.attisdropped <> a.attisdropped
	OR pa.attislocal <> a.attislocal
	--OR pa.attinhcount <> a.attinhcount
	OR pa.attcollation <> a.attcollation
	OR pa.attmissingval <> a.attmissingval
	OR pa.is_deleted = true



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

// TextSQL_Create - возвращает текст SQL для создания таблиц
func TextSQL_Create(Settings *config.SettingsINI, MassAttribute []AttributeSQL) (string, error) {
	Otvet := ""
	var err error

	for _, v := range MassAttribute {
		//длина текста
		TextLength := ""
		if v.Atttypmod > 0 {
			TextLength = "(" + micro.StringFromInt32(v.Atttypmod) + ")"
		}

		//
		TextCollation := ""
		if v.CollationName != "" {
			TextCollation = " COLLATE " + v.CollationName
		}

		//
		Otvet1 := `ALTER TABLE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.TableName + `"` + " ADD COLUMN " + v.Attname + " " + v.TypeName + TextLength + TextCollation + ";\n"

		Otvet = Otvet + Otvet1
	}

	return Otvet, err
}
