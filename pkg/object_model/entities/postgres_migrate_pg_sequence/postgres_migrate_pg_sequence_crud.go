//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package postgres_migrate_pg_sequence

import (
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
	"strconv"
)

// versionPostgresMigratePgSequence - версия структуры модели, с учётом имен и типов полей
var versionPostgresMigratePgSequence uint32

// Crud_PostgresMigratePgSequence - объект контроллер crud операций
var Crud_PostgresMigratePgSequence ICrud_PostgresMigratePgSequence

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PostgresMigratePgSequence interface {
	Read(*PostgresMigratePgSequence) error
	Save(*PostgresMigratePgSequence) error
	Update(*PostgresMigratePgSequence) error
	Create(*PostgresMigratePgSequence) error
	Delete(*PostgresMigratePgSequence) error
	Restore(*PostgresMigratePgSequence) error
	ReadFromCache(Seqrelid int64, VersionID int64) (PostgresMigratePgSequence, error)
	UpdateManyFields(*PostgresMigratePgSequence, []string) error
	Update_Seqcache(*PostgresMigratePgSequence) error
	Update_Seqcycle(*PostgresMigratePgSequence) error
	Update_Seqincrement(*PostgresMigratePgSequence) error
	Update_Seqmax(*PostgresMigratePgSequence) error
	Update_Seqmin(*PostgresMigratePgSequence) error
	Update_Seqrelid(*PostgresMigratePgSequence) error
	Update_Seqstart(*PostgresMigratePgSequence) error
	Update_Seqtypid(*PostgresMigratePgSequence) error
	Update_VersionID(*PostgresMigratePgSequence) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PostgresMigratePgSequence) TableNameDB() string {
	return "postgres_migrate_pg_sequence"
}

// NewPostgresMigratePgSequence - возвращает новый	объект
func NewPostgresMigratePgSequence() PostgresMigratePgSequence {
	return PostgresMigratePgSequence{}
}

// AsPostgresMigratePgSequence - создаёт объект из упакованного объекта в массиве байтов
func AsPostgresMigratePgSequence(b []byte) (PostgresMigratePgSequence, error) {
	c := NewPostgresMigratePgSequence()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPostgresMigratePgSequence(), err
	}
	return c, nil
}

// PostgresMigratePgSequenceAsBytes - упаковывает объект в массив байтов
func PostgresMigratePgSequenceAsBytes(m *PostgresMigratePgSequence) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PostgresMigratePgSequence) GetStructVersion() uint32 {
	if versionPostgresMigratePgSequence == 0 {
		versionPostgresMigratePgSequence = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPostgresMigratePgSequence
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PostgresMigratePgSequence) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PostgresMigratePgSequence) GetJSON() (string, error) {
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
func (m *PostgresMigratePgSequence) Read() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PostgresMigratePgSequence) Save() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PostgresMigratePgSequence) Update() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PostgresMigratePgSequence) Create() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *PostgresMigratePgSequence) Delete() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *PostgresMigratePgSequence) Restore() error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.Restore(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PostgresMigratePgSequence) ReadFromCache(Seqrelid int64, VersionID int64) (PostgresMigratePgSequence, error) {
	Otvet := PostgresMigratePgSequence{}
	var err error

	if Crud_PostgresMigratePgSequence == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PostgresMigratePgSequence.ReadFromCache(Seqrelid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgSequence) SetCrudInterface(crud ICrud_PostgresMigratePgSequence) {
	Crud_PostgresMigratePgSequence = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PostgresMigratePgSequence) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PostgresMigratePgSequence == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgSequence.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------

// StringIdentifier - возвращает строковое представление PrimaryKey
func StringIdentifier(Seqrelid int64, VersionID int64) string {
	Otvet := ""
	Otvet = Otvet + strconv.Itoa(int(Seqrelid)) + "_"
	Otvet = Otvet + strconv.Itoa(int(VersionID)) + "_"

	return Otvet
}
