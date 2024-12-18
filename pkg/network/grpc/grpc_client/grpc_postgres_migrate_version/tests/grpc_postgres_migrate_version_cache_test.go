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

func TestReadFromCache(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_version.Crud_GRPC{}
	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet, err := crud.ReadFromCache(ID_Test)

	if err != nil {
		t.Error("ReadFromCache() error: ", err)
	}

	if Otvet.ID == 0 {
		t.Error("ReadFromCache() error: ID =0")
	}
}
