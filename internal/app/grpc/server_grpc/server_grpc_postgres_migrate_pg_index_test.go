//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

// PostgresMigratePgIndex_ID_Test - ID таблицы для тестирования
const PostgresMigratePgIndex_INDEXRELID_Test = 678300828
const PostgresMigratePgIndex_VERSIONID_Test = 1

func Test_server_PostgresMigratePgIndex_Read(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Otvet, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Read() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Read() error: ModelString=''")
	}
}

func Test_server_PostgresMigratePgIndex_Delete(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()

	server1 := &ServerGRPC{}

	//прочитаем
	Response, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ", err)
	}
	if Response.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ModelString=''")
	}

	Otvet := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), Otvet)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Delete() Unmarshal() error: ", err)
	}

	if Otvet.IsDeleted == false {
		//пометим на удаление
		_, err = server1.PostgresMigratePgIndex_Delete(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ", err)
		}
		if Otvet.Indexrelid == 0 {
			t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ID =0")
		}

		//снимем пометку на удаление
		_, err = server1.PostgresMigratePgIndex_Restore(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ", err)
		}
		if Otvet.Indexrelid == 0 {
			t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ID =0")
		}
	} else {
		//снимем пометку на удаление
		_, err = server1.PostgresMigratePgIndex_Restore(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ", err)
		}
		if Otvet.Indexrelid == 0 {
			t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ID =0")
		}

		//пометим на удаление
		_, err = server1.PostgresMigratePgIndex_Delete(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ", err)
		}
		if Otvet.Indexrelid == 0 {
			t.Error("Test_server_PostgresMigratePgIndex_Delete() error: ID =0")
		}
	}

}

func Test_server_PostgresMigratePgIndex_Create(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var ModelString string
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = -1
	m.VersionID = -1
	ModelString, err := m.GetJSON()
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Create() error: ", err)
		return
	}

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	ctx := context.Background()
	server1 := &ServerGRPC{}
	Otvet, err := server1.PostgresMigratePgIndex_Create(ctx, &RequestModel)
	if err == nil {
		t.Error("Test_server_PostgresMigratePgIndex_Create() error: ", err)
	}
	if Otvet.ModelString != "" {
		t.Error("Test_server_PostgresMigratePgIndex_Create() error: ModelString !=''")
	}
}

func Test_server_PostgresMigratePgIndex_Update(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update() error: ", err)
		return
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update() error: ModelString=''")
	}

	var ModelString string
	ModelString = Response1.ModelString

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PostgresMigratePgIndex_Update(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Update() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Update() error: ModelString=''")
	}

}

func Test_server_PostgresMigratePgIndex_Save(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgIndex_INDEXRELID_Test
	Request.Int64_2 = PostgresMigratePgIndex_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgIndex_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Save() error: ", err)
		return
	}

	var ModelString string
	ModelString = Response1.ModelString

	//sModel, _ := GetJSON(Otvet)
	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PostgresMigratePgIndex_Save(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgIndex_Save() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgIndex_Save() error: ModelString=''")
	}

}
