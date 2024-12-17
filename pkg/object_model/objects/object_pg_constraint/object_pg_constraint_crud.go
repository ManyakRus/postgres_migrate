package object_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionPgConstraint - версия структуры модели, с учётом имен и типов полей
var versionObjectPgConstraint uint32

// Crud_PgConstraint - объект контроллер crud операций
var Crud_ObjectPgConstraint ICrud_ObjectPgConstraint

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPgConstraint interface {
	ICrud_ObjectPgConstraint_manual
	ReadObject(*ObjectPgConstraint) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPgConstraint) GetStructVersion() uint32 {
	if versionObjectPgConstraint == 0 {
		versionObjectPgConstraint = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPgConstraint
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPgConstraint) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPgConstraint) GetJSON() (string, error) {
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
func (m *ObjectPgConstraint) ReadObject() error {
	if Crud_ObjectPgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPgConstraint.ReadObject(m)

	return err
}
