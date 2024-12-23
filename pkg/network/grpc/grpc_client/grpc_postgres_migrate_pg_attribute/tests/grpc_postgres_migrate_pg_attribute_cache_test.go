//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestReadFromCache(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_attribute.Crud_GRPC{}
	Otvet := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	Otvet, err := crud.ReadFromCache(ATTNAME_Test, ATTRELID_Test, VERSIONID_Test)

	if err != nil {
		t.Error("ReadFromCache() error: ", err)
	}

	if (Otvet.Attname == "") || (Otvet.Attrelid == 0) || (Otvet.VersionID == 0) {
		t.Error("ReadFromCache() error: ID =0")
	}
}
