//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestCrud_GRPC_UpdateManyFields(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	//прочитаем из БД
	crud := grpc_postgres_migrate_pg_description.Crud_GRPC{}
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet.Classoid = CLASSOID_Test
	Otvet.Objoid = OBJOID_Test
	Otvet.Objsubid = OBJSUBID_Test
	Otvet.VersionID = VERSIONID_Test
	err := crud.Read(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ", err)
		return
	}

	//запишем в БД, пустой список полей (не изменится)
	err = crud.UpdateManyFields(&Otvet, nil)

	if err != nil {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ", err)
	}

	if (Otvet.Classoid == 0) || (Otvet.Objoid == 0) || (Otvet.Objsubid == 0) || (Otvet.VersionID == 0) {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ID =0")
	}
}

func TestCrud_GRPC_Update_Classoid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_description.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = CLASSOID_Test
	m.Objoid = OBJOID_Test
	m.Objsubid = OBJSUBID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Classoid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet.Classoid = m.Classoid
	Otvet.Objoid = m.Objoid
	Otvet.Objsubid = m.Objsubid
	Otvet.VersionID = m.VersionID
	Otvet.Classoid = m.Classoid
	err = crud.Update_Classoid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Classoid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Description(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_description.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = CLASSOID_Test
	m.Objoid = OBJOID_Test
	m.Objsubid = OBJSUBID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Description() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet.Classoid = m.Classoid
	Otvet.Objoid = m.Objoid
	Otvet.Objsubid = m.Objsubid
	Otvet.VersionID = m.VersionID
	Otvet.Description = m.Description
	err = crud.Update_Description(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Description() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Objoid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_description.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = CLASSOID_Test
	m.Objoid = OBJOID_Test
	m.Objsubid = OBJSUBID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Objoid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet.Classoid = m.Classoid
	Otvet.Objoid = m.Objoid
	Otvet.Objsubid = m.Objsubid
	Otvet.VersionID = m.VersionID
	Otvet.Objoid = m.Objoid
	err = crud.Update_Objoid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Objoid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Objsubid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_description.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = CLASSOID_Test
	m.Objoid = OBJOID_Test
	m.Objsubid = OBJSUBID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Objsubid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet.Classoid = m.Classoid
	Otvet.Objoid = m.Objoid
	Otvet.Objsubid = m.Objsubid
	Otvet.VersionID = m.VersionID
	Otvet.Objsubid = m.Objsubid
	err = crud.Update_Objsubid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Objsubid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_description.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = CLASSOID_Test
	m.Objoid = OBJOID_Test
	m.Objsubid = OBJSUBID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	Otvet.Classoid = m.Classoid
	Otvet.Objoid = m.Objoid
	Otvet.Objsubid = m.Objsubid
	Otvet.VersionID = m.VersionID
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Update() error: ", err)
	}
}
