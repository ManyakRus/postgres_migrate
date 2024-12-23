//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_func"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_version"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_version"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestReadObject(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_version.Crud_GRPC{}
	Otvet := object_postgres_migrate_version.ObjectPostgresMigrateVersion{}
	Otvet.ID = ID_Test
	err := crud.ReadObject(&Otvet)

	if err != nil && crud_func.IsRecordNotFound(err) == false {
		t.Error("TestReadObject() error: ", err)
	}
}
