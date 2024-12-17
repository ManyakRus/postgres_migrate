//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package pg_namespace

import (
	"strconv"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPgNamespace - версия структуры модели, с учётом имен и типов полей
var versionPgNamespace uint32

// Crud_PgNamespace - объект контроллер crud операций
var Crud_PgNamespace ICrud_PgNamespace

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PgNamespace interface {
	Read(*PgNamespace) error
	Save(*PgNamespace) error
	Update(*PgNamespace) error
	Create(*PgNamespace) error
	ReadFromCache(Oid int64, VersionID int64) (PgNamespace, error)
	UpdateManyFields(*PgNamespace, []string) error
	Update_Nspacl(*PgNamespace) error
	Update_Nspname(*PgNamespace) error
	Update_Nspowner(*PgNamespace) error
	Update_Oid(*PgNamespace) error
	Update_VersionID(*PgNamespace) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PgNamespace) TableNameDB() string {
	return "pg_namespace"
}

// NewPgNamespace - возвращает новый	объект
func NewPgNamespace() PgNamespace {
	return PgNamespace{}
}

// AsPgNamespace - создаёт объект из упакованного объекта в массиве байтов
func AsPgNamespace(b []byte) (PgNamespace, error) {
	c := NewPgNamespace()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPgNamespace(), err
	}
	return c, nil
}

// PgNamespaceAsBytes - упаковывает объект в массив байтов
func PgNamespaceAsBytes(m *PgNamespace) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PgNamespace) GetStructVersion() uint32 {
	if versionPgNamespace == 0 {
		versionPgNamespace = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPgNamespace
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PgNamespace) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PgNamespace) GetJSON() (string, error) {
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
func (m *PgNamespace) Read() error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgNamespace.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PgNamespace) Save() error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgNamespace.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PgNamespace) Update() error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgNamespace.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PgNamespace) Create() error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgNamespace.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PgNamespace) ReadFromCache(Oid int64, VersionID int64) (PgNamespace, error) {
	Otvet := PgNamespace{}
	var err error

	if Crud_PgNamespace == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PgNamespace.ReadFromCache(Oid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgNamespace) SetCrudInterface(crud ICrud_PgNamespace) {
	Crud_PgNamespace = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PgNamespace) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PgNamespace == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgNamespace.UpdateManyFields(m, MassNeedUpdateFields)

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