//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package postgres_migrate_pg_namespace

import (
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
	"strconv"
)

// versionPostgresMigratePgNamespace - версия структуры модели, с учётом имен и типов полей
var versionPostgresMigratePgNamespace uint32

// Crud_PostgresMigratePgNamespace - объект контроллер crud операций
var Crud_PostgresMigratePgNamespace ICrud_PostgresMigratePgNamespace

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PostgresMigratePgNamespace interface {
	Read(*PostgresMigratePgNamespace) error
	Save(*PostgresMigratePgNamespace) error
	Update(*PostgresMigratePgNamespace) error
	Create(*PostgresMigratePgNamespace) error
	Delete(*PostgresMigratePgNamespace) error
	Restore(*PostgresMigratePgNamespace) error
	ReadFromCache(Oid int64, VersionID int64) (PostgresMigratePgNamespace, error)
	UpdateManyFields(*PostgresMigratePgNamespace, []string) error
	Update_Nspacl(*PostgresMigratePgNamespace) error
	Update_Nspname(*PostgresMigratePgNamespace) error
	Update_Nspowner(*PostgresMigratePgNamespace) error
	Update_Oid(*PostgresMigratePgNamespace) error
	Update_VersionID(*PostgresMigratePgNamespace) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PostgresMigratePgNamespace) TableNameDB() string {
	return "postgres_migrate_pg_namespace"
}

// NewPostgresMigratePgNamespace - возвращает новый	объект
func NewPostgresMigratePgNamespace() PostgresMigratePgNamespace {
	return PostgresMigratePgNamespace{}
}

// AsPostgresMigratePgNamespace - создаёт объект из упакованного объекта в массиве байтов
func AsPostgresMigratePgNamespace(b []byte) (PostgresMigratePgNamespace, error) {
	c := NewPostgresMigratePgNamespace()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPostgresMigratePgNamespace(), err
	}
	return c, nil
}

// PostgresMigratePgNamespaceAsBytes - упаковывает объект в массив байтов
func PostgresMigratePgNamespaceAsBytes(m *PostgresMigratePgNamespace) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PostgresMigratePgNamespace) GetStructVersion() uint32 {
	if versionPostgresMigratePgNamespace == 0 {
		versionPostgresMigratePgNamespace = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPostgresMigratePgNamespace
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PostgresMigratePgNamespace) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PostgresMigratePgNamespace) GetJSON() (string, error) {
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
func (m *PostgresMigratePgNamespace) Read() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PostgresMigratePgNamespace) Save() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PostgresMigratePgNamespace) Update() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PostgresMigratePgNamespace) Create() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *PostgresMigratePgNamespace) Delete() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *PostgresMigratePgNamespace) Restore() error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.Restore(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PostgresMigratePgNamespace) ReadFromCache(Oid int64, VersionID int64) (PostgresMigratePgNamespace, error) {
	Otvet := PostgresMigratePgNamespace{}
	var err error

	if Crud_PostgresMigratePgNamespace == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PostgresMigratePgNamespace.ReadFromCache(Oid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgNamespace) SetCrudInterface(crud ICrud_PostgresMigratePgNamespace) {
	Crud_PostgresMigratePgNamespace = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PostgresMigratePgNamespace) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PostgresMigratePgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgNamespace.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------

// StringIdentifier - возвращает строковое представление PrimaryKey
func StringIdentifier(Oid int64, VersionID int64) string {
	Otvet := ""
	Otvet = Otvet + strconv.Itoa(int(Oid)) + "_"
	Otvet = Otvet + strconv.Itoa(int(VersionID)) + "_"

	return Otvet
}
