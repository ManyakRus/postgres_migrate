package object_postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"encoding/json"
	"reflect"
)

// versionPostgresMigrateVersion - версия структуры модели, с учётом имен и типов полей
var versionObjectPostgresMigrateVersion uint32

// Crud_PostgresMigrateVersion - объект контроллер crud операций
var Crud_ObjectPostgresMigrateVersion ICrud_ObjectPostgresMigrateVersion

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_ObjectPostgresMigrateVersion interface {
	ICrud_ObjectPostgresMigrateVersion_manual
	ReadObject(*ObjectPostgresMigrateVersion) error
}

// GetStructVersion - возвращает версию модели
func (m ObjectPostgresMigrateVersion) GetStructVersion() uint32 {
	if versionObjectPostgresMigrateVersion == 0 {
		versionObjectPostgresMigrateVersion = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionObjectPostgresMigrateVersion
}

// GetModelFromJSON - создаёт модель из строки json
func (m *ObjectPostgresMigrateVersion) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m ObjectPostgresMigrateVersion) GetJSON() (string, error) {
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
func (m *ObjectPostgresMigrateVersion) ReadObject() error {
	if Crud_ObjectPostgresMigrateVersion == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_ObjectPostgresMigrateVersion.ReadObject(m)

	return err
}
