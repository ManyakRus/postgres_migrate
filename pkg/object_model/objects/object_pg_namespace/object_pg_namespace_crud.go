package object_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionPgNamespace - версия структуры модели, с учётом имен и типов полей
var versionObjectPgNamespace uint32

// Crud_PgNamespace - объект контроллер crud операций
var Crud_ObjectPgNamespace ICrud_ObjectPgNamespace

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPgNamespace interface {
	ICrud_ObjectPgNamespace_manual
	ReadObject(*ObjectPgNamespace) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPgNamespace) GetStructVersion() uint32 {
	if versionObjectPgNamespace == 0 {
		versionObjectPgNamespace = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPgNamespace
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPgNamespace) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPgNamespace) GetJSON() (string, error) {
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
func (m *ObjectPgNamespace) ReadObject() error {
	if Crud_ObjectPgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPgNamespace.ReadObject(m)

	return err
}
