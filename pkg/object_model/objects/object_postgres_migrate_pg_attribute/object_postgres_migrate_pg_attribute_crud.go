package object_postgres_migrate_pg_attribute

import (
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"reflect"
)

// versionPostgresMigratePgAttribute - версия структуры модели, с учётом имен и типов полей
var versionObjectPostgresMigratePgAttribute uint32

// Crud_PostgresMigratePgAttribute - объект контроллер crud операций
var Crud_ObjectPostgresMigratePgAttribute ICrud_ObjectPostgresMigratePgAttribute

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPostgresMigratePgAttribute interface {
	ICrud_ObjectPostgresMigratePgAttribute_manual
	ReadObject(*ObjectPostgresMigratePgAttribute) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPostgresMigratePgAttribute) GetStructVersion() uint32 {
	if versionObjectPostgresMigratePgAttribute == 0 {
		versionObjectPostgresMigratePgAttribute = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPostgresMigratePgAttribute
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPostgresMigratePgAttribute) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPostgresMigratePgAttribute) GetJSON() (string, error) {
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
func (m *ObjectPostgresMigratePgAttribute) ReadObject() error {
	if Crud_ObjectPostgresMigratePgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPostgresMigratePgAttribute.ReadObject(m)

	return err
}
