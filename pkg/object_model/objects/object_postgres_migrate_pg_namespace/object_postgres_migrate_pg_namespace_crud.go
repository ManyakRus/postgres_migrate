package object_postgres_migrate_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionPostgresMigratePgNamespace - версия структуры модели, с учётом имен и типов полей
var versionObjectPostgresMigratePgNamespace uint32

// Crud_PostgresMigratePgNamespace - объект контроллер crud операций
var Crud_ObjectPostgresMigratePgNamespace ICrud_ObjectPostgresMigratePgNamespace

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPostgresMigratePgNamespace interface {
	ICrud_ObjectPostgresMigratePgNamespace_manual
	ReadObject(*ObjectPostgresMigratePgNamespace) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPostgresMigratePgNamespace) GetStructVersion() uint32 {
	if versionObjectPostgresMigratePgNamespace == 0 {
		versionObjectPostgresMigratePgNamespace = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPostgresMigratePgNamespace
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPostgresMigratePgNamespace) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPostgresMigratePgNamespace) GetJSON() (string, error) {
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
func (m *ObjectPostgresMigratePgNamespace) ReadObject() error {
	if Crud_ObjectPostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPostgresMigratePgNamespace.ReadObject(m)

	return err
}
