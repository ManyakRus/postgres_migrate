//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_namespace"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"reflect"
	"testing"
)

func TestReadFromCache(t *testing.T) {
	t.SkipNow() //now rows in DB

	var err error

	config_main.LoadEnv()
	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//читаем из БД
	m1 := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m1.Oid = OID_Test
	m1.VersionID = VERSIONID_Test
	err = Crud_DB{}.Read(&m1)
	if err != nil {
		t.Errorf("TestReadFromCache() error:t %v", err)
	}

	//читаем из Кеша
	m2, err := Crud_DB{}.ReadFromCache(OID_Test, VERSIONID_Test)
	if err != nil {
		t.Errorf("TestReadFromCache() error: %v", err)
	}

	//сравниваем
	if reflect.DeepEqual(m1, m2) != true {
		t.Errorf("TestReadFromCache() error: m1 != m2")
	}

}