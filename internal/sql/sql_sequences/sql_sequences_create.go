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

// SequencesSQL - структура для атрибутов и дополнительных данных
type SequencesSQL struct {
	TableName     string `json:"TableName" gorm:"column:TableName"`
	TypeName      string `json:"TypeName" gorm:"column:TypeName"`
	CollationName string `json:"CollationName" gorm:"column:CollationName"`
	postgres_migrate_pg_attribute.PostgresMigratePgAttribute
}

// Start_Sequences_create - добавляет текст SQL в Text
func Start_Sequences_create(Settings *config.SettingsINI, VersionID int64) (string, error) {
	var err error
	Otvet := ""

	//найдём массив новых
	MassClass, err := Find_MassSequenceSQL_Create(Settings, VersionID)
	if err != nil {
		err = fmt.Errorf("Find_MassSequenceSQL_Create() error: %w", err)
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

// Find_MassSequenceSQL_Create - возвращает массив postgres_migrate_pg_class
func Find_MassSequenceSQL_Create(Settings *config.SettingsINI, VersionID int64) ([]SequencesSQL, error) {
	Otvet := make([]SequencesSQL, 0)
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
	seqcycle bool,
	Name text
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
	pc.relname	
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

WHERE 1=1
	AND (ps.seqrelid is null  OR  ps.is_deleted = true)


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
func TextSQL_Create(Settings *config.SettingsINI, MassAttribute []SequencesSQL) (string, error) {
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
