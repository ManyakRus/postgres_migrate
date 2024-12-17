//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_description"
	"context"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func Test_server_PgDescription_ReadFromCache(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = PgDescription_CLASSOID_Test
	Request.Int64_2 = PgDescription_OBJOID_Test
	Request.Int32_1 = PgDescription_OBJSUBID_Test
	Request.Int64_3 = PgDescription_VERSIONID_Test
	Request.VersionModel = pg_description.PgDescription{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Otvet, err := server1.PgDescription_ReadFromCache(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgDescription_ReadFromCache() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PgDescription_ReadFromCache() error: ModelString=''")
	}
}
