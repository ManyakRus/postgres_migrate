//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestServerGRPC_PostgresMigratePgClass_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var err error

	//прочитаем из БД
	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgClass_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД с пустым списком полей (не запишется)
	var ModelString string
	ModelString = Response1.ModelString
	RequestModel := grpc_proto.Request_Model_MassString{}
	RequestModel.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	RequestModel.ModelString = ModelString
	RequestModel.MassNames = nil

	_, err = server1.PostgresMigratePgClass_UpdateManyFields(ctx, &RequestModel)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgClass_UpdateManyFields() error: ", err)
	}

}

func Test_server_PostgresMigratePgClass_Update_Oid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Oid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Oid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Oid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_1 = m.Oid
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Oid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Oid() Update_Oid() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relallvisible(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relallvisible() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relallvisible() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relallvisible() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int32_1 = m.Relallvisible
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relallvisible(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relallvisible() Update_Relallvisible() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relam(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relam() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relam() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relam() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_3 = m.Relam
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relam(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relam() Update_Relam() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relchecks(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relchecks() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relchecks() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relchecks() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int32_1 = m.Relchecks
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relchecks(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relchecks() Update_Relchecks() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relfilenode(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relfilenode() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relfilenode() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relfilenode() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_3 = m.Relfilenode
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relfilenode(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relfilenode() Update_Relfilenode() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relforcerowsecurity(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relforcerowsecurity() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relforcerowsecurity() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relforcerowsecurity() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Bool_1 = m.Relforcerowsecurity
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relforcerowsecurity(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relforcerowsecurity() Update_Relforcerowsecurity() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relfrozenxid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relfrozenxid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relfrozenxid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relfrozenxid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int32_1 = m.Relfrozenxid
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relfrozenxid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relfrozenxid() Update_Relfrozenxid() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relhasindex(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhasindex() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhasindex() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhasindex() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Bool_1 = m.Relhasindex
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relhasindex(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhasindex() Update_Relhasindex() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relhasrules(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhasrules() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhasrules() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhasrules() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Bool_1 = m.Relhasrules
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relhasrules(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhasrules() Update_Relhasrules() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relhassubclass(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhassubclass() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhassubclass() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhassubclass() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Bool_1 = m.Relhassubclass
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relhassubclass(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhassubclass() Update_Relhassubclass() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relhastriggers(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhastriggers() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhastriggers() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhastriggers() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Bool_1 = m.Relhastriggers
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relhastriggers(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relhastriggers() Update_Relhastriggers() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relispartition(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relispartition() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relispartition() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relispartition() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Bool_1 = m.Relispartition
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relispartition(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relispartition() Update_Relispartition() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relispopulated(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relispopulated() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relispopulated() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relispopulated() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Bool_1 = m.Relispopulated
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relispopulated(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relispopulated() Update_Relispopulated() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relisshared(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relisshared() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relisshared() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relisshared() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Bool_1 = m.Relisshared
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relisshared(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relisshared() Update_Relisshared() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relkind(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relkind() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relkind() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relkind() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.String_1 = m.Relkind
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relkind(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relkind() Update_Relkind() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relminmxid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relminmxid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relminmxid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relminmxid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int32_1 = m.Relminmxid
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relminmxid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relminmxid() Update_Relminmxid() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relname(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relname() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relname() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relname() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.String_1 = m.Relname
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relname(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relname() Update_Relname() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relnamespace(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relnamespace() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relnamespace() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relnamespace() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_3 = m.Relnamespace
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relnamespace(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relnamespace() Update_Relnamespace() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relnatts(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relnatts() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relnatts() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relnatts() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int32_1 = m.Relnatts
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relnatts(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relnatts() Update_Relnatts() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Reloftype(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reloftype() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reloftype() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reloftype() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_3 = m.Reloftype
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Reloftype(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reloftype() Update_Reloftype() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relowner(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relowner() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relowner() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relowner() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_3 = m.Relowner
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relowner(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relowner() Update_Relowner() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relpages(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relpages() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relpages() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relpages() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int32_1 = m.Relpages
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relpages(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relpages() Update_Relpages() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relpersistence(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relpersistence() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relpersistence() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relpersistence() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.String_1 = m.Relpersistence
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relpersistence(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relpersistence() Update_Relpersistence() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relreplident(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relreplident() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relreplident() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relreplident() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.String_1 = m.Relreplident
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relreplident(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relreplident() Update_Relreplident() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relrewrite(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relrewrite() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relrewrite() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relrewrite() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_3 = m.Relrewrite
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relrewrite(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relrewrite() Update_Relrewrite() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Relrowsecurity(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relrowsecurity() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relrowsecurity() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relrowsecurity() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Bool_1 = m.Relrowsecurity
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Relrowsecurity(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Relrowsecurity() Update_Relrowsecurity() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Reltablespace(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltablespace() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltablespace() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltablespace() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_3 = m.Reltablespace
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Reltablespace(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltablespace() Update_Reltablespace() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Reltoastrelid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltoastrelid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltoastrelid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltoastrelid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_3 = m.Reltoastrelid
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Reltoastrelid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltoastrelid() Update_Reltoastrelid() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Reltuples(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltuples() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltuples() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltuples() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Float32{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Float32_1 = m.Reltuples
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Reltuples(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltuples() Update_Reltuples() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_Reltype(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltype() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltype() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltype() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_3 = m.Reltype
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_Reltype(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_Reltype() Update_Reltype() error: ", err)
	}
}

func Test_server_PostgresMigratePgClass_Update_VersionID(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgClass_OID_Test
	Request.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgClass_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_VersionID() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgClass_Update_VersionID() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_VersionID() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgClass_OID_Test
	Request2.Int64_2 = PostgresMigratePgClass_VERSIONID_Test
	Request2.Int64_2 = m.VersionID
	Request2.VersionModel = postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	_, err = server1.PostgresMigratePgClass_Update_VersionID(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgClass_Update_VersionID() Update_VersionID() error: ", err)
	}
}
