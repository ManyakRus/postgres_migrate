//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_namespace"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

const OID_Test = 0
const VERSIONID_Test = 0

func TestRead(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := pg_namespace.PgNamespace{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestRead() error: ", err)
	}

	if Otvet.Oid != OID_Test {
		t.Error(TableName + "_test.TestRead() error ID != ", OID_Test)
	} else {
		t.Log(TableName+"_test.TestRead() Otvet: ", Otvet.Oid)
	}
}

func TestSave(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	crud := Crud_DB{}
	Otvet := pg_namespace.PgNamespace{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
	}

	if Otvet.Oid != OID_Test {
		t.Error(TableName + "_test.TestSave() error ID != ", OID_Test)
	}

	err = crud.Save(&Otvet)
	if err != nil {
		t.Error("TestSave() error: ", err)
	}
	t.Log(TableName+"_test.TestSave() Otvet: ", Otvet.Oid)

}
