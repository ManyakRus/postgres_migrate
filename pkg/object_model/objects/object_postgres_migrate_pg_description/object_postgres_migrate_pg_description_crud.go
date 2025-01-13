package object_postgres_migrate_pg_description

import (
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"reflect"
)

// versionPostgresMigratePgDescription - версия структуры модели, с учётом имен и типов полей
var versionObjectPostgresMigratePgDescription uint32

// Crud_PostgresMigratePgDescription - объект контроллер crud операций
var Crud_ObjectPostgresMigratePgDescription ICrud_ObjectPostgresMigratePgDescription

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPostgresMigratePgDescription interface {
	ICrud_ObjectPostgresMigratePgDescription_manual
	ReadObject(*ObjectPostgresMigratePgDescription) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPostgresMigratePgDescription) GetStructVersion() uint32 {
	if versionObjectPostgresMigratePgDescription == 0 {
		versionObjectPostgresMigratePgDescription = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPostgresMigratePgDescription
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPostgresMigratePgDescription) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPostgresMigratePgDescription) GetJSON() (string, error) {
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
func (m *ObjectPostgresMigratePgDescription) ReadObject() error {
	if Crud_ObjectPostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPostgresMigratePgDescription.ReadObject(m)

	return err
}
