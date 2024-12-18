//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package postgres_migrate_pg_class

import (
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
	"strconv"
)

// versionPostgresMigratePgClass - версия структуры модели, с учётом имен и типов полей
var versionPostgresMigratePgClass uint32

// Crud_PostgresMigratePgClass - объект контроллер crud операций
var Crud_PostgresMigratePgClass ICrud_PostgresMigratePgClass

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PostgresMigratePgClass interface {
	Read(*PostgresMigratePgClass) error
	Save(*PostgresMigratePgClass) error
	Update(*PostgresMigratePgClass) error
	Create(*PostgresMigratePgClass) error
	ReadFromCache(Oid int64, VersionID int64) (PostgresMigratePgClass, error)
	UpdateManyFields(*PostgresMigratePgClass, []string) error
	Update_Oid(*PostgresMigratePgClass) error
	Update_Relallvisible(*PostgresMigratePgClass) error
	Update_Relam(*PostgresMigratePgClass) error
	Update_Relchecks(*PostgresMigratePgClass) error
	Update_Relfilenode(*PostgresMigratePgClass) error
	Update_Relforcerowsecurity(*PostgresMigratePgClass) error
	Update_Relfrozenxid(*PostgresMigratePgClass) error
	Update_Relhasindex(*PostgresMigratePgClass) error
	Update_Relhasrules(*PostgresMigratePgClass) error
	Update_Relhassubclass(*PostgresMigratePgClass) error
	Update_Relhastriggers(*PostgresMigratePgClass) error
	Update_Relispartition(*PostgresMigratePgClass) error
	Update_Relispopulated(*PostgresMigratePgClass) error
	Update_Relisshared(*PostgresMigratePgClass) error
	Update_Relkind(*PostgresMigratePgClass) error
	Update_Relminmxid(*PostgresMigratePgClass) error
	Update_Relname(*PostgresMigratePgClass) error
	Update_Relnamespace(*PostgresMigratePgClass) error
	Update_Relnatts(*PostgresMigratePgClass) error
	Update_Reloftype(*PostgresMigratePgClass) error
	Update_Relowner(*PostgresMigratePgClass) error
	Update_Relpages(*PostgresMigratePgClass) error
	Update_Relpersistence(*PostgresMigratePgClass) error
	Update_Relreplident(*PostgresMigratePgClass) error
	Update_Relrewrite(*PostgresMigratePgClass) error
	Update_Relrowsecurity(*PostgresMigratePgClass) error
	Update_Reltablespace(*PostgresMigratePgClass) error
	Update_Reltoastrelid(*PostgresMigratePgClass) error
	Update_Reltuples(*PostgresMigratePgClass) error
	Update_Reltype(*PostgresMigratePgClass) error
	Update_VersionID(*PostgresMigratePgClass) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PostgresMigratePgClass) TableNameDB() string {
	return "postgres_migrate_pg_class"
}

// NewPostgresMigratePgClass - возвращает новый	объект
func NewPostgresMigratePgClass() PostgresMigratePgClass {
	return PostgresMigratePgClass{}
}

// AsPostgresMigratePgClass - создаёт объект из упакованного объекта в массиве байтов
func AsPostgresMigratePgClass(b []byte) (PostgresMigratePgClass, error) {
	c := NewPostgresMigratePgClass()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPostgresMigratePgClass(), err
	}
	return c, nil
}

// PostgresMigratePgClassAsBytes - упаковывает объект в массив байтов
func PostgresMigratePgClassAsBytes(m *PostgresMigratePgClass) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PostgresMigratePgClass) GetStructVersion() uint32 {
	if versionPostgresMigratePgClass == 0 {
		versionPostgresMigratePgClass = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPostgresMigratePgClass
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PostgresMigratePgClass) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PostgresMigratePgClass) GetJSON() (string, error) {
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
func (m *PostgresMigratePgClass) Read() error {
	if Crud_PostgresMigratePgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgClass.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PostgresMigratePgClass) Save() error {
	if Crud_PostgresMigratePgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgClass.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PostgresMigratePgClass) Update() error {
	if Crud_PostgresMigratePgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgClass.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PostgresMigratePgClass) Create() error {
	if Crud_PostgresMigratePgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgClass.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PostgresMigratePgClass) ReadFromCache(Oid int64, VersionID int64) (PostgresMigratePgClass, error) {
	Otvet := PostgresMigratePgClass{}
	var err error

	if Crud_PostgresMigratePgClass == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PostgresMigratePgClass.ReadFromCache(Oid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgClass) SetCrudInterface(crud ICrud_PostgresMigratePgClass) {
	Crud_PostgresMigratePgClass = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PostgresMigratePgClass) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PostgresMigratePgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgClass.UpdateManyFields(m, MassNeedUpdateFields)

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
