//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package pg_description

import (
	"strconv"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPgDescription - версия структуры модели, с учётом имен и типов полей
var versionPgDescription uint32

// Crud_PgDescription - объект контроллер crud операций
var Crud_PgDescription ICrud_PgDescription

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PgDescription interface {
	Read(*PgDescription) error
	Save(*PgDescription) error
	Update(*PgDescription) error
	Create(*PgDescription) error
	ReadFromCache(Classoid int64, Objoid int64, Objsubid int32, VersionID int64) (PgDescription, error)
	UpdateManyFields(*PgDescription, []string) error
	Update_Classoid(*PgDescription) error
	Update_Description(*PgDescription) error
	Update_Objoid(*PgDescription) error
	Update_Objsubid(*PgDescription) error
	Update_VersionID(*PgDescription) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PgDescription) TableNameDB() string {
	return "pg_description"
}

// NewPgDescription - возвращает новый	объект
func NewPgDescription() PgDescription {
	return PgDescription{}
}

// AsPgDescription - создаёт объект из упакованного объекта в массиве байтов
func AsPgDescription(b []byte) (PgDescription, error) {
	c := NewPgDescription()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPgDescription(), err
	}
	return c, nil
}

// PgDescriptionAsBytes - упаковывает объект в массив байтов
func PgDescriptionAsBytes(m *PgDescription) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PgDescription) GetStructVersion() uint32 {
	if versionPgDescription == 0 {
		versionPgDescription = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPgDescription
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PgDescription) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PgDescription) GetJSON() (string, error) {
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
func (m *PgDescription) Read() error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgDescription.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PgDescription) Save() error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgDescription.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PgDescription) Update() error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgDescription.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PgDescription) Create() error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgDescription.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PgDescription) ReadFromCache(Classoid int64, Objoid int64, Objsubid int32, VersionID int64) (PgDescription, error) {
	Otvet := PgDescription{}
	var err error

	if Crud_PgDescription == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PgDescription.ReadFromCache(Classoid, Objoid, Objsubid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgDescription) SetCrudInterface(crud ICrud_PgDescription) {
	Crud_PgDescription = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PgDescription) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgDescription.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------

// StringIdentifier - возвращает строковое представление PrimaryKey
func StringIdentifier(Classoid int64, Objoid int64, Objsubid int32, VersionID int64) string {
	Otvet := ""
	Otvet = Otvet + strconv.Itoa(int(Classoid)) + "_"
	Otvet = Otvet + strconv.Itoa(int(Objoid)) + "_"
	Otvet = Otvet + strconv.Itoa(int(Objsubid)) + "_"
	Otvet = Otvet + strconv.Itoa(int(VersionID)) + "_"

	return Otvet
}