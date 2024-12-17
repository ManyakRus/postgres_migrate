package object_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionPgDescription - версия структуры модели, с учётом имен и типов полей
var versionObjectPgDescription uint32

// Crud_PgDescription - объект контроллер crud операций
var Crud_ObjectPgDescription ICrud_ObjectPgDescription

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPgDescription interface {
	ICrud_ObjectPgDescription_manual
	ReadObject(*ObjectPgDescription) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPgDescription) GetStructVersion() uint32 {
	if versionObjectPgDescription == 0 {
		versionObjectPgDescription = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPgDescription
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPgDescription) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPgDescription) GetJSON() (string, error) {
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
func (m *ObjectPgDescription) ReadObject() error {
	if Crud_ObjectPgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPgDescription.ReadObject(m)

	return err
}
