//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestReadFromCache(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet, err := crud.ReadFromCache(OID_Test, VERSIONID_Test)

	if err != nil {
		t.Error("ReadFromCache() error: ", err)
	}

	if (Otvet.Oid == 0) ||  (Otvet.VersionID == 0) {
		t.Error("ReadFromCache() error: ID =0")
	}
}
