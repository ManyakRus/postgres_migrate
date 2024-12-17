//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_func"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter"
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_index"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestReadObject(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()
	crud_starter.InitCrudTransport_DB()

	crud := crud_object_pg_index.Crud_DB{}
	Otvet := object_pg_index.ObjectPgIndex{}
	err := crud.ReadObject(&Otvet)
	if err != nil && crud_func.IsRecordNotFound(err) == false {
		t.Error("TestReadObject() error: ", err)
	}
}
