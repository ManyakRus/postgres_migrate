//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

const INDEXRELID_Test = 678300828
const VERSIONID_Test = 1

func TestRead(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestRead() error: ", err)
	}

	if Otvet.Indexrelid != INDEXRELID_Test {
		t.Error(TableName + "_test.TestRead() error ID != ", INDEXRELID_Test)
	} else {
		t.Log(TableName+"_test.TestRead() Otvet: ", Otvet.Indexrelid)
	}
}

func TestSave(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
	}

	if Otvet.Indexrelid != INDEXRELID_Test {
		t.Error(TableName + "_test.TestSave() error ID != ", INDEXRELID_Test)
	}

	err = crud.Save(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
	}
	t.Log(TableName+"_test.TestSave() Otvet: ", Otvet.Indexrelid)

}

func TestDelete(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestDelete() error: ", err)
	}

	if Otvet.IsDeleted == false {
		err = crud.Delete(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}

		err = crud.Restore(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}
	} else {
		err = crud.Restore(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}

		err = crud.Delete(&Otvet)
		if err != nil {
			t.Error("TestDelete() error: ", err)
		}

	}

}
