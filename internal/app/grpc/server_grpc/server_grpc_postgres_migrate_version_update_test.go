//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
	"context"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestServerGRPC_PostgresMigrateVersion_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var err error

	//прочитаем из БД
	ctx := context.Background()
	Request := grpc_proto.Request_Int64{}
	Request.Int64_1 = PostgresMigrateVersion_ID_Test
	Request.VersionModel = postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigrateVersion_Read(ctx, &Request)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigrateVersion_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД с пустым списком полей (не запишется)
	var ModelString string
	ModelString = Response1.ModelString
	RequestModel := grpc_proto.Request_Model_MassString{}
	RequestModel.VersionModel = postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()
	RequestModel.ModelString = ModelString
	RequestModel.MassNames = nil

	_, err = server1.PostgresMigrateVersion_UpdateManyFields(ctx, &RequestModel)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigrateVersion_UpdateManyFields() error: ", err)
	}

}

func Test_server_PostgresMigrateVersion_Update_Description(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64{}
	Request.Int64_1 = PostgresMigrateVersion_ID_Test
	Request.VersionModel = postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()
	Response1, err := server1.PostgresMigrateVersion_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigrateVersion_Update_Description() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigrateVersion_Update_Description() Read() error: ModelString=''")
	}
	m := postgres_migrate_version.PostgresMigrateVersion{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigrateVersion_Update_Description() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_String{}
	Request2.Int64_1 = PostgresMigrateVersion_ID_Test
	Request2.String_1 = m.Description
	Request2.VersionModel = postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()
	_, err = server1.PostgresMigrateVersion_Update_Description(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigrateVersion_Update_Description() Update_Description() error: ", err)
	}
}

func Test_server_PostgresMigrateVersion_Update_Name(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64{}
	Request.Int64_1 = PostgresMigrateVersion_ID_Test
	Request.VersionModel = postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()
	Response1, err := server1.PostgresMigrateVersion_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigrateVersion_Update_Name() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigrateVersion_Update_Name() Read() error: ModelString=''")
	}
	m := postgres_migrate_version.PostgresMigrateVersion{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigrateVersion_Update_Name() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_String{}
	Request2.Int64_1 = PostgresMigrateVersion_ID_Test
	Request2.String_1 = m.Name
	Request2.VersionModel = postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()
	_, err = server1.PostgresMigrateVersion_Update_Name(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigrateVersion_Update_Name() Update_Name() error: ", err)
	}
}
