//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestCrud_GRPC_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	//прочитаем из БД
	crud := grpc_postgres_migrate_pg_namespace.Crud_GRPC{}
	Otvet := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	Otvet.Oid = OID_Test
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

	if (Otvet.Oid == 0) ||  (Otvet.VersionID == 0) {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ID =0")
	}
}

func TestCrud_GRPC_Update_Nspacl(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_namespace.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Nspacl() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Nspacl = m.Nspacl
	err = crud.Update_Nspacl(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Nspacl() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Nspname(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_namespace.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Nspname() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Nspname = m.Nspname
	err = crud.Update_Nspname(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Nspname() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Nspowner(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_namespace.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Nspowner() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Nspowner = m.Nspowner
	err = crud.Update_Nspowner(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Nspowner() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Oid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_namespace.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Oid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Oid = m.Oid
	err = crud.Update_Oid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Oid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_VersionID(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_namespace.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Update() error: ", err)
	}
}
