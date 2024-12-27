//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package postgres_migrate_pg_description

import (
	"strconv"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPostgresMigratePgDescription - версия структуры модели, с учётом имен и типов полей
var versionPostgresMigratePgDescription uint32

// Crud_PostgresMigratePgDescription - объект контроллер crud операций
var Crud_PostgresMigratePgDescription ICrud_PostgresMigratePgDescription

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PostgresMigratePgDescription interface {
	Read(*PostgresMigratePgDescription) error
	Save(*PostgresMigratePgDescription) error
	Update(*PostgresMigratePgDescription) error
	Create(*PostgresMigratePgDescription) error
	Delete(*PostgresMigratePgDescription) error
	Restore(*PostgresMigratePgDescription) error
	ReadFromCache(Classoid int64, Objoid int64, Objsubid int32, VersionID int64) (PostgresMigratePgDescription, error)
	UpdateManyFields(*PostgresMigratePgDescription, []string) error
	Update_Classoid(*PostgresMigratePgDescription) error
	Update_Description(*PostgresMigratePgDescription) error
	Update_Objoid(*PostgresMigratePgDescription) error
	Update_Objsubid(*PostgresMigratePgDescription) error
	Update_VersionID(*PostgresMigratePgDescription) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PostgresMigratePgDescription) TableNameDB() string {
	return "postgres_migrate_pg_description"
}

// NewPostgresMigratePgDescription - возвращает новый	объект
func NewPostgresMigratePgDescription() PostgresMigratePgDescription {
	return PostgresMigratePgDescription{}
}

// AsPostgresMigratePgDescription - создаёт объект из упакованного объекта в массиве байтов
func AsPostgresMigratePgDescription(b []byte) (PostgresMigratePgDescription, error) {
	c := NewPostgresMigratePgDescription()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPostgresMigratePgDescription(), err
	}
	return c, nil
}

// PostgresMigratePgDescriptionAsBytes - упаковывает объект в массив байтов
func PostgresMigratePgDescriptionAsBytes(m *PostgresMigratePgDescription) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PostgresMigratePgDescription) GetStructVersion() uint32 {
	if versionPostgresMigratePgDescription == 0 {
		versionPostgresMigratePgDescription = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPostgresMigratePgDescription
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PostgresMigratePgDescription) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PostgresMigratePgDescription) GetJSON() (string, error) {
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
func (m *PostgresMigratePgDescription) Read() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgDescription.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PostgresMigratePgDescription) Save() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgDescription.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PostgresMigratePgDescription) Update() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgDescription.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PostgresMigratePgDescription) Create() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgDescription.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *PostgresMigratePgDescription) Delete() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgDescription.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *PostgresMigratePgDescription) Restore() error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgDescription.Restore(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PostgresMigratePgDescription) ReadFromCache(Classoid int64, Objoid int64, Objsubid int32, VersionID int64) (PostgresMigratePgDescription, error) {
	Otvet := PostgresMigratePgDescription{}
	var err error

	if Crud_PostgresMigratePgDescription == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PostgresMigratePgDescription.ReadFromCache(Classoid, Objoid, Objsubid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgDescription) SetCrudInterface(crud ICrud_PostgresMigratePgDescription) {
	Crud_PostgresMigratePgDescription = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PostgresMigratePgDescription) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PostgresMigratePgDescription == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgDescription.UpdateManyFields(m, MassNeedUpdateFields)

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