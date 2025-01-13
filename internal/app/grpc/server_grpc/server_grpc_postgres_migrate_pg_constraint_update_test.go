//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_constraint"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestServerGRPC_PostgresMigratePgConstraint_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var err error

	//прочитаем из БД
	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgConstraint_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД с пустым списком полей (не запишется)
	var ModelString string
	ModelString = Response1.ModelString
	RequestModel := grpc_proto.Request_Model_MassString{}
	RequestModel.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	RequestModel.ModelString = ModelString
	RequestModel.MassNames = nil

	_, err = server1.PostgresMigratePgConstraint_UpdateManyFields(ctx, &RequestModel)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgConstraint_UpdateManyFields() error: ", err)
	}

}

func Test_server_PostgresMigratePgConstraint_Update_Condeferrable(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Condeferrable() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Condeferrable() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Condeferrable() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Bool_1 = m.Condeferrable
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Condeferrable(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Condeferrable() Update_Condeferrable() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Condeferred(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Condeferred() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Condeferred() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Condeferred() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Bool_1 = m.Condeferred
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Condeferred(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Condeferred() Update_Condeferred() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conexclop(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conexclop() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conexclop() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conexclop() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Conexclop
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conexclop(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conexclop() Update_Conexclop() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Confdeltype(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confdeltype() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confdeltype() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confdeltype() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Confdeltype
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Confdeltype(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confdeltype() Update_Confdeltype() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conffeqop(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conffeqop() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conffeqop() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conffeqop() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Conffeqop
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conffeqop(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conffeqop() Update_Conffeqop() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Confkey(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confkey() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confkey() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confkey() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Confkey
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Confkey(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confkey() Update_Confkey() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Confmatchtype(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confmatchtype() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confmatchtype() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confmatchtype() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Confmatchtype
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Confmatchtype(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confmatchtype() Update_Confmatchtype() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Confrelid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confrelid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confrelid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confrelid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Int64_3 = m.Confrelid
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Confrelid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confrelid() Update_Confrelid() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Confupdtype(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confupdtype() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confupdtype() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confupdtype() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Confupdtype
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Confupdtype(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Confupdtype() Update_Confupdtype() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conindid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conindid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conindid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conindid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Int64_3 = m.Conindid
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conindid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conindid() Update_Conindid() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Coninhcount(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Coninhcount() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Coninhcount() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Coninhcount() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int32{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Int32_1 = m.Coninhcount
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Coninhcount(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Coninhcount() Update_Coninhcount() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conislocal(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conislocal() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conislocal() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conislocal() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Bool_1 = m.Conislocal
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conislocal(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conislocal() Update_Conislocal() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conkey(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conkey() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conkey() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conkey() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Conkey
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conkey(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conkey() Update_Conkey() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conname(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conname() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conname() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conname() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Conname
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conname(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conname() Update_Conname() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Connamespace(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Connamespace() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Connamespace() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Connamespace() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Int64_3 = m.Connamespace
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Connamespace(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Connamespace() Update_Connamespace() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Connoinherit(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Connoinherit() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Connoinherit() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Connoinherit() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Bool_1 = m.Connoinherit
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Connoinherit(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Connoinherit() Update_Connoinherit() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conparentid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conparentid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conparentid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conparentid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Int64_3 = m.Conparentid
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conparentid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conparentid() Update_Conparentid() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conpfeqop(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conpfeqop() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conpfeqop() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conpfeqop() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Conpfeqop
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conpfeqop(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conpfeqop() Update_Conpfeqop() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conppeqop(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conppeqop() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conppeqop() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conppeqop() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Conppeqop
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conppeqop(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conppeqop() Update_Conppeqop() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Conrelid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conrelid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conrelid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conrelid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Int64_3 = m.Conrelid
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Conrelid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Conrelid() Update_Conrelid() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Contype(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Contype() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Contype() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Contype() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_String{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.String_1 = m.Contype
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Contype(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Contype() Update_Contype() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Contypid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Contypid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Contypid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Contypid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Int64_3 = m.Contypid
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Contypid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Contypid() Update_Contypid() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Convalidated(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Convalidated() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Convalidated() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Convalidated() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Bool_1 = m.Convalidated
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Convalidated(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Convalidated() Update_Convalidated() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_Oid(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Oid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Oid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Oid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Int64_1 = m.Oid
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_Oid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_Oid() Update_Oid() error: ", err)
	}
}

func Test_server_PostgresMigratePgConstraint_Update_VersionID(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_VersionID() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_VersionID() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_VersionID() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request2.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request2.Int64_2 = m.VersionID
	Request2.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	_, err = server1.PostgresMigratePgConstraint_Update_VersionID(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update_VersionID() Update_VersionID() error: ", err)
	}
}
