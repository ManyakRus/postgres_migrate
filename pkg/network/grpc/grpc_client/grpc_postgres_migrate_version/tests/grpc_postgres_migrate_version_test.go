//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_version"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

// ID_Test - ID таблицы для тестирования
const ID_Test = 1

func TestGetVersionModel(t *testing.T) {

	crud := grpc_postgres_migrate_version.Crud_GRPC{}
	Otvet := crud.GetVersionModel()
	if Otvet == 0 {
		t.Error("TestGetVersionModel() error: Otvet =0")
	}
}

func TestRead(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_version.Crud_GRPC{}
	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet.ID = ID_Test
	err := crud.Read(&Otvet)

	if err != nil {
		t.Error("TestRead() error: ", err)
	}

	if Otvet.ID == 0 {
		t.Error("TestRead() error: ID =0")
	}
}

func TestCreate(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet.ID = -1

	crud := grpc_postgres_migrate_version.Crud_GRPC{}
	err := crud.Create(&Otvet)

	if err == nil {
		t.Error("TestCreate() error: ", err)
	}

}

func TestUpdate(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet.ID = 0

	crud := grpc_postgres_migrate_version.Crud_GRPC{}
	err := crud.Update(&Otvet)

	if err == nil {
		t.Error("TestUpdate() error: ", err)
	}

	if Otvet.ID != 0 {
		t.Error("TestUpdate() error: ID =0")
	}
}

func TestSave(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_version.Crud_GRPC{}
	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet.ID = ID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
		return
	}

	err = crud.Save(&Otvet)

	if err != nil {
		t.Error("TestSave() error: ", err)
	}

	if Otvet.ID == 0 {
		t.Error("TestSave() error: ID =0")
	}
}
