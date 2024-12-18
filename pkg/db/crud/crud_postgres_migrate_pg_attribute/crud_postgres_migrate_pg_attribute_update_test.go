//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
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
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
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

func TestUpdate_Attalign(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attalign() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attalign = m.Attalign
	err = crud.Update_Attalign(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attalign() Update() error: ", err)
	}
}

func TestUpdate_Attbyval(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attbyval() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attbyval = m.Attbyval
	err = crud.Update_Attbyval(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attbyval() Update() error: ", err)
	}
}

func TestUpdate_Attcacheoff(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attcacheoff() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attcacheoff = m.Attcacheoff
	err = crud.Update_Attcacheoff(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attcacheoff() Update() error: ", err)
	}
}

func TestUpdate_Attcollation(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attcollation() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attcollation = m.Attcollation
	err = crud.Update_Attcollation(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attcollation() Update() error: ", err)
	}
}

func TestUpdate_Attgenerated(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attgenerated() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attgenerated = m.Attgenerated
	err = crud.Update_Attgenerated(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attgenerated() Update() error: ", err)
	}
}

func TestUpdate_Atthasdef(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Atthasdef() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Atthasdef = m.Atthasdef
	err = crud.Update_Atthasdef(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Atthasdef() Update() error: ", err)
	}
}

func TestUpdate_Atthasmissing(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Atthasmissing() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Atthasmissing = m.Atthasmissing
	err = crud.Update_Atthasmissing(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Atthasmissing() Update() error: ", err)
	}
}

func TestUpdate_Attidentity(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attidentity() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attidentity = m.Attidentity
	err = crud.Update_Attidentity(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attidentity() Update() error: ", err)
	}
}

func TestUpdate_Attinhcount(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attinhcount() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attinhcount = m.Attinhcount
	err = crud.Update_Attinhcount(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attinhcount() Update() error: ", err)
	}
}

func TestUpdate_Attisdropped(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attisdropped() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attisdropped = m.Attisdropped
	err = crud.Update_Attisdropped(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attisdropped() Update() error: ", err)
	}
}

func TestUpdate_Attislocal(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attislocal() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attislocal = m.Attislocal
	err = crud.Update_Attislocal(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attislocal() Update() error: ", err)
	}
}

func TestUpdate_Attlen(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attlen() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attlen = m.Attlen
	err = crud.Update_Attlen(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attlen() Update() error: ", err)
	}
}

func TestUpdate_Attname(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attname() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attname = m.Attname
	err = crud.Update_Attname(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attname() Update() error: ", err)
	}
}

func TestUpdate_Attndims(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attndims() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attndims = m.Attndims
	err = crud.Update_Attndims(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attndims() Update() error: ", err)
	}
}

func TestUpdate_Attnotnull(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attnotnull() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attnotnull = m.Attnotnull
	err = crud.Update_Attnotnull(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attnotnull() Update() error: ", err)
	}
}

func TestUpdate_Attnum(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attnum() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attnum = m.Attnum
	err = crud.Update_Attnum(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attnum() Update() error: ", err)
	}
}

func TestUpdate_Attrelid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attrelid = m.Attrelid
	err = crud.Update_Attrelid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attrelid() Update() error: ", err)
	}
}

func TestUpdate_Attstattarget(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attstattarget() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attstattarget = m.Attstattarget
	err = crud.Update_Attstattarget(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attstattarget() Update() error: ", err)
	}
}

func TestUpdate_Attstorage(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Attstorage() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Attstorage = m.Attstorage
	err = crud.Update_Attstorage(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Attstorage() Update() error: ", err)
	}
}

func TestUpdate_Atttypid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Atttypid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Atttypid = m.Atttypid
	err = crud.Update_Atttypid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Atttypid() Update() error: ", err)
	}
}

func TestUpdate_Atttypmod(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Atttypmod() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Atttypmod = m.Atttypmod
	err = crud.Update_Atttypmod(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Atttypmod() Update() error: ", err)
	}
}

func TestUpdate_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestUpdate_VersionID() Update() error: ", err)
	}
}
