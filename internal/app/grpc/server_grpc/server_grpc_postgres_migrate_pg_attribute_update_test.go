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

func TestServerGRPC_PostgresMigratePgAttribute_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var err error

	//прочитаем из БД
	ctx := context.Background()
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgAttribute_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД с пустым списком полей (не запишется)
	var ModelString string
	ModelString = Response1.ModelString
	RequestModel := grpc_proto.Request_Model_MassString{}
	RequestModel.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	RequestModel.ModelString = ModelString
	RequestModel.MassNames = nil

	_, err = server1.PostgresMigratePgAttribute_UpdateManyFields(ctx, &RequestModel)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgAttribute_UpdateManyFields() error: ", err)
	}

}

func Test_server_PostgresMigratePgAttribute_Update_Attalign(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attalign() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attalign() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attalign() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_String{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.String_2 = m.Attalign
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attalign(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attalign() Update_Attalign() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attbyval(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attbyval() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attbyval() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attbyval() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Bool{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Bool_1 = m.Attbyval
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attbyval(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attbyval() Update_Attbyval() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attcacheoff(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attcacheoff() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attcacheoff() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attcacheoff() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Int32{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int32_1 = m.Attcacheoff
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attcacheoff(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attcacheoff() Update_Attcacheoff() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attcollation(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attcollation() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attcollation() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attcollation() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Int64{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int64_3 = m.Attcollation
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attcollation(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attcollation() Update_Attcollation() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attgenerated(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attgenerated() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attgenerated() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attgenerated() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_String{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.String_2 = m.Attgenerated
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attgenerated(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attgenerated() Update_Attgenerated() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Atthasdef(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atthasdef() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atthasdef() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atthasdef() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Bool{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Bool_1 = m.Atthasdef
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Atthasdef(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atthasdef() Update_Atthasdef() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Atthasmissing(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atthasmissing() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atthasmissing() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atthasmissing() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Bool{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Bool_1 = m.Atthasmissing
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Atthasmissing(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atthasmissing() Update_Atthasmissing() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attidentity(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attidentity() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attidentity() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attidentity() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_String{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.String_2 = m.Attidentity
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attidentity(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attidentity() Update_Attidentity() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attinhcount(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attinhcount() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attinhcount() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attinhcount() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Int32{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int32_1 = m.Attinhcount
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attinhcount(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attinhcount() Update_Attinhcount() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attisdropped(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attisdropped() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attisdropped() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attisdropped() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Bool{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Bool_1 = m.Attisdropped
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attisdropped(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attisdropped() Update_Attisdropped() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attislocal(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attislocal() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attislocal() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attislocal() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Bool{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Bool_1 = m.Attislocal
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attislocal(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attislocal() Update_Attislocal() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attlen(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attlen() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attlen() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attlen() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Int32{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int32_1 = m.Attlen
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attlen(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attlen() Update_Attlen() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attname(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attname() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attname() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attname() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.String_1 = m.Attname
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attname(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attname() Update_Attname() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attndims(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attndims() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attndims() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attndims() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Int32{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int32_1 = m.Attndims
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attndims(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attndims() Update_Attndims() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attnotnull(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attnotnull() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attnotnull() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attnotnull() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Bool{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Bool_1 = m.Attnotnull
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attnotnull(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attnotnull() Update_Attnotnull() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attnum(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attnum() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attnum() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attnum() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Int32{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int32_1 = m.Attnum
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attnum(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attnum() Update_Attnum() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attrelid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attrelid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attrelid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attrelid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int64_1 = m.Attrelid
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attrelid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attrelid() Update_Attrelid() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attstattarget(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attstattarget() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attstattarget() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attstattarget() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Int32{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int32_1 = m.Attstattarget
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attstattarget(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attstattarget() Update_Attstattarget() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Attstorage(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attstorage() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attstorage() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attstorage() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_String{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.String_2 = m.Attstorage
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Attstorage(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Attstorage() Update_Attstorage() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Atttypid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atttypid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atttypid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atttypid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Int64{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int64_3 = m.Atttypid
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Atttypid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atttypid() Update_Atttypid() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_Atttypmod(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atttypmod() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atttypmod() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atttypmod() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64_Int32{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int32_1 = m.Atttypmod
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_Atttypmod(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_Atttypmod() Update_Atttypmod() error: ", err)
	}
}

func Test_server_PostgresMigratePgAttribute_Update_VersionID(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgAttribute_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_VersionID() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_VersionID() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_VersionID() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_String_Int64_Int64{}
	Request2.String_1 = PostgresMigratePgAttribute_ATTNAME_Test
	Request2.Int64_1 = PostgresMigratePgAttribute_ATTRELID_Test
	Request2.Int64_2 = PostgresMigratePgAttribute_VERSIONID_Test
	Request2.Int64_2 = m.VersionID
	Request2.VersionModel = postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	_, err = server1.PostgresMigratePgAttribute_Update_VersionID(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgAttribute_Update_VersionID() Update_VersionID() error: ", err)
	}
}
