//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestServerGRPC_PostgresMigratePgSequence_UpdateManyFields(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var err error

	//прочитаем из БД
	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgSequence_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД с пустым списком полей (не запишется)
	var ModelString string
	ModelString = Response1.ModelString
	RequestModel := grpc_proto.Request_Model_MassString{}
	RequestModel.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	RequestModel.ModelString = ModelString
	RequestModel.MassNames = nil

	_, err = server1.PostgresMigratePgSequence_UpdateManyFields(ctx, &RequestModel)
	if err != nil {
		t.Error("TestServerGRPC_PostgresMigratePgSequence_UpdateManyFields() error: ", err)
	}

}

func Test_server_PostgresMigratePgSequence_Update_Seqcache(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqcache() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqcache() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqcache() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request2.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request2.Int64_3 = m.Seqcache
	Request2.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	_, err = server1.PostgresMigratePgSequence_Update_Seqcache(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqcache() Update_Seqcache() error: ", err)
	}
}

func Test_server_PostgresMigratePgSequence_Update_Seqcycle(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqcycle() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqcycle() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqcycle() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Bool{}
	Request2.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request2.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request2.Bool_1 = m.Seqcycle
	Request2.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	_, err = server1.PostgresMigratePgSequence_Update_Seqcycle(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqcycle() Update_Seqcycle() error: ", err)
	}
}

func Test_server_PostgresMigratePgSequence_Update_Seqincrement(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqincrement() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqincrement() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqincrement() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request2.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request2.Int64_3 = m.Seqincrement
	Request2.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	_, err = server1.PostgresMigratePgSequence_Update_Seqincrement(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqincrement() Update_Seqincrement() error: ", err)
	}
}

func Test_server_PostgresMigratePgSequence_Update_Seqmax(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqmax() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqmax() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqmax() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request2.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request2.Int64_3 = m.Seqmax
	Request2.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	_, err = server1.PostgresMigratePgSequence_Update_Seqmax(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqmax() Update_Seqmax() error: ", err)
	}
}

func Test_server_PostgresMigratePgSequence_Update_Seqmin(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqmin() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqmin() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqmin() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request2.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request2.Int64_3 = m.Seqmin
	Request2.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	_, err = server1.PostgresMigratePgSequence_Update_Seqmin(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqmin() Update_Seqmin() error: ", err)
	}
}

func Test_server_PostgresMigratePgSequence_Update_Seqrelid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqrelid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqrelid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqrelid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request2.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request2.Int64_1 = m.Seqrelid
	Request2.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	_, err = server1.PostgresMigratePgSequence_Update_Seqrelid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqrelid() Update_Seqrelid() error: ", err)
	}
}

func Test_server_PostgresMigratePgSequence_Update_Seqstart(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqstart() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqstart() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqstart() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request2.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request2.Int64_3 = m.Seqstart
	Request2.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	_, err = server1.PostgresMigratePgSequence_Update_Seqstart(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqstart() Update_Seqstart() error: ", err)
	}
}

func Test_server_PostgresMigratePgSequence_Update_Seqtypid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqtypid() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqtypid() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqtypid() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request2.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request2.Int64_3 = m.Seqtypid
	Request2.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	_, err = server1.PostgresMigratePgSequence_Update_Seqtypid(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_Seqtypid() Update_Seqtypid() error: ", err)
	}
}

func Test_server_PostgresMigratePgSequence_Update_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	server1 := &ServerGRPC{}
	ctx := context.Background()

	//прочитаем из БД
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_VersionID() Read() error: ", err)
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update_VersionID() Read() error: ModelString=''")
	}
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Response1.ModelString)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_VersionID() GetModelFromJSON() error: ", err)
	}

	//запишем в БД это же значение
	Request2 := grpc_proto.Request_Int64_Int64{}
	Request2.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request2.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request2.Int64_2 = m.VersionID
	Request2.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	_, err = server1.PostgresMigratePgSequence_Update_VersionID(ctx, &Request2)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update_VersionID() Update_VersionID() error: ", err)
	}
}
