//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestCrud_GRPC_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	//прочитаем из БД
	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
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

func TestCrud_GRPC_Update_Condeferrable(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Condeferrable() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Condeferrable = m.Condeferrable
	err = crud.Update_Condeferrable(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Condeferrable() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Condeferred(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Condeferred() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Condeferred = m.Condeferred
	err = crud.Update_Condeferred(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Condeferred() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conexclop(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conexclop() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conexclop = m.Conexclop
	err = crud.Update_Conexclop(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conexclop() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Confdeltype(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confdeltype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Confdeltype = m.Confdeltype
	err = crud.Update_Confdeltype(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confdeltype() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conffeqop(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conffeqop() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conffeqop = m.Conffeqop
	err = crud.Update_Conffeqop(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conffeqop() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Confkey(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confkey() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Confkey = m.Confkey
	err = crud.Update_Confkey(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confkey() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Confmatchtype(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confmatchtype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Confmatchtype = m.Confmatchtype
	err = crud.Update_Confmatchtype(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confmatchtype() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Confrelid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Confrelid = m.Confrelid
	err = crud.Update_Confrelid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confrelid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Confupdtype(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confupdtype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Confupdtype = m.Confupdtype
	err = crud.Update_Confupdtype(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Confupdtype() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conindid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conindid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conindid = m.Conindid
	err = crud.Update_Conindid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conindid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Coninhcount(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Coninhcount() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Coninhcount = m.Coninhcount
	err = crud.Update_Coninhcount(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Coninhcount() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conislocal(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conislocal() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conislocal = m.Conislocal
	err = crud.Update_Conislocal(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conislocal() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conkey(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conkey() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conkey = m.Conkey
	err = crud.Update_Conkey(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conkey() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conname(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conname() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conname = m.Conname
	err = crud.Update_Conname(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conname() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Connamespace(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Connamespace() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Connamespace = m.Connamespace
	err = crud.Update_Connamespace(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Connamespace() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Connoinherit(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Connoinherit() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Connoinherit = m.Connoinherit
	err = crud.Update_Connoinherit(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Connoinherit() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conparentid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conparentid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conparentid = m.Conparentid
	err = crud.Update_Conparentid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conparentid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conpfeqop(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conpfeqop() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conpfeqop = m.Conpfeqop
	err = crud.Update_Conpfeqop(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conpfeqop() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conppeqop(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conppeqop() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conppeqop = m.Conppeqop
	err = crud.Update_Conppeqop(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conppeqop() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Conrelid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Conrelid = m.Conrelid
	err = crud.Update_Conrelid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Conrelid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Contype(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Contype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Contype = m.Contype
	err = crud.Update_Contype(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Contype() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Contypid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Contypid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Contypid = m.Contypid
	err = crud.Update_Contypid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Contypid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Convalidated(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Convalidated() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Convalidated = m.Convalidated
	err = crud.Update_Convalidated(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Convalidated() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Oid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Oid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
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

	crud := grpc_postgres_migrate_pg_constraint.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Update() error: ", err)
	}
}
