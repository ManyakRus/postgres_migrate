package sql_attributes

import (
	"context"
	"fmt"
	"github.com/ManyakRus/postgres_migrate/internal/config"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_gorm"
	"strings"
	"time"
)

// SequenceAlter - хранит старое и новое название таблиц
type SequenceAlter struct {
	Name                   string
	Name_Old               string
	Sequence_type_id_Old   int64
	Sequence_start_Old     int64
	Sequence_increment_Old int64
	Sequence_max_Old       int64
	Sequence_min_Old       int64
	Sequence_cache_Old     int64
	Sequence_cycle_Old     bool
	TypeName               string
	TypeName_Old           string
	postgres_migrate_pg_sequence.PostgresMigratePgSequence
}

// Start_Secuences_alter - добавляет текст SQL в Text
func Start_Secuences_alter(Settings *config.SettingsINI, VersionID int64) (string, error) {
	var err error
	Otvet := ""

	//найдём массив новых
	MassClass, err := Find_MassSequence_Alter(Settings, VersionID)
	if err != nil {
		err = fmt.Errorf("Find_MassClass_Create() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//создадим файлы
	Otvet1, err := TextSQL_Alter(Settings, MassClass)
	if err != nil {
		err = fmt.Errorf("Create_files() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//
	Otvet = Otvet + Otvet1

	return Otvet, err
}

// Find_MassSequence_Alter - возвращает массив postgres_migrate_pg_class
func Find_MassSequence_Alter(Settings *config.SettingsINI, VersionID int64) ([]SequenceAlter, error) {
	Otvet := make([]SequenceAlter, 0)
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
	and pmpc.relpersistence = 'p' --постоянная таблица (permanent)
	and pmpc.relkind = 'r' --обычная таблица

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
	relminmxid xid
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
	pmpc.relminmxid
		
FROM
    SCHEMA_PM.postgres_migrate_pg_class as pmpc
	
JOIN
	temp_pg_class_max
ON 
	temp_pg_class_max.oid = pmpc.oid
	and temp_pg_class_max.version_id = pmpc.version_id

WHERE 1=1

;
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
	is_deleted bool,
	Name_Old text,
	TypeName_Old text
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
	ps.is_deleted,
	pc.relname,
	pt.typname
		
FROM
    SCHEMA_PM.postgres_migrate_pg_sequence as ps
	
JOIN
	temp_pg_sequence_max
ON 
	temp_pg_sequence_max.seqrelid = ps.seqrelid
	and temp_pg_sequence_max.version_id = ps.version_id


JOIN
	temp_pm_pg_class as pc
ON
	pc.oid = ps.seqrelid


JOIN
	pg_catalog.pg_type as pt
ON 
	pt.oid = ps.seqtypid

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
	seqcycle bool,
	Name Text,
	TypeName Text
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
	ps.seqcycle,
	pc.relname,
	pt.typname as TypeName
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


JOIN
	pg_catalog.pg_class as c
ON
	c.oid = ps.seqtypid


JOIN
	pg_catalog.pg_type as pt
ON 
	pt.oid = ps.seqtypid


WHERE 1=1
	and pn.nspname = 'SCHEMA_DB'
;

------------------------------ сравнение -------------------------------------------
SELECT
	s.seqrelid,
	s.seqtypid,
	s.seqstart,
	s.seqincrement,
	s.seqmax,
	s.seqmin,
	s.seqcache,
	s.seqcycle,
	s.relname as Name_Old,
	s.Name as Name,
	s.seqtypid as Sequence_type_id_Old,
	s.seqstart as Sequence_start_Old,
	s.seqincrement as Sequence_increment_Old,
	s.seqmax as Sequence_max_Old,
	s.seqmin as Sequence_min_Old,
	s.seqcache as Sequence_cache_Old,
	s.seqcycle as Sequence_cycle_Old,
	ps.TypeName_Old,
	s.TypeName
FROM
	temp_pm_pg_sequence as ps

JOIN
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
	OR ps.Name_Old <> s.Name
	OR ps.TypeName_Old <> s.TypeName
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

// TextSQL_Alter - возвращает текст SQL для создания таблиц
func TextSQL_Alter(Settings *config.SettingsINI, MassNames []SequenceAlter) (string, error) {
	Otvet := ""
	var err error

	for _, v := range MassNames {
		//имя
		if v.Name != v.Name_Old {
			Otvet1 := `ALTER SEQUENCE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.Name_Old + `"` + " RENAME TO " + `"` + v.Name + `"` + ";\n"
			Otvet = Otvet + Otvet1
		}

		//тип
		if v.Seqtypid != v.Sequence_type_id_Old {
			Otvet1 := `ALTER SEQUENCE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.Name + `"` + " AS " + v.TypeName + ";\n"
			Otvet = Otvet + Otvet1
		}

		//MINVALUE
		if v.Seqmin != v.Sequence_min_Old {
			sValue := micro.StringFromInt64(v.Seqmin)
			Otvet1 := `ALTER SEQUENCE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.Name + `"` + " MINVALUE " + sValue + ";\n"
			Otvet = Otvet + Otvet1
		}

		//MAXVALUE
		if v.Seqmax != v.Sequence_max_Old {
			sValue := micro.StringFromInt64(v.Seqmax)
			Otvet1 := `ALTER SEQUENCE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.Name + `"` + " MAXVALUE " + sValue + ";\n"
			Otvet = Otvet + Otvet1
		}

		//START
		if v.Seqstart != v.Sequence_start_Old {
			sValue := micro.StringFromInt64(v.Seqstart)
			Otvet1 := `ALTER SEQUENCE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.Name + `"` + " START " + sValue + ";\n"
			Otvet = Otvet + Otvet1
		}

		//CACHE
		if v.Seqcache != v.Sequence_cache_Old {
			sValue := micro.StringFromInt64(v.Seqcache)
			Otvet1 := `ALTER SEQUENCE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.Name + `"` + " CACHE " + sValue + ";\n"
			Otvet = Otvet + Otvet1
		}

		//INCREMENT
		if v.Seqincrement != v.Sequence_increment_Old {
			sValue := micro.StringFromInt64(v.Seqincrement)
			Otvet1 := `ALTER SEQUENCE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.Name + `"` + " INCREMENT " + sValue + ";\n"
			Otvet = Otvet + Otvet1
		}

		//CYCLE
		if v.Seqcache != v.Sequence_cache_Old {
			TextNO := Find_TextNo(v.Seqcycle)
			Otvet1 := `ALTER SEQUENCE "` + Settings.DB_SCHEME_DATABASE + `"."` + v.Name + `"` + TextNO + " CYCLE " + ";\n"
			Otvet = Otvet + Otvet1
		}

	}

	return Otvet, err
}

// Find_TextNo - возвращает " NO" для true
func Find_TextNo(Seqcycle bool) string {
	Otvet := ""
	if Seqcycle == true {
		return " NO"
	}

	return Otvet
}
