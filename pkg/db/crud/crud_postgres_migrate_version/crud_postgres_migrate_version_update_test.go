//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestUpdateManyFields(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_version.PostgresMigrateVersion{}
	m.ID = ID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdateManyFields() Read() error: ", err)
	}

	//запишем в БД это же значение
	err = crud.UpdateManyFields(&m, nil)
	if err != nil {
		t.Error("TestUpdateManyFields() UpdateManyFields() error: ", err)
	}
}

func TestUpdate_Description(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_version.PostgresMigrateVersion{}
	m.ID = ID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Description() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet.ID = ID_Test
	Otvet.Description = m.Description
	err = crud.Update_Description(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Description() Update() error: ", err)
	}
}

func TestUpdate_Name(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_version.PostgresMigrateVersion{}
	m.ID = ID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Name() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_version.PostgresMigrateVersion{}
	Otvet.ID = ID_Test
	Otvet.Name = m.Name
	err = crud.Update_Name(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Name() Update() error: ", err)
	}
}
