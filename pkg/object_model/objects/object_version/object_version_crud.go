package object_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionVersion - версия структуры модели, с учётом имен и типов полей
var versionObjectVersion uint32

// Crud_Version - объект контроллер crud операций
var Crud_ObjectVersion ICrud_ObjectVersion

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectVersion interface {
	ICrud_ObjectVersion_manual
	ReadObject(*ObjectVersion) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectVersion) GetStructVersion() uint32 {
	if versionObjectVersion == 0 {
		versionObjectVersion = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectVersion
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectVersion) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectVersion) GetJSON() (string, error) {
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
func (m *ObjectVersion) ReadObject() error {
	if Crud_ObjectVersion == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectVersion.ReadObject(m)

	return err
}
