//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

const CLASSOID_Test = 1259
const OBJOID_Test = 678300790
const OBJSUBID_Test = 1
const VERSIONID_Test = 1

func TestRead(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet.Classoid = CLASSOID_Test
	Otvet.Objoid = OBJOID_Test
	Otvet.Objsubid = OBJSUBID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestRead() error: ", err)
	}

	if Otvet.Classoid != CLASSOID_Test {
		t.Error(TableName+"_test.TestRead() error ID != ", CLASSOID_Test)
	} else {
		t.Log(TableName+"_test.TestRead() Otvet: ", Otvet.Classoid)
	}
}

func TestSave(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet.Classoid = CLASSOID_Test
	Otvet.Objoid = OBJOID_Test
	Otvet.Objsubid = OBJSUBID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
	}

	if Otvet.Classoid != CLASSOID_Test {
		t.Error(TableName+"_test.TestSave() error ID != ", CLASSOID_Test)
	}

	err = crud.Save(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
	}
	t.Log(TableName+"_test.TestSave() Otvet: ", Otvet.Classoid)

}

func TestDelete(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet.Classoid = CLASSOID_Test
	Otvet.Objoid = OBJOID_Test
	Otvet.Objsubid = OBJSUBID_Test
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
