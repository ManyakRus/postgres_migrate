//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_namespace"
	"context"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestServerGRPC_PgNamespace_UpdateManyFields(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var err error

	//прочитаем из БД
	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PgNamespace_OID_Test
	Request.Int64_2 = PgNamespace_VERSIONID_Test
	Request.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PgNamespace_Read(ctx, &Request)
	if err != nil {
		t.Error("TestServerGRPC_PgNamespace_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД с пустым списком полей (не запишется)
	var ModelString string
	ModelString = Response1.ModelString
	RequestModel := grpc_proto.Request_Model_MassString{}
	RequestModel.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	RequestModel.ModelString = ModelString
	RequestModel.MassNames = nil

	_, err = server1.PgNamespace_UpdateManyFields(ctx, &RequestModel)
	if err != nil {
		t.Error("TestServerGRPC_PgNamespace_UpdateManyFields() error: ", err)
	}

}

func Test_server_PgNamespace_Update_Nspacl(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PgNamespace_OID_Test
	Request.Int64_2 = PgNamespace_VERSIONID_Test
	Request.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	Response1, err := server1.PgNamespace_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Nspacl() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PgNamespace_Update_Nspacl() Read() error: ModelString=''")
	}
	m := pg_namespace.PgNamespace{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Nspacl() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PgNamespace_OID_Test
	Request2.Int64_2 = PgNamespace_VERSIONID_Test
	Request2.String_1 = m.Nspacl
	Request2.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	_, err = server1.PgNamespace_Update_Nspacl(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Nspacl() Update_Nspacl() error: ", err)
	}
}

func Test_server_PgNamespace_Update_Nspname(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PgNamespace_OID_Test
	Request.Int64_2 = PgNamespace_VERSIONID_Test
	Request.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	Response1, err := server1.PgNamespace_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Nspname() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PgNamespace_Update_Nspname() Read() error: ModelString=''")
	}
	m := pg_namespace.PgNamespace{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Nspname() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PgNamespace_OID_Test
	Request2.Int64_2 = PgNamespace_VERSIONID_Test
	Request2.String_1 = m.Nspname
	Request2.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	_, err = server1.PgNamespace_Update_Nspname(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Nspname() Update_Nspname() error: ", err)
	}
}

func Test_server_PgNamespace_Update_Nspowner(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PgNamespace_OID_Test
	Request.Int64_2 = PgNamespace_VERSIONID_Test
	Request.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	Response1, err := server1.PgNamespace_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Nspowner() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PgNamespace_Update_Nspowner() Read() error: ModelString=''")
	}
	m := pg_namespace.PgNamespace{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Nspowner() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PgNamespace_OID_Test
	Request2.Int64_2 = PgNamespace_VERSIONID_Test
	Request2.Int64_3 = m.Nspowner
	Request2.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	_, err = server1.PgNamespace_Update_Nspowner(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Nspowner() Update_Nspowner() error: ", err)
	}
}

func Test_server_PgNamespace_Update_Oid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PgNamespace_OID_Test
	Request.Int64_2 = PgNamespace_VERSIONID_Test
	Request.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	Response1, err := server1.PgNamespace_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Oid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PgNamespace_Update_Oid() Read() error: ModelString=''")
	}
	m := pg_namespace.PgNamespace{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Oid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PgNamespace_OID_Test
	Request2.Int64_2 = PgNamespace_VERSIONID_Test
	Request2.Int64_1 = m.Oid
	Request2.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	_, err = server1.PgNamespace_Update_Oid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_Oid() Update_Oid() error: ", err)
	}
}

func Test_server_PgNamespace_Update_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PgNamespace_OID_Test
	Request.Int64_2 = PgNamespace_VERSIONID_Test
	Request.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	Response1, err := server1.PgNamespace_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_VersionID() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PgNamespace_Update_VersionID() Read() error: ModelString=''")
	}
	m := pg_namespace.PgNamespace{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_VersionID() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PgNamespace_OID_Test
	Request2.Int64_2 = PgNamespace_VERSIONID_Test
	Request2.Int64_2 = m.VersionID
	Request2.VersionModel = pg_namespace.PgNamespace{}.GetStructVersion()
	_, err = server1.PgNamespace_Update_VersionID(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PgNamespace_Update_VersionID() Update_VersionID() error: ", err)
	}
}
