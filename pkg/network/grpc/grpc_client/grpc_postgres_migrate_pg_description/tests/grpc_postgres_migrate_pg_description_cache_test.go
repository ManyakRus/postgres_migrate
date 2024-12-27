//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestReadFromCache(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_description.Crud_GRPC{}
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet, err := crud.ReadFromCache(CLASSOID_Test, OBJOID_Test, OBJSUBID_Test, VERSIONID_Test)

	if err != nil {
		t.Error("ReadFromCache() error: ", err)
	}

	if (Otvet.Classoid == 0) ||  (Otvet.Objoid == 0) ||  (Otvet.Objsubid == 0) ||  (Otvet.VersionID == 0) {
		t.Error("ReadFromCache() error: ID =0")
	}
}
