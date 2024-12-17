//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_attribute"
	"context"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

// PgAttribute_ID_Test - ID таблицы для тестирования
const PgAttribute_ATTNAME_Test = ""
const PgAttribute_ATTRELID_Test = 0
const PgAttribute_VERSIONID_Test = 0

func Test_server_PgAttribute_Read(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PgAttribute_ATTNAME_Test
	Request.Int64_1 = PgAttribute_ATTRELID_Test
	Request.Int64_2 = PgAttribute_VERSIONID_Test
	Request.VersionModel = pg_attribute.PgAttribute{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Otvet, err := server1.PgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgAttribute_Read() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PgAttribute_Read() error: ModelString=''")
	}
}

func Test_server_PgAttribute_Create(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var ModelString string
	m := pg_attribute.PgAttribute{}
	m.Attname = ""
	m.Attrelid = -1
	m.VersionID = -1
	ModelString, err := m.GetJSON()
	if err != nil {
		t.Error("Test_server_PgAttribute_Create() error: ", err)
		return
	}

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = pg_attribute.PgAttribute{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	ctx := context.Background()
	server1 := &ServerGRPC{}
	Otvet, err := server1.PgAttribute_Create(ctx, &RequestModel)
	if err == nil {
		t.Error("Test_server_PgAttribute_Create() error: ", err)
	}
	if Otvet.ModelString != "" {
		t.Error("Test_server_PgAttribute_Create() error: ModelString !=''")
	}
}

func Test_server_PgAttribute_Update(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PgAttribute_ATTNAME_Test
	Request.Int64_1 = PgAttribute_ATTRELID_Test
	Request.Int64_2 = PgAttribute_VERSIONID_Test
	Request.VersionModel = pg_attribute.PgAttribute{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgAttribute_Update() error: ", err)
		return
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PgAttribute_Update() error: ModelString=''")
	}

	var ModelString string
	ModelString = Response1.ModelString

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = pg_attribute.PgAttribute{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PgAttribute_Update(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PgAttribute_Update() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PgAttribute_Update() error: ModelString=''")
	}

}

func Test_server_PgAttribute_Save(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PgAttribute_ATTNAME_Test
	Request.Int64_1 = PgAttribute_ATTRELID_Test
	Request.Int64_2 = PgAttribute_VERSIONID_Test
	Request.VersionModel = pg_attribute.PgAttribute{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgAttribute_Save() error: ", err)
		return
	}

	var ModelString string
	ModelString = Response1.ModelString

	//sModel, _ := GetJSON(Otvet)
	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = pg_attribute.PgAttribute{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PgAttribute_Save(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PgAttribute_Save() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PgAttribute_Save() error: ModelString=''")
	}

}
