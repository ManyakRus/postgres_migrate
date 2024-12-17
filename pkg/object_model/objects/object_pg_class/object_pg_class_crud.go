package object_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionPgClass - версия структуры модели, с учётом имен и типов полей
var versionObjectPgClass uint32

// Crud_PgClass - объект контроллер crud операций
var Crud_ObjectPgClass ICrud_ObjectPgClass

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPgClass interface {
	ICrud_ObjectPgClass_manual
	ReadObject(*ObjectPgClass) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPgClass) GetStructVersion() uint32 {
	if versionObjectPgClass == 0 {
		versionObjectPgClass = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPgClass
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPgClass) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPgClass) GetJSON() (string, error) {
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
func (m *ObjectPgClass) ReadObject() error {
	if Crud_ObjectPgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPgClass.ReadObject(m)

	return err
}
