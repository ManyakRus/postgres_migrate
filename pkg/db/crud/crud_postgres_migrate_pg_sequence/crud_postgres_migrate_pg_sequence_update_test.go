//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_sequence

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestUpdateManyFields(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdateManyFields() Read() error: ", err)
	}

	//запишем в БД это же значение
	err = crud.UpdateManyFields(&m, nil)
	if err != nil {
		t.Error("TestUpdateManyFields() UpdateManyFields() error: ", err)
	}
}

func TestUpdate_Seqcache(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Seqcache() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Seqcache = m.Seqcache
	err = crud.Update_Seqcache(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Seqcache() Update() error: ", err)
	}
}

func TestUpdate_Seqcycle(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Seqcycle() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Seqcycle = m.Seqcycle
	err = crud.Update_Seqcycle(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Seqcycle() Update() error: ", err)
	}
}

func TestUpdate_Seqincrement(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Seqincrement() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Seqincrement = m.Seqincrement
	err = crud.Update_Seqincrement(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Seqincrement() Update() error: ", err)
	}
}

func TestUpdate_Seqmax(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Seqmax() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Seqmax = m.Seqmax
	err = crud.Update_Seqmax(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Seqmax() Update() error: ", err)
	}
}

func TestUpdate_Seqmin(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Seqmin() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Seqmin = m.Seqmin
	err = crud.Update_Seqmin(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Seqmin() Update() error: ", err)
	}
}

func TestUpdate_Seqrelid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Seqrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Seqrelid = m.Seqrelid
	err = crud.Update_Seqrelid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Seqrelid() Update() error: ", err)
	}
}

func TestUpdate_Seqstart(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Seqstart() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Seqstart = m.Seqstart
	err = crud.Update_Seqstart(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Seqstart() Update() error: ", err)
	}
}

func TestUpdate_Seqtypid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Seqtypid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Seqtypid = m.Seqtypid
	err = crud.Update_Seqtypid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Seqtypid() Update() error: ", err)
	}
}

func TestUpdate_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestUpdate_VersionID() Update() error: ", err)
	}
}
