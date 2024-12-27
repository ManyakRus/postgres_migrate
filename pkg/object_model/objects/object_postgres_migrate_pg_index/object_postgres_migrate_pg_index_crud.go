package object_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionPostgresMigratePgIndex - версия структуры модели, с учётом имен и типов полей
var versionObjectPostgresMigratePgIndex uint32

// Crud_PostgresMigratePgIndex - объект контроллер crud операций
var Crud_ObjectPostgresMigratePgIndex ICrud_ObjectPostgresMigratePgIndex

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPostgresMigratePgIndex interface {
	ICrud_ObjectPostgresMigratePgIndex_manual
	ReadObject(*ObjectPostgresMigratePgIndex) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPostgresMigratePgIndex) GetStructVersion() uint32 {
	if versionObjectPostgresMigratePgIndex == 0 {
		versionObjectPostgresMigratePgIndex = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPostgresMigratePgIndex
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPostgresMigratePgIndex) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPostgresMigratePgIndex) GetJSON() (string, error) {
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
func (m *ObjectPostgresMigratePgIndex) ReadObject() error {
	if Crud_ObjectPostgresMigratePgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPostgresMigratePgIndex.ReadObject(m)

	return err
}
