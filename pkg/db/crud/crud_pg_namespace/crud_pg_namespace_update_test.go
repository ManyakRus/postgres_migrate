//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_namespace"
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
	m := pg_namespace.PgNamespace{}
	m.Oid = OID_Test
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

func TestUpdate_Nspacl(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_namespace.PgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Nspacl() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_namespace.PgNamespace{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Nspacl = m.Nspacl
	err = crud.Update_Nspacl(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Nspacl() Update() error: ", err)
	}
}

func TestUpdate_Nspname(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_namespace.PgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Nspname() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_namespace.PgNamespace{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Nspname = m.Nspname
	err = crud.Update_Nspname(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Nspname() Update() error: ", err)
	}
}

func TestUpdate_Nspowner(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_namespace.PgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Nspowner() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_namespace.PgNamespace{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Nspowner = m.Nspowner
	err = crud.Update_Nspowner(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Nspowner() Update() error: ", err)
	}
}

func TestUpdate_Oid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_namespace.PgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Oid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_namespace.PgNamespace{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Oid = m.Oid
	err = crud.Update_Oid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Oid() Update() error: ", err)
	}
}

func TestUpdate_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_namespace.PgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_namespace.PgNamespace{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestUpdate_VersionID() Update() error: ", err)
	}
}
