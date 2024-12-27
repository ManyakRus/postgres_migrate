//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package postgres_migrate_pg_index

import (
	"strconv"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPostgresMigratePgIndex - версия структуры модели, с учётом имен и типов полей
var versionPostgresMigratePgIndex uint32

// Crud_PostgresMigratePgIndex - объект контроллер crud операций
var Crud_PostgresMigratePgIndex ICrud_PostgresMigratePgIndex

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PostgresMigratePgIndex interface {
	Read(*PostgresMigratePgIndex) error
	Save(*PostgresMigratePgIndex) error
	Update(*PostgresMigratePgIndex) error
	Create(*PostgresMigratePgIndex) error
	Delete(*PostgresMigratePgIndex) error
	Restore(*PostgresMigratePgIndex) error
	ReadFromCache(Indexrelid int64, VersionID int64) (PostgresMigratePgIndex, error)
	UpdateManyFields(*PostgresMigratePgIndex, []string) error
	Update_Indcheckxmin(*PostgresMigratePgIndex) error
	Update_Indclass(*PostgresMigratePgIndex) error
	Update_Indcollation(*PostgresMigratePgIndex) error
	Update_Indexprs(*PostgresMigratePgIndex) error
	Update_Indexrelid(*PostgresMigratePgIndex) error
	Update_Indimmediate(*PostgresMigratePgIndex) error
	Update_Indisclustered(*PostgresMigratePgIndex) error
	Update_Indisexclusion(*PostgresMigratePgIndex) error
	Update_Indislive(*PostgresMigratePgIndex) error
	Update_Indisprimary(*PostgresMigratePgIndex) error
	Update_Indisready(*PostgresMigratePgIndex) error
	Update_Indisreplident(*PostgresMigratePgIndex) error
	Update_Indisunique(*PostgresMigratePgIndex) error
	Update_Indisvalid(*PostgresMigratePgIndex) error
	Update_Indkey(*PostgresMigratePgIndex) error
	Update_Indnatts(*PostgresMigratePgIndex) error
	Update_Indnkeyatts(*PostgresMigratePgIndex) error
	Update_Indoption(*PostgresMigratePgIndex) error
	Update_Indpred(*PostgresMigratePgIndex) error
	Update_Indrelid(*PostgresMigratePgIndex) error
	Update_VersionID(*PostgresMigratePgIndex) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PostgresMigratePgIndex) TableNameDB() string {
	return "postgres_migrate_pg_index"
}

// NewPostgresMigratePgIndex - возвращает новый	объект
func NewPostgresMigratePgIndex() PostgresMigratePgIndex {
	return PostgresMigratePgIndex{}
}

// AsPostgresMigratePgIndex - создаёт объект из упакованного объекта в массиве байтов
func AsPostgresMigratePgIndex(b []byte) (PostgresMigratePgIndex, error) {
	c := NewPostgresMigratePgIndex()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPostgresMigratePgIndex(), err
	}
	return c, nil
}

// PostgresMigratePgIndexAsBytes - упаковывает объект в массив байтов
func PostgresMigratePgIndexAsBytes(m *PostgresMigratePgIndex) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PostgresMigratePgIndex) GetStructVersion() uint32 {
	if versionPostgresMigratePgIndex == 0 {
		versionPostgresMigratePgIndex = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPostgresMigratePgIndex
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PostgresMigratePgIndex) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PostgresMigratePgIndex) GetJSON() (string, error) {
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
func (m *PostgresMigratePgIndex) Read() error {
	if Crud_PostgresMigratePgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgIndex.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PostgresMigratePgIndex) Save() error {
	if Crud_PostgresMigratePgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgIndex.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PostgresMigratePgIndex) Update() error {
	if Crud_PostgresMigratePgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgIndex.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PostgresMigratePgIndex) Create() error {
	if Crud_PostgresMigratePgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgIndex.Create(m)

	return err
}

// Delete - устанавливает признак пометки удаления в БД
func (m *PostgresMigratePgIndex) Delete() error {
	if Crud_PostgresMigratePgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgIndex.Delete(m)

	return err
}

// Restore - снимает признак пометки удаления в БД
func (m *PostgresMigratePgIndex) Restore() error {
	if Crud_PostgresMigratePgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PostgresMigratePgIndex.Restore(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PostgresMigratePgIndex) ReadFromCache(Indexrelid int64, VersionID int64) (PostgresMigratePgIndex, error) {
	Otvet := PostgresMigratePgIndex{}
	var err error

	if Crud_PostgresMigratePgIndex == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PostgresMigratePgIndex.ReadFromCache(Indexrelid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgIndex) SetCrudInterface(crud ICrud_PostgresMigratePgIndex) {
	Crud_PostgresMigratePgIndex = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PostgresMigratePgIndex) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PostgresMigratePgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgIndex.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------

// StringIdentifier - возвращает строковое представление PrimaryKey
func StringIdentifier(Indexrelid int64, VersionID int64) string {
	Otvet := ""
	Otvet = Otvet + strconv.Itoa(int(Indexrelid)) + "_"
	Otvet = Otvet + strconv.Itoa(int(VersionID)) + "_"

	return Otvet
}