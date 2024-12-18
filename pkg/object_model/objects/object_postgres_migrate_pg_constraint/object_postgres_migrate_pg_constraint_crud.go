package object_postgres_migrate_pg_constraint

import (
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"reflect"
)

// versionPostgresMigratePgConstraint - версия структуры модели, с учётом имен и типов полей
var versionObjectPostgresMigratePgConstraint uint32

// Crud_PostgresMigratePgConstraint - объект контроллер crud операций
var Crud_ObjectPostgresMigratePgConstraint ICrud_ObjectPostgresMigratePgConstraint

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPostgresMigratePgConstraint interface {
	ICrud_ObjectPostgresMigratePgConstraint_manual
	ReadObject(*ObjectPostgresMigratePgConstraint) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPostgresMigratePgConstraint) GetStructVersion() uint32 {
	if versionObjectPostgresMigratePgConstraint == 0 {
		versionObjectPostgresMigratePgConstraint = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPostgresMigratePgConstraint
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPostgresMigratePgConstraint) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPostgresMigratePgConstraint) GetJSON() (string, error) {
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
func (m *ObjectPostgresMigratePgConstraint) ReadObject() error {
	if Crud_ObjectPostgresMigratePgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPostgresMigratePgConstraint.ReadObject(m)

	return err
}
