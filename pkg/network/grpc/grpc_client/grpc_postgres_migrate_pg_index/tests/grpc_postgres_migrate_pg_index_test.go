//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

// INDEXRELID_Test - ID таблицы для тестирования
const INDEXRELID_Test = 678300828
const VERSIONID_Test = 1

func TestGetVersionModel(t *testing.T) {

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}
	Otvet := crud.GetVersionModel()
	if Otvet == 0 {
		t.Error("TestGetVersionModel() error: Otvet =0")
	}
}

func TestRead(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)

	if err != nil {
		t.Error("TestRead() error: ", err)
	}

	if (Otvet.Indexrelid == 0) || (Otvet.VersionID == 0) {
		t.Error("TestRead() error: ID =0")
	}
}

func TestCreate(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = -1
	Otvet.VersionID = -1

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}
	err := crud.Create(&Otvet)

	if err == nil {
		t.Error("TestCreate() error: ", err)
	}

}

func TestUpdate(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = 0
	Otvet.VersionID = 0

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}
	err := crud.Update(&Otvet)

	if err == nil {
		t.Error("TestUpdate() error: ", err)
	}

	if Otvet.Indexrelid != 0 || Otvet.VersionID != 0 {
		t.Error("TestUpdate() error: ID =0")
	}
}

func TestSave(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
		return
	}

	err = crud.Save(&Otvet)

	if err != nil {
		t.Error("TestSave() error: ", err)
	}

	if (Otvet.Indexrelid == 0) || (Otvet.VersionID == 0) {
		t.Error("TestSave() error: ID =0")
	}
}

func TestDelete(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestRead() error: ", err)
	}

	if Otvet.IsDeleted == false {
		err = crud.Delete(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}
		if (Otvet.Indexrelid == 0) || (Otvet.VersionID == 0) {
			t.Error("TestDelete() error: ID =0")
		}

		err = crud.Restore(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}
		if (Otvet.Indexrelid == 0) || (Otvet.VersionID == 0) {
			t.Error("TestDelete() error: ID =0")
		}
	} else {
		err = crud.Restore(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}
		if (Otvet.Indexrelid == 0) || (Otvet.VersionID == 0) {
			t.Error("TestDelete() error: ID =0")
		}

		err = crud.Delete(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}
		if (Otvet.Indexrelid == 0) || (Otvet.VersionID == 0) {
			t.Error("TestDelete() error: ID =0")
		}
	}
}
