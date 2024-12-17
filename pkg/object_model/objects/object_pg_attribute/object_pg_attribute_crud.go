package object_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionPgAttribute - версия структуры модели, с учётом имен и типов полей
var versionObjectPgAttribute uint32

// Crud_PgAttribute - объект контроллер crud операций
var Crud_ObjectPgAttribute ICrud_ObjectPgAttribute

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPgAttribute interface {
	ICrud_ObjectPgAttribute_manual
	ReadObject(*ObjectPgAttribute) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPgAttribute) GetStructVersion() uint32 {
	if versionObjectPgAttribute == 0 {
		versionObjectPgAttribute = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPgAttribute
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPgAttribute) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPgAttribute) GetJSON() (string, error) {
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
func (m *ObjectPgAttribute) ReadObject() error {
	if Crud_ObjectPgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPgAttribute.ReadObject(m)

	return err
}
