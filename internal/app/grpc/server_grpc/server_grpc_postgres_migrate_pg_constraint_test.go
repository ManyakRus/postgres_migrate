//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_constraint"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

// PostgresMigratePgConstraint_ID_Test - ID таблицы для тестирования
const PostgresMigratePgConstraint_OID_Test = 678300829
const PostgresMigratePgConstraint_VERSIONID_Test = 1

func Test_server_PostgresMigratePgConstraint_Read(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Otvet, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Read() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Read() error: ModelString=''")
	}
}

func Test_server_PostgresMigratePgConstraint_Delete(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()

	server1 := &ServerGRPC{}

	//прочитаем
	Response, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ", err)
	}
	if Response.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ModelString=''")
	}

	Otvet := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), Otvet)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Delete() Unmarshal() error: ", err)
	}

	if Otvet.IsDeleted == false {
		//пометим на удаление
		_, err = server1.PostgresMigratePgConstraint_Delete(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ", err)
		}
		if Otvet.Oid == 0 {
			t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ID =0")
		}

		//снимем пометку на удаление
		_, err = server1.PostgresMigratePgConstraint_Restore(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ", err)
		}
		if Otvet.Oid == 0 {
			t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ID =0")
		}
	} else {
		//снимем пометку на удаление
		_, err = server1.PostgresMigratePgConstraint_Restore(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ", err)
		}
		if Otvet.Oid == 0 {
			t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ID =0")
		}

		//пометим на удаление
		_, err = server1.PostgresMigratePgConstraint_Delete(ctx, &Request)
		if err != nil {
			t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ", err)
		}
		if Otvet.Oid == 0 {
			t.Error("Test_server_PostgresMigratePgConstraint_Delete() error: ID =0")
		}
	}

}

func Test_server_PostgresMigratePgConstraint_Create(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	var ModelString string
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = -1
	m.VersionID = -1
	ModelString, err := m.GetJSON()
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Create() error: ", err)
		return
	}

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	ctx := context.Background()
	server1 := &ServerGRPC{}
	Otvet, err := server1.PostgresMigratePgConstraint_Create(ctx, &RequestModel)
	if err == nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Create() error: ", err)
	}
	if Otvet.ModelString != "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Create() error: ModelString !=''")
	}
}

func Test_server_PostgresMigratePgConstraint_Update(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update() error: ", err)
		return
	}
	if Response1.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update() error: ModelString=''")
	}

	var ModelString string
	ModelString = Response1.ModelString

	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PostgresMigratePgConstraint_Update(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Update() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Update() error: ModelString=''")
	}

}

func Test_server_PostgresMigratePgConstraint_Save(t *testing.T) {
	config_main.LoadEnv()
	crud_starter.InitCrudTransport_DB()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	ctx := context.Background()
	Request := grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = PostgresMigratePgConstraint_OID_Test
	Request.Int64_2 = PostgresMigratePgConstraint_VERSIONID_Test
	Request.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()

	server1 := &ServerGRPC{}
	Response1, err := server1.PostgresMigratePgConstraint_Read(ctx, &Request)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Save() error: ", err)
		return
	}

	var ModelString string
	ModelString = Response1.ModelString

	//sModel, _ := GetJSON(Otvet)
	RequestModel := grpc_proto.RequestModel{}
	RequestModel.VersionModel = postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	RequestModel.ModelString = ModelString

	Otvet, err := server1.PostgresMigratePgConstraint_Save(ctx, &RequestModel)
	if err != nil {
		t.Error("Test_server_PostgresMigratePgConstraint_Save() error: ", err)
	}
	if Otvet.ModelString == "" {
		t.Error("Test_server_PostgresMigratePgConstraint_Save() error: ModelString=''")
	}

}
