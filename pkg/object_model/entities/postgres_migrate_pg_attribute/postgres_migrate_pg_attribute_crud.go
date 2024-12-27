//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package postgres_migrate_pg_attribute

import (
	"strconv"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPostgresMigratePgAttribute - версия структуры модели, с учётом имен и типов полей
var versionPostgresMigratePgAttribute uint32

// Crud_PostgresMigratePgAttribute - объект контроллер crud операций
var Crud_PostgresMigratePgAttribute ICrud_PostgresMigratePgAttribute

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PostgresMigratePgAttribute interface {
	Read(*PostgresMigratePgAttribute) error
	Save(*PostgresMigratePgAttribute) error
	Update(*PostgresMigratePgAttribute) error
	Create(*PostgresMigratePgAttribute) error
	Delete(*PostgresMigratePgAttribute) error
	Restore(*PostgresMigratePgAttribute) error
	ReadFromCache(Attname string, Attrelid int64, VersionID int64) (PostgresMigratePgAttribute, error)
	UpdateManyFields(*PostgresMigratePgAttribute, []string) error
	Update_Attalign(*PostgresMigratePgAttribute) error
	Update_Attbyval(*PostgresMigratePgAttribute) error
	Update_Attcacheoff(*PostgresMigratePgAttribute) error
	Update_Attcollation(*PostgresMigratePgAttribute) error
	Update_Attgenerated(*PostgresMigratePgAttribute) error
	Update_Atthasdef(*PostgresMigratePgAttribute) error
	Update_Atthasmissing(*PostgresMigratePgAttribute) error
	Update_Attidentity(*PostgresMigratePgAttribute) error
	Update_Attinhcount(*PostgresMigratePgAttribute) error
	Update_Attisdropped(*PostgresMigratePgAttribute) error
	Update_Attislocal(*PostgresMigratePgAttribute) error
	Update_Attlen(*PostgresMigratePgAttribute) error
	Update_Attmissingval(*PostgresMigratePgAttribute) error
	Update_Attname(*PostgresMigratePgAttribute) error
	Update_Attndims(*PostgresMigratePgAttribute) error
	Update_Attnotnull(*PostgresMigratePgAttribute) error
	Update_Attnum(*PostgresMigratePgAttribute) error
	Update_Attrelid(*PostgresMigratePgAttribute) error
	Update_Attstattarget(*PostgresMigratePgAttribute) error
	Update_Attstorage(*PostgresMigratePgAttribute) error
	Update_Atttypid(*PostgresMigratePgAttribute) error
	Update_Atttypmod(*PostgresMigratePgAttribute) error
	Update_VersionID(*PostgresMigratePgAttribute) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PostgresMigratePgAttribute) TableNameDB() string {
	return "postgres_migrate_pg_attribute"
}

// NewPostgresMigratePgAttribute - возвращает новый	объект
func NewPostgresMigratePgAttribute() PostgresMigratePgAttribute {
	return PostgresMigratePgAttribute{}
}

// AsPostgresMigratePgAttribute - создаёт объект из упакованного объекта в массиве байтов
func AsPostgresMigratePgAttribute(b []byte) (PostgresMigratePgAttribute, error) {
	c := NewPostgresMigratePgAttribute()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPostgresMigratePgAttribute(), err
	}
	return c, nil
}

// PostgresMigratePgAttributeAsBytes - упаковывает объект в массив байтов
func PostgresMigratePgAttributeAsBytes(m *PostgresMigratePgAttribute) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PostgresMigratePgAttribute) GetStructVersion() uint32 {
	if versionPostgresMigratePgAttribute == 0 {
		versionPostgresMigratePgAttribute = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPostgresMigratePgAttribute
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PostgresMigratePgAttribute) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PostgresMigratePgAttribute) GetJSON() (string, error) {
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
func (m *PostgresMigratePgAttribute) Read() error {
	if Crud_PostgresMigratePgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgAttribute.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PostgresMigratePgAttribute) Save() error {
	if Crud_PostgresMigratePgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgAttribute.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PostgresMigratePgAttribute) Update() error {
	if Crud_PostgresMigratePgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgAttribute.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PostgresMigratePgAttribute) Create() error {
	if Crud_PostgresMigratePgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgAttribute.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *PostgresMigratePgAttribute) Delete() error {
	if Crud_PostgresMigratePgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgAttribute.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *PostgresMigratePgAttribute) Restore() error {
	if Crud_PostgresMigratePgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgAttribute.Restore(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PostgresMigratePgAttribute) ReadFromCache(Attname string, Attrelid int64, VersionID int64) (PostgresMigratePgAttribute, error) {
	Otvet := PostgresMigratePgAttribute{}
	var err error

	if Crud_PostgresMigratePgAttribute == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PostgresMigratePgAttribute.ReadFromCache(Attname, Attrelid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgAttribute) SetCrudInterface(crud ICrud_PostgresMigratePgAttribute) {
	Crud_PostgresMigratePgAttribute = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PostgresMigratePgAttribute) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PostgresMigratePgAttribute == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgAttribute.UpdateManyFields(m, MassNeedUpdateFields)

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