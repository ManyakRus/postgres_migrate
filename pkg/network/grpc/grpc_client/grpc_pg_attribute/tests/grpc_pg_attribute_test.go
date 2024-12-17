//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

// ATTNAME_Test - ID таблицы для тестирования
const ATTNAME_Test = ""
const ATTRELID_Test = 0
const VERSIONID_Test = 0

func TestGetVersionModel(t *testing.T) {
	t.SkipNow() //now rows in DB

	crud := grpc_pg_attribute.Crud_GRPC{}
	Otvet := crud.GetVersionModel()
	if Otvet == 0 {
		t.Error("TestGetVersionModel() error: Otvet =0")
	}
}

func TestRead(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)

	if err != nil {
		t.Error("TestRead() error: ", err)
	}

	if (Otvet.Attname == "") ||  (Otvet.Attrelid == 0) ||  (Otvet.VersionID == 0) {
		t.Error("TestRead() error: ID =0")
	}
}

func TestCreate(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = ""
	Otvet.Attrelid = -1
	Otvet.VersionID = -1

	crud := grpc_pg_attribute.Crud_GRPC{}
	err := crud.Create(&Otvet)

	if err == nil {
		t.Error("TestCreate() error: ", err)
	}

}

func TestUpdate(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = ""
	Otvet.Attrelid = 0
	Otvet.VersionID = 0

	crud := grpc_pg_attribute.Crud_GRPC{}
	err := crud.Update(&Otvet)

	if err == nil {
		t.Error("TestUpdate() error: ", err)
	}

	if Otvet.Attname != "" || Otvet.Attrelid != 0 || Otvet.VersionID != 0 {
		t.Error("TestUpdate() error: ID =0")
	}
}

func TestSave(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
		return
	}

	err = crud.Save(&Otvet)

	if err != nil {
		t.Error("TestSave() error: ", err)
	}

	if (Otvet.Attname == "") ||  (Otvet.Attrelid == 0) ||  (Otvet.VersionID == 0) {
		t.Error("TestSave() error: ID =0")
	}
}
