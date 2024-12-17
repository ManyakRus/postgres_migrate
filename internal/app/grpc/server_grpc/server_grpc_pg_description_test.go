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

// PgDescription_ID_Test - ID таблицы для тестирования
const PgDescription_CLASSOID_Test = 0
const PgDescription_OBJOID_Test = 0
const PgDescription_OBJSUBID_Test = 0
const PgDescription_VERSIONID_Test = 0

func Test_server_PgDescription_Read(t *testing.T) {
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
	Otvet, err := server1.PgDescription_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgDescription_Read() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PgDescription_Read() error: ModelString=''")
	}
}

func Test_server_PgDescription_Create(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var ModelString string
	m := pg_description.PgDescription{}
	m.Classoid = -1
	m.Objoid = -1
	m.Objsubid = -1
	m.VersionID = -1
	ModelString, err := m.GetJSON()
	if err != nil {
		t.Error("Test_server_PgDescription_Create() error: ", err)
		return
	}

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = pg_description.PgDescription{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	ctx := context.Background()
	server1 := &ServerGRPC{}
	Otvet, err := server1.PgDescription_Create(ctx, &RequestModel)
	if err == nil {
		t.Error("Test_server_PgDescription_Create() error: ", err)
	}
	if Otvet.ModelString != "" {
		t.Error("Test_server_PgDescription_Create() error: ModelString !=''")
	}
}

func Test_server_PgDescription_Update(t *testing.T) {
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
	Response1, err := server1.PgDescription_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgDescription_Update() error: ", err)
		return
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PgDescription_Update() error: ModelString=''")
	}

	var ModelString string
	ModelString = Response1.ModelString

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = pg_description.PgDescription{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PgDescription_Update(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PgDescription_Update() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PgDescription_Update() error: ModelString=''")
	}

}

func Test_server_PgDescription_Save(t *testing.T) {
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
	Response1, err := server1.PgDescription_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgDescription_Save() error: ", err)
		return
	}

	var ModelString string
	ModelString = Response1.ModelString

	//sModel, _ := GetJSON(Otvet)
	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = pg_description.PgDescription{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PgDescription_Save(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PgDescription_Save() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PgDescription_Save() error: ModelString=''")
	}

}
