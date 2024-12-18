//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package postgres_migrate_version

import (
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPostgresMigrateVersion - версия структуры модели, с учётом имен и типов полей
var versionPostgresMigrateVersion uint32

// Crud_PostgresMigrateVersion - объект контроллер crud операций
var Crud_PostgresMigrateVersion ICrud_PostgresMigrateVersion

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PostgresMigrateVersion interface {
	Read(*PostgresMigrateVersion) error
	Save(*PostgresMigrateVersion) error
	Update(*PostgresMigrateVersion) error
	Create(*PostgresMigrateVersion) error
	ReadFromCache(ID int64) (PostgresMigrateVersion, error)
	UpdateManyFields(*PostgresMigrateVersion, []string) error
	Update_Description(*PostgresMigrateVersion) error
	Update_Name(*PostgresMigrateVersion) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PostgresMigrateVersion) TableNameDB() string {
	return "postgres_migrate_version"
}

// NewPostgresMigrateVersion - возвращает новый	объект
func NewPostgresMigrateVersion() PostgresMigrateVersion {
	return PostgresMigrateVersion{}
}

// AsPostgresMigrateVersion - создаёт объект из упакованного объекта в массиве байтов
func AsPostgresMigrateVersion(b []byte) (PostgresMigrateVersion, error) {
	c := NewPostgresMigrateVersion()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPostgresMigrateVersion(), err
	}
	return c, nil
}

// PostgresMigrateVersionAsBytes - упаковывает объект в массив байтов
func PostgresMigrateVersionAsBytes(m *PostgresMigrateVersion) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PostgresMigrateVersion) GetStructVersion() uint32 {
	if versionPostgresMigrateVersion == 0 {
		versionPostgresMigrateVersion = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPostgresMigrateVersion
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PostgresMigrateVersion) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PostgresMigrateVersion) GetJSON() (string, error) {
	var ReturnVar string
	var err error

	bytes, err := json.Marshal(m)
	if err != nil {
		return ReturnVar, err
	}
	ReturnVar = string(bytes)
	return ReturnVar, err
}

// ---------------------------- CRUD операции ------------------------------------------------------------

// Read - находит запись в БД по ID, и заполняет в объект
func (m *PostgresMigrateVersion) Read() error {
	if Crud_PostgresMigrateVersion == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigrateVersion.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PostgresMigrateVersion) Save() error {
	if Crud_PostgresMigrateVersion == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigrateVersion.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PostgresMigrateVersion) Update() error {
	if Crud_PostgresMigrateVersion == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigrateVersion.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PostgresMigrateVersion) Create() error {
	if Crud_PostgresMigrateVersion == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigrateVersion.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PostgresMigrateVersion) ReadFromCache(ID int64) (PostgresMigrateVersion, error) {
	Otvet := PostgresMigrateVersion{}
	var err error

	if Crud_PostgresMigrateVersion == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PostgresMigrateVersion.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigrateVersion) SetCrudInterface(crud ICrud_PostgresMigrateVersion) {
	Crud_PostgresMigrateVersion = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PostgresMigrateVersion) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PostgresMigrateVersion == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigrateVersion.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
