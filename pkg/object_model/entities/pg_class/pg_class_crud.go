//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package pg_class

import (
	"strconv"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPgClass - версия структуры модели, с учётом имен и типов полей
var versionPgClass uint32

// Crud_PgClass - объект контроллер crud операций
var Crud_PgClass ICrud_PgClass

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PgClass interface {
	Read(*PgClass) error
	Save(*PgClass) error
	Update(*PgClass) error
	Create(*PgClass) error
	ReadFromCache(Oid int64, VersionID int64) (PgClass, error)
	UpdateManyFields(*PgClass, []string) error
	Update_Oid(*PgClass) error
	Update_Relallvisible(*PgClass) error
	Update_Relam(*PgClass) error
	Update_Relchecks(*PgClass) error
	Update_Relfilenode(*PgClass) error
	Update_Relforcerowsecurity(*PgClass) error
	Update_Relfrozenxid(*PgClass) error
	Update_Relhasindex(*PgClass) error
	Update_Relhasrules(*PgClass) error
	Update_Relhassubclass(*PgClass) error
	Update_Relhastriggers(*PgClass) error
	Update_Relispartition(*PgClass) error
	Update_Relispopulated(*PgClass) error
	Update_Relisshared(*PgClass) error
	Update_Relkind(*PgClass) error
	Update_Relminmxid(*PgClass) error
	Update_Relname(*PgClass) error
	Update_Relnamespace(*PgClass) error
	Update_Relnatts(*PgClass) error
	Update_Reloftype(*PgClass) error
	Update_Relowner(*PgClass) error
	Update_Relpages(*PgClass) error
	Update_Relpersistence(*PgClass) error
	Update_Relreplident(*PgClass) error
	Update_Relrewrite(*PgClass) error
	Update_Relrowsecurity(*PgClass) error
	Update_Reltablespace(*PgClass) error
	Update_Reltoastrelid(*PgClass) error
	Update_Reltuples(*PgClass) error
	Update_Reltype(*PgClass) error
	Update_VersionID(*PgClass) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PgClass) TableNameDB() string {
	return "pg_class"
}

// NewPgClass - возвращает новый	объект
func NewPgClass() PgClass {
	return PgClass{}
}

// AsPgClass - создаёт объект из упакованного объекта в массиве байтов
func AsPgClass(b []byte) (PgClass, error) {
	c := NewPgClass()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPgClass(), err
	}
	return c, nil
}

// PgClassAsBytes - упаковывает объект в массив байтов
func PgClassAsBytes(m *PgClass) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PgClass) GetStructVersion() uint32 {
	if versionPgClass == 0 {
		versionPgClass = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPgClass
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PgClass) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PgClass) GetJSON() (string, error) {
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
func (m *PgClass) Read() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgClass.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PgClass) Save() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgClass.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PgClass) Update() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgClass.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PgClass) Create() error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgClass.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PgClass) ReadFromCache(Oid int64, VersionID int64) (PgClass, error) {
	Otvet := PgClass{}
	var err error

	if Crud_PgClass == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PgClass.ReadFromCache(Oid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgClass) SetCrudInterface(crud ICrud_PgClass) {
	Crud_PgClass = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PgClass) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PgClass == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgClass.UpdateManyFields(m, MassNeedUpdateFields)

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