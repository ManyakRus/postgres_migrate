//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionVersion - версия структуры модели, с учётом имен и типов полей
var versionVersion uint32

// Crud_Version - объект контроллер crud операций
var Crud_Version ICrud_Version

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_Version interface {
	Read(*Version) error
	Save(*Version) error
	Update(*Version) error
	Create(*Version) error
	ReadFromCache(ID int64) (Version, error)
	UpdateManyFields(*Version, []string) error
	Update_Description(*Version) error
	Update_Name(*Version) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m Version) TableNameDB() string {
	return "version"
}

// NewVersion - возвращает новый	объект
func NewVersion() Version {
	return Version{}
}

// AsVersion - создаёт объект из упакованного объекта в массиве байтов
func AsVersion(b []byte) (Version, error) {
	c := NewVersion()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewVersion(), err
	}
	return c, nil
}

// VersionAsBytes - упаковывает объект в массив байтов
func VersionAsBytes(m *Version) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m Version) GetStructVersion() uint32 {
	if versionVersion == 0 {
		versionVersion = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionVersion
}

// GetModelFromJSON - создаёт модель из строки json
func (m *Version) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m Version) GetJSON() (string, error) {
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
func (m *Version) Read() error {
	if Crud_Version == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Version.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *Version) Save() error {
	if Crud_Version == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Version.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *Version) Update() error {
	if Crud_Version == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Version.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *Version) Create() error {
	if Crud_Version == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_Version.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *Version) ReadFromCache(ID int64) (Version, error) {
	Otvet := Version{}
	var err error

	if Crud_Version == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_Version.ReadFromCache(ID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m Version) SetCrudInterface(crud ICrud_Version) {
	Crud_Version = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *Version) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_Version == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_Version.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------
