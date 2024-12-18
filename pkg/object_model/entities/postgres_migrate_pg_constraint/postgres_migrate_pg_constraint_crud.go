//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package postgres_migrate_pg_constraint

import (
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/pkg/db/calc_struct_version"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
	"strconv"
)

// versionPostgresMigratePgConstraint - версия структуры модели, с учётом имен и типов полей
var versionPostgresMigratePgConstraint uint32

// Crud_PostgresMigratePgConstraint - объект контроллер crud операций
var Crud_PostgresMigratePgConstraint ICrud_PostgresMigratePgConstraint

// интерфейс стандартных CRUD операций, для использования в DB или GRPC или NRPC
type ICrud_PostgresMigratePgConstraint interface {
	Read(*PostgresMigratePgConstraint) error
	Save(*PostgresMigratePgConstraint) error
	Update(*PostgresMigratePgConstraint) error
	Create(*PostgresMigratePgConstraint) error
	ReadFromCache(Oid int64, VersionID int64) (PostgresMigratePgConstraint, error)
	UpdateManyFields(*PostgresMigratePgConstraint, []string) error
	Update_Condeferrable(*PostgresMigratePgConstraint) error
	Update_Condeferred(*PostgresMigratePgConstraint) error
	Update_Conexclop(*PostgresMigratePgConstraint) error
	Update_Confdeltype(*PostgresMigratePgConstraint) error
	Update_Conffeqop(*PostgresMigratePgConstraint) error
	Update_Confkey(*PostgresMigratePgConstraint) error
	Update_Confmatchtype(*PostgresMigratePgConstraint) error
	Update_Confrelid(*PostgresMigratePgConstraint) error
	Update_Confupdtype(*PostgresMigratePgConstraint) error
	Update_Conindid(*PostgresMigratePgConstraint) error
	Update_Coninhcount(*PostgresMigratePgConstraint) error
	Update_Conislocal(*PostgresMigratePgConstraint) error
	Update_Conkey(*PostgresMigratePgConstraint) error
	Update_Conname(*PostgresMigratePgConstraint) error
	Update_Connamespace(*PostgresMigratePgConstraint) error
	Update_Connoinherit(*PostgresMigratePgConstraint) error
	Update_Conparentid(*PostgresMigratePgConstraint) error
	Update_Conpfeqop(*PostgresMigratePgConstraint) error
	Update_Conppeqop(*PostgresMigratePgConstraint) error
	Update_Conrelid(*PostgresMigratePgConstraint) error
	Update_Contype(*PostgresMigratePgConstraint) error
	Update_Contypid(*PostgresMigratePgConstraint) error
	Update_Convalidated(*PostgresMigratePgConstraint) error
	Update_Oid(*PostgresMigratePgConstraint) error
	Update_VersionID(*PostgresMigratePgConstraint) error
}

// TableNameDB - возвращает имя таблицы в БД
func (m PostgresMigratePgConstraint) TableNameDB() string {
	return "postgres_migrate_pg_constraint"
}

// NewPostgresMigratePgConstraint - возвращает новый	объект
func NewPostgresMigratePgConstraint() PostgresMigratePgConstraint {
	return PostgresMigratePgConstraint{}
}

// AsPostgresMigratePgConstraint - создаёт объект из упакованного объекта в массиве байтов
func AsPostgresMigratePgConstraint(b []byte) (PostgresMigratePgConstraint, error) {
	c := NewPostgresMigratePgConstraint()
	err := msgpack.Unmarshal(b, &c)
	if err != nil {
		return NewPostgresMigratePgConstraint(), err
	}
	return c, nil
}

// PostgresMigratePgConstraintAsBytes - упаковывает объект в массив байтов
func PostgresMigratePgConstraintAsBytes(m *PostgresMigratePgConstraint) ([]byte, error) {
	b, err := msgpack.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetStructVersion - возвращает версию модели
func (m PostgresMigratePgConstraint) GetStructVersion() uint32 {
	if versionPostgresMigratePgConstraint == 0 {
		versionPostgresMigratePgConstraint = calc_struct_version.CalcStructVersion(reflect.TypeOf(m))
	}

	return versionPostgresMigratePgConstraint
}

// GetModelFromJSON - создаёт модель из строки json
func (m *PostgresMigratePgConstraint) GetModelFromJSON(sModel string) error {
	var err error

	var bytes []byte
	bytes = []byte(sModel)

	err = json.Unmarshal(bytes, m)

	return err
}

// GetJSON - возвращает строку json из модели
func (m PostgresMigratePgConstraint) GetJSON() (string, error) {
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
func (m *PostgresMigratePgConstraint) Read() error {
	if Crud_PostgresMigratePgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgConstraint.Read(m)

	return err
}

// Save - записывает объект в БД по ID
func (m *PostgresMigratePgConstraint) Save() error {
	if Crud_PostgresMigratePgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgConstraint.Save(m)

	return err
}

// Update - обновляет объект в БД по ID
func (m *PostgresMigratePgConstraint) Update() error {
	if Crud_PostgresMigratePgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgConstraint.Update(m)

	return err
}

// Create - создаёт объект в БД с новым ID
func (m *PostgresMigratePgConstraint) Create() error {
	if Crud_PostgresMigratePgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgConstraint.Create(m)

	return err
}

// ReadFromCache - находит запись в кэше или в БД по ID, и заполняет в объект
func (m *PostgresMigratePgConstraint) ReadFromCache(Oid int64, VersionID int64) (PostgresMigratePgConstraint, error) {
	Otvet := PostgresMigratePgConstraint{}
	var err error

	if Crud_PostgresMigratePgConstraint == nil {
		return Otvet, db_constants.ErrorCrudIsNotInit
	}

	Otvet, err = Crud_PostgresMigratePgConstraint.ReadFromCache(Oid, VersionID)

	return Otvet, err
}

// SetCrudInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func (m PostgresMigratePgConstraint) SetCrudInterface(crud ICrud_PostgresMigratePgConstraint) {
	Crud_PostgresMigratePgConstraint = crud

	return
}

// UpdateManyFields - находит запись в БД по ID, и изменяет только нужные колонки
func (m *PostgresMigratePgConstraint) UpdateManyFields(MassNeedUpdateFields []string) error {
	if Crud_PostgresMigratePgConstraint == nil {
		return db_constants.ErrorCrudIsNotInit
	}

	err := Crud_PostgresMigratePgConstraint.UpdateManyFields(m, MassNeedUpdateFields)

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
