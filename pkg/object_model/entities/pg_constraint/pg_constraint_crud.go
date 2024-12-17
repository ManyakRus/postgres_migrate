//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package pg_constraint

import (
	"strconv"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

// versionPgConstraint - версия структуры модели, с учётом имен и типов полей
var versionPgConstraint uint32

// Crud_PgConstraint - объект контроллер crud операций
var Crud_PgConstraint ICrud_PgConstraint

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PgConstraint interface {
	Read(*PgConstraint) error
	Save(*PgConstraint) error
	Update(*PgConstraint) error
	Create(*PgConstraint) error
	ReadFromCache(Oid int64, VersionID int64) (PgConstraint, error)
	UpdateManyFields(*PgConstraint, []string) error
	Update_Condeferrable(*PgConstraint) error
	Update_Condeferred(*PgConstraint) error
	Update_Conexclop(*PgConstraint) error
	Update_Confdeltype(*PgConstraint) error
	Update_Conffeqop(*PgConstraint) error
	Update_Confkey(*PgConstraint) error
	Update_Confmatchtype(*PgConstraint) error
	Update_Confrelid(*PgConstraint) error
	Update_Confupdtype(*PgConstraint) error
	Update_Conindid(*PgConstraint) error
	Update_Coninhcount(*PgConstraint) error
	Update_Conislocal(*PgConstraint) error
	Update_Conkey(*PgConstraint) error
	Update_Conname(*PgConstraint) error
	Update_Connamespace(*PgConstraint) error
	Update_Connoinherit(*PgConstraint) error
	Update_Conparentid(*PgConstraint) error
	Update_Conpfeqop(*PgConstraint) error
	Update_Conppeqop(*PgConstraint) error
	Update_Conrelid(*PgConstraint) error
	Update_Contype(*PgConstraint) error
	Update_Contypid(*PgConstraint) error
	Update_Convalidated(*PgConstraint) error
	Update_Oid(*PgConstraint) error
	Update_VersionID(*PgConstraint) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PgConstraint) TableNameDB() string {
	return "pg_constraint"
}

// NewPgConstraint - возвращает новый	объект
func NewPgConstraint() PgConstraint {
	return PgConstraint{}
}

// AsPgConstraint - создаёт объект из упакованного объекта в массиве байтов
func AsPgConstraint(b []byte) (PgConstraint, error) {
	c := NewPgConstraint()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPgConstraint(), err
	}
	return c, nil
}

// PgConstraintAsBytes - упаковывает объект в массив байтов
func PgConstraintAsBytes(m *PgConstraint) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PgConstraint) GetStructVersion() uint32 {
	if versionPgConstraint == 0 {
		versionPgConstraint = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPgConstraint
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PgConstraint) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PgConstraint) GetJSON() (string, error) {
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
func (m *PgConstraint) Read() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgConstraint.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PgConstraint) Save() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgConstraint.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PgConstraint) Update() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgConstraint.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PgConstraint) Create() error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}
	
	err := Crud_PgConstraint.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PgConstraint) ReadFromCache(Oid int64, VersionID int64) (PgConstraint, error) {
	Otvet := PgConstraint{}
	var err error

	if Crud_PgConstraint == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PgConstraint.ReadFromCache(Oid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PgConstraint) SetCrudInterface(crud ICrud_PgConstraint) {
	Crud_PgConstraint = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PgConstraint) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PgConstraint.UpdateManyFields(m, MassNeedUpdateFields)

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