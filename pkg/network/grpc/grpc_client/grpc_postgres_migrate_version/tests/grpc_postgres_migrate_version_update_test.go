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

func TestCrud_GRPC_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	//прочитаем из БД
	crud := grpc_postgres_migrate_version.Crud_GRPC{}
	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet.ID = ID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД, пустой список полей (не изменится)
	err = crud.UpdateManyFields(&Otvet, nil)

	if err != nil {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ", err)
	}

	if Otvet.ID == 0 {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ID =0")
	}
}

func TestCrud_GRPC_Update_Description(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_version.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_version.PostgresMigrateVersion{}
	m.ID = ID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Description() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet.ID = m.ID
	Otvet.Description = m.Description
	err = crud.Update_Description(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Description() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Name(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_version.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_version.PostgresMigrateVersion{}
	m.ID = ID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Name() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet.ID = m.ID
	Otvet.Name = m.Name
	err = crud.Update_Name(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Name() Update() error: ", err)
	}
}
