//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

// PostgresMigratePgAttribute_ID_Test - ID таблицы для тестирования
const PostgresMigratePgAttribute_ATTNAME_Test = ""
const PostgresMigratePgAttribute_ATTRELID_Test = 0
const PostgresMigratePgAttribute_VERSIONID_Test = 0

func Test_server_PostgresMigratePgAttribute_Read(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Otvet, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Read() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Read() error: ModelString=''")
	}
}

func Test_server_PostgresMigratePgAttribute_Create(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var ModelString string
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = ""
	m.Attrelid = -1
	m.VersionID = -1
	ModelString, err := m.GetJSON()
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Create() error: ", err)
		return
	}

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	ctx := context.Background()
	server1 := &ServerGRPC{}
	Otvet, err := server1.PostgresMigratePgAttribute_Create(ctx, &RequestModel)
	if err == nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Create() error: ", err)
	}
	if Otvet.ModelString != "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Create() error: ModelString !=''")
	}
}

func Test_server_PostgresMigratePgAttribute_Update(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update() error: ", err)
		return
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update() error: ModelString=''")
	}

	var ModelString string
	ModelString = Response1.ModelString

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PostgresMigratePgAttribute_Update(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update() error: ModelString=''")
	}

}

func Test_server_PostgresMigratePgAttribute_Save(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Save() error: ", err)
		return
	}

	var ModelString string
	ModelString = Response1.ModelString

	//sModel, _ := GetJSON(Otvet)
	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PostgresMigratePgAttribute_Save(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Save() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Save() error: ModelString=''")
	}

}
