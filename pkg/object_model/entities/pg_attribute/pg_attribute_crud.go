//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package pg_attribute

import (
	"strconv"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPgAttribute - версия структуры модели, с учётом имен и типов полей
var versionPgAttribute uint32

// Crud_PgAttribute - объект контроллер crud операций
var Crud_PgAttribute ICrud_PgAttribute

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PgAttribute interface {
	Read(*PgAttribute) error
	Save(*PgAttribute) error
	Update(*PgAttribute) error
	Create(*PgAttribute) error
	ReadFromCache(Attname string, Attrelid int64, VersionID int64) (PgAttribute, error)
	UpdateManyFields(*PgAttribute, []string) error
	Update_Attalign(*PgAttribute) error
	Update_Attbyval(*PgAttribute) error
	Update_Attcacheoff(*PgAttribute) error
	Update_Attcollation(*PgAttribute) error
	Update_Attgenerated(*PgAttribute) error
	Update_Atthasdef(*PgAttribute) error
	Update_Atthasmissing(*PgAttribute) error
	Update_Attidentity(*PgAttribute) error
	Update_Attinhcount(*PgAttribute) error
	Update_Attisdropped(*PgAttribute) error
	Update_Attislocal(*PgAttribute) error
	Update_Attlen(*PgAttribute) error
	Update_Attname(*PgAttribute) error
	Update_Attndims(*PgAttribute) error
	Update_Attnotnull(*PgAttribute) error
	Update_Attnum(*PgAttribute) error
	Update_Attrelid(*PgAttribute) error
	Update_Attstattarget(*PgAttribute) error
	Update_Attstorage(*PgAttribute) error
	Update_Atttypid(*PgAttribute) error
	Update_Atttypmod(*PgAttribute) error
	Update_VersionID(*PgAttribute) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PgAttribute) TableNameDB() string {
	return "pg_attribute"
}

// NewPgAttribute - возвращает новый	объект
func NewPgAttribute() PgAttribute {
	return PgAttribute{}
}

// AsPgAttribute - создаёт объект из упакованного объекта в массиве байтов
func AsPgAttribute(b []byte) (PgAttribute, error) {
	c := NewPgAttribute()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPgAttribute(), err
	}
	return c, nil
}

// PgAttributeAsBytes - упаковывает объект в массив байтов
func PgAttributeAsBytes(m *PgAttribute) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PgAttribute) GetStructVersion() uint32 {
	if versionPgAttribute == 0 {
		versionPgAttribute = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPgAttribute
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PgAttribute) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PgAttribute) GetJSON() (string, error) {
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
func (m *PgAttribute) Read() error {
	if Crud_PgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgAttribute.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PgAttribute) Save() error {
	if Crud_PgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgAttribute.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PgAttribute) Update() error {
	if Crud_PgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgAttribute.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PgAttribute) Create() error {
	if Crud_PgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgAttribute.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PgAttribute) ReadFromCache(Attname string, Attrelid int64, VersionID int64) (PgAttribute, error) {
	Otvet := PgAttribute{}
	var err error

	if Crud_PgAttribute == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PgAttribute.ReadFromCache(Attname, Attrelid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgAttribute) SetCrudInterface(crud ICrud_PgAttribute) {
	Crud_PgAttribute = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PgAttribute) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgAttribute.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------

// StringIdentifier - возвращает строковое представление PrimaryKey
func StringIdentifier(Attname string, Attrelid int64, VersionID int64) string {
	Otvet := ""
	Otvet = Otvet + Attname + "_"
	Otvet = Otvet + strconv.Itoa(int(Attrelid)) + "_"
	Otvet = Otvet + strconv.Itoa(int(VersionID)) + "_"

	return Otvet
}