package object_postgres_migrate_pg_sequence

import (
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"reflect"
)

// versionPostgresMigratePgSequence - версия структуры модели, с учётом имен и типов полей
var versionObjectPostgresMigratePgSequence uint32

// Crud_PostgresMigratePgSequence - объект контроллер crud операций
var Crud_ObjectPostgresMigratePgSequence ICrud_ObjectPostgresMigratePgSequence

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPostgresMigratePgSequence interface {
	ICrud_ObjectPostgresMigratePgSequence_manual
	ReadObject(*ObjectPostgresMigratePgSequence) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPostgresMigratePgSequence) GetStructVersion() uint32 {
	if versionObjectPostgresMigratePgSequence == 0 {
		versionObjectPostgresMigratePgSequence = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPostgresMigratePgSequence
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPostgresMigratePgSequence) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPostgresMigratePgSequence) GetJSON() (string, error) {
	var Result string
	var err error

	bytes, err := json.Marshal(m)
	if err != nil {
		return Result, err
	}
	Result = string(bytes)
	return Result, err
}

// ---------------------------- CRUD операции ------------------------------------------------------------

// ReadObject - находит запись в БД по ID, и заполняет в объект, а также заполняет все поля у которых есть foreign key
func (m *ObjectPostgresMigratePgSequence) ReadObject() error {
	if Crud_ObjectPostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPostgresMigratePgSequence.ReadObject(m)

	return err
}
