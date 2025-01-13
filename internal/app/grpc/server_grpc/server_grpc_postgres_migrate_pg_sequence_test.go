//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

// PostgresMigratePgSequence_ID_Test - ID таблицы для тестирования
const PostgresMigratePgSequence_SEQRELID_Test = 0
const PostgresMigratePgSequence_VERSIONID_Test = 0

func Test_server_PostgresMigratePgSequence_Read(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Otvet, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Read() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Read() error: ModelString=''")
	}
}

func Test_server_PostgresMigratePgSequence_Delete(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()

	server1 := &ServerGRPC{}

	//прочитаем
	Response, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ", err)
	}
	if Response.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ModelString=''")
	}

	Otvet := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), Otvet)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Delete() Unmarshal() error: ", err)
	}

	if Otvet.IsDeleted == false {
		//пометим на удаление
		_, err = server1.PostgresMigratePgSequence_Delete(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ", err)
		}
		if Otvet.Seqrelid == 0 {
			t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ID =0")
		}

		//снимем пометку на удаление
		_, err = server1.PostgresMigratePgSequence_Restore(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ", err)
		}
		if Otvet.Seqrelid == 0 {
			t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ID =0")
		}
	} else {
		//снимем пометку на удаление
		_, err = server1.PostgresMigratePgSequence_Restore(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ", err)
		}
		if Otvet.Seqrelid == 0 {
			t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ID =0")
		}

		//пометим на удаление
		_, err = server1.PostgresMigratePgSequence_Delete(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ", err)
		}
		if Otvet.Seqrelid == 0 {
			t.Error("Test_server_PostgresMigratePgSequence_Delete() error: ID =0")
		}
	}

}

func Test_server_PostgresMigratePgSequence_Create(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var ModelString string
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = -1
	m.VersionID = -1
	ModelString, err := m.GetJSON()
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Create() error: ", err)
		return
	}

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	ctx := context.Background()
	server1 := &ServerGRPC{}
	Otvet, err := server1.PostgresMigratePgSequence_Create(ctx, &RequestModel)
	if err == nil {
		t.Error("Test_server_PostgresMigratePgSequence_Create() error: ", err)
	}
	if Otvet.ModelString != "" {
		t.Error("Test_server_PostgresMigratePgSequence_Create() error: ModelString !=''")
	}
}

func Test_server_PostgresMigratePgSequence_Update(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update() error: ", err)
		return
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update() error: ModelString=''")
	}

	var ModelString string
	ModelString = Response1.ModelString

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PostgresMigratePgSequence_Update(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Update() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Update() error: ModelString=''")
	}

}

func Test_server_PostgresMigratePgSequence_Save(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgSequence_SEQRELID_Test
	Request.Int64_2 = PostgresMigratePgSequence_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgSequence_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Save() error: ", err)
		return
	}

	var ModelString string
	ModelString = Response1.ModelString

	//sModel, _ := GetJSON(Otvet)
	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PostgresMigratePgSequence_Save(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgSequence_Save() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgSequence_Save() error: ModelString=''")
	}

}
