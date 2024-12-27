//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"context"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestServerGRPC_PostgresMigratePgIndex_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var err error

	//прочитаем из БД
	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgIndex_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД с пустым списком полей (не запишется)
	var ModelString string
	ModelString = Response1.ModelString
	RequestModel := grpc_proto.Request_Model_MassString{}
	RequestModel.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	RequestModel.ModelString = ModelString
	RequestModel.MassNames = nil

	_, err = server1.PostgresMigratePgIndex_UpdateManyFields(ctx, &RequestModel)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgIndex_UpdateManyFields() error: ", err)
	}

}

func Test_server_PostgresMigratePgIndex_Update_Indcheckxmin(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indcheckxmin() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indcheckxmin() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indcheckxmin() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indcheckxmin
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indcheckxmin(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indcheckxmin() Update_Indcheckxmin() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indclass(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indclass() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indclass() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indclass() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.String_1 = m.Indclass
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indclass(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indclass() Update_Indclass() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indcollation(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indcollation() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indcollation() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indcollation() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.String_1 = m.Indcollation
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indcollation(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indcollation() Update_Indcollation() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indexprs(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indexprs() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indexprs() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indexprs() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.String_1 = m.Indexprs
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indexprs(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indexprs() Update_Indexprs() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indexrelid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indexrelid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indexrelid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indexrelid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Int64_1 = m.Indexrelid
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indexrelid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indexrelid() Update_Indexrelid() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indimmediate(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indimmediate() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indimmediate() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indimmediate() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indimmediate
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indimmediate(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indimmediate() Update_Indimmediate() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indisclustered(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisclustered() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisclustered() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisclustered() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indisclustered
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indisclustered(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisclustered() Update_Indisclustered() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indisexclusion(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisexclusion() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisexclusion() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisexclusion() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indisexclusion
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indisexclusion(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisexclusion() Update_Indisexclusion() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indislive(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indislive() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indislive() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indislive() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indislive
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indislive(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indislive() Update_Indislive() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indisprimary(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisprimary() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisprimary() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisprimary() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indisprimary
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indisprimary(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisprimary() Update_Indisprimary() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indisready(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisready() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisready() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisready() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indisready
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indisready(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisready() Update_Indisready() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indisreplident(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisreplident() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisreplident() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisreplident() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indisreplident
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indisreplident(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisreplident() Update_Indisreplident() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indisunique(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisunique() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisunique() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisunique() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indisunique
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indisunique(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisunique() Update_Indisunique() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indisvalid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisvalid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisvalid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisvalid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Bool_1 = m.Indisvalid
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indisvalid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indisvalid() Update_Indisvalid() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indkey(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indkey() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indkey() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indkey() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.String_1 = m.Indkey
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indkey(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indkey() Update_Indkey() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indnatts(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indnatts() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indnatts() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indnatts() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Int32_1 = m.Indnatts
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indnatts(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indnatts() Update_Indnatts() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indnkeyatts(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indnkeyatts() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indnkeyatts() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indnkeyatts() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Int32_1 = m.Indnkeyatts
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indnkeyatts(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indnkeyatts() Update_Indnkeyatts() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indoption(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indoption() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indoption() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indoption() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.String_1 = m.Indoption
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indoption(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indoption() Update_Indoption() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indpred(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indpred() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indpred() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indpred() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.String_1 = m.Indpred
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indpred(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indpred() Update_Indpred() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_Indrelid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indrelid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indrelid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indrelid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Int64_3 = m.Indrelid
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_Indrelid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_Indrelid() Update_Indrelid() error: ", err)
	}
}

func Test_server_PostgresMigratePgIndex_Update_VersionID(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_VersionID() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update_VersionID() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_VersionID() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request2.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request2.Int64_2 = m.VersionID
	Request2.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	_, err = server1.PostgresMigratePgIndex_Update_VersionID(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update_VersionID() Update_VersionID() error: ", err)
	}
}
