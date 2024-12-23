//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func Test_server_PostgresMigratePgDescription_ReadFromCache(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Otvet, err := server1.PostgresMigratePgDescription_ReadFromCache(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_ReadFromCache() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgDescription_ReadFromCache() error: ModelString=''")
	}
}
