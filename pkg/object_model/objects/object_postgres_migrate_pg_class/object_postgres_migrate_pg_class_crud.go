package object_postgres_migrate_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionPostgresMigratePgClass - версия структуры модели, с учётом имен и типов полей
var versionObjectPostgresMigratePgClass uint32

// Crud_PostgresMigratePgClass - объект контроллер crud операций
var Crud_ObjectPostgresMigratePgClass ICrud_ObjectPostgresMigratePgClass

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPostgresMigratePgClass interface {
	ICrud_ObjectPostgresMigratePgClass_manual
	ReadObject(*ObjectPostgresMigratePgClass) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPostgresMigratePgClass) GetStructVersion() uint32 {
	if versionObjectPostgresMigratePgClass == 0 {
		versionObjectPostgresMigratePgClass = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPostgresMigratePgClass
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPostgresMigratePgClass) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPostgresMigratePgClass) GetJSON() (string, error) {
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
func (m *ObjectPostgresMigratePgClass) ReadObject() error {
	if Crud_ObjectPostgresMigratePgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPostgresMigratePgClass.ReadObject(m)

	return err
}
