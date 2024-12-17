//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package pg_index

import (
	"strconv"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPgIndex - версия структуры модели, с учётом имен и типов полей
var versionPgIndex uint32

// Crud_PgIndex - объект контроллер crud операций
var Crud_PgIndex ICrud_PgIndex

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PgIndex interface {
	Read(*PgIndex) error
	Save(*PgIndex) error
	Update(*PgIndex) error
	Create(*PgIndex) error
	ReadFromCache(Indexrelid int64, Indrelid int64, VersionID int64) (PgIndex, error)
	UpdateManyFields(*PgIndex, []string) error
	Update_Indcheckxmin(*PgIndex) error
	Update_Indclass(*PgIndex) error
	Update_Indcollation(*PgIndex) error
	Update_Indexprs(*PgIndex) error
	Update_Indexrelid(*PgIndex) error
	Update_Indimmediate(*PgIndex) error
	Update_Indisclustered(*PgIndex) error
	Update_Indisexclusion(*PgIndex) error
	Update_Indislive(*PgIndex) error
	Update_Indisprimary(*PgIndex) error
	Update_Indisready(*PgIndex) error
	Update_Indisreplident(*PgIndex) error
	Update_Indisunique(*PgIndex) error
	Update_Indisvalid(*PgIndex) error
	Update_Indkey(*PgIndex) error
	Update_Indnatts(*PgIndex) error
	Update_Indnkeyatts(*PgIndex) error
	Update_Indoption(*PgIndex) error
	Update_Indpred(*PgIndex) error
	Update_Indrelid(*PgIndex) error
	Update_VersionID(*PgIndex) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PgIndex) TableNameDB() string {
	return "pg_index"
}

// NewPgIndex - возвращает новый	объект
func NewPgIndex() PgIndex {
	return PgIndex{}
}

// AsPgIndex - создаёт объект из упакованного объекта в массиве байтов
func AsPgIndex(b []byte) (PgIndex, error) {
	c := NewPgIndex()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPgIndex(), err
	}
	return c, nil
}

// PgIndexAsBytes - упаковывает объект в массив байтов
func PgIndexAsBytes(m *PgIndex) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PgIndex) GetStructVersion() uint32 {
	if versionPgIndex == 0 {
		versionPgIndex = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPgIndex
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PgIndex) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PgIndex) GetJSON() (string, error) {
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
func (m *PgIndex) Read() error {
	if Crud_PgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgIndex.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PgIndex) Save() error {
	if Crud_PgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgIndex.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PgIndex) Update() error {
	if Crud_PgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgIndex.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PgIndex) Create() error {
	if Crud_PgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgIndex.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PgIndex) ReadFromCache(Indexrelid int64, Indrelid int64, VersionID int64) (PgIndex, error) {
	Otvet := PgIndex{}
	var err error

	if Crud_PgIndex == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PgIndex.ReadFromCache(Indexrelid, Indrelid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgIndex) SetCrudInterface(crud ICrud_PgIndex) {
	Crud_PgIndex = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PgIndex) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PgIndex == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgIndex.UpdateManyFields(m, MassNeedUpdateFields)

	return err
}

// ---------------------------- конец CRUD операции ------------------------------------------------------------

// StringIdentifier - возвращает строковое представление PrimaryKey
func StringIdentifier(Indexrelid int64, Indrelid int64, VersionID int64) string {
	Otvet := ""
	Otvet = Otvet + strconv.Itoa(int(Indexrelid)) + "_"
	Otvet = Otvet + strconv.Itoa(int(Indrelid)) + "_"
	Otvet = Otvet + strconv.Itoa(int(VersionID)) + "_"

	return Otvet
}