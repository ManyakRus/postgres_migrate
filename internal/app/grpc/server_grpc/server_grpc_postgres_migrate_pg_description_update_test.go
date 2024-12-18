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

func TestServerGRPC_PostgresMigratePgDescription_UpdateManyFields(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var err error

	//прочитаем из БД
	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgDescription_Read(ctx, &Request)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgDescription_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД с пустым списком полей (не запишется)
	var ModelString string
	ModelString = Response1.ModelString
	RequestModel := grpc_proto.Request_Model_MassString{}
	RequestModel.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	RequestModel.ModelString = ModelString
	RequestModel.MassNames = nil

	_, err = server1.PostgresMigratePgDescription_UpdateManyFields(ctx, &RequestModel)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgDescription_UpdateManyFields() error: ", err)
	}

}

func Test_server_PostgresMigratePgDescription_Update_Classoid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgDescription_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Classoid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Classoid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Classoid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request2.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request2.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request2.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request2.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request2.Int64_1 = m.Classoid
	Request2.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	_, err = server1.PostgresMigratePgDescription_Update_Classoid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Classoid() Update_Classoid() error: ", err)
	}
}

func Test_server_PostgresMigratePgDescription_Update_Description(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgDescription_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Description() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Description() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Description() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request2.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request2.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request2.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request2.String_1 = m.Description
	Request2.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	_, err = server1.PostgresMigratePgDescription_Update_Description(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Description() Update_Description() error: ", err)
	}
}

func Test_server_PostgresMigratePgDescription_Update_Objoid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgDescription_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Objoid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Objoid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Objoid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request2.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request2.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request2.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request2.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request2.Int64_2 = m.Objoid
	Request2.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	_, err = server1.PostgresMigratePgDescription_Update_Objoid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Objoid() Update_Objoid() error: ", err)
	}
}

func Test_server_PostgresMigratePgDescription_Update_Objsubid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgDescription_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Objsubid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Objsubid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Objsubid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request2.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request2.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request2.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request2.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request2.Int32_1 = m.Objsubid
	Request2.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	_, err = server1.PostgresMigratePgDescription_Update_Objsubid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_Objsubid() Update_Objsubid() error: ", err)
	}
}

func Test_server_PostgresMigratePgDescription_Update_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgDescription_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_VersionID() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgDescription_Update_VersionID() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_VersionID() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request2.Int64_1 = PostgresMigratePgDescription_CLASSOID_Test
	Request2.Int64_2 = PostgresMigratePgDescription_OBJOID_Test
	Request2.Int32_1 = PostgresMigratePgDescription_OBJSUBID_Test
	Request2.Int64_3 = PostgresMigratePgDescription_VERSIONID_Test
	Request2.Int64_3 = m.VersionID
	Request2.VersionModel = postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	_, err = server1.PostgresMigratePgDescription_Update_VersionID(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgDescription_Update_VersionID() Update_VersionID() error: ", err)
	}
}
