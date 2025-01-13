//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestCrud_GRPC_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	//прочитаем из БД
	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
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

	if (Otvet.Oid == 0) || (Otvet.VersionID == 0) {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ID =0")
	}
}

func TestCrud_GRPC_Update_Oid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Oid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Oid = m.Oid
	err = crud.Update_Oid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Oid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relallvisible(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relallvisible() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relallvisible = m.Relallvisible
	err = crud.Update_Relallvisible(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relallvisible() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relam(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relam() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relam = m.Relam
	err = crud.Update_Relam(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relam() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relchecks(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relchecks() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relchecks = m.Relchecks
	err = crud.Update_Relchecks(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relchecks() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relfilenode(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relfilenode() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relfilenode = m.Relfilenode
	err = crud.Update_Relfilenode(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relfilenode() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relforcerowsecurity(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relforcerowsecurity() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relforcerowsecurity = m.Relforcerowsecurity
	err = crud.Update_Relforcerowsecurity(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relforcerowsecurity() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relfrozenxid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relfrozenxid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relfrozenxid = m.Relfrozenxid
	err = crud.Update_Relfrozenxid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relfrozenxid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relhasindex(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relhasindex() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relhasindex = m.Relhasindex
	err = crud.Update_Relhasindex(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relhasindex() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relhasrules(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relhasrules() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relhasrules = m.Relhasrules
	err = crud.Update_Relhasrules(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relhasrules() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relhassubclass(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relhassubclass() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relhassubclass = m.Relhassubclass
	err = crud.Update_Relhassubclass(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relhassubclass() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relhastriggers(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relhastriggers() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relhastriggers = m.Relhastriggers
	err = crud.Update_Relhastriggers(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relhastriggers() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relispartition(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relispartition() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relispartition = m.Relispartition
	err = crud.Update_Relispartition(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relispartition() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relispopulated(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relispopulated() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relispopulated = m.Relispopulated
	err = crud.Update_Relispopulated(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relispopulated() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relisshared(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relisshared() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relisshared = m.Relisshared
	err = crud.Update_Relisshared(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relisshared() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relkind(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relkind() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relkind = m.Relkind
	err = crud.Update_Relkind(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relkind() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relminmxid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relminmxid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relminmxid = m.Relminmxid
	err = crud.Update_Relminmxid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relminmxid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relname(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relname() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relname = m.Relname
	err = crud.Update_Relname(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relname() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relnamespace(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relnamespace() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relnamespace = m.Relnamespace
	err = crud.Update_Relnamespace(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relnamespace() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relnatts(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relnatts() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relnatts = m.Relnatts
	err = crud.Update_Relnatts(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relnatts() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Reloftype(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reloftype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Reloftype = m.Reloftype
	err = crud.Update_Reloftype(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reloftype() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relowner(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relowner() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relowner = m.Relowner
	err = crud.Update_Relowner(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relowner() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relpages(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relpages() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relpages = m.Relpages
	err = crud.Update_Relpages(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relpages() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relpersistence(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relpersistence() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relpersistence = m.Relpersistence
	err = crud.Update_Relpersistence(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relpersistence() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relreplident(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relreplident() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relreplident = m.Relreplident
	err = crud.Update_Relreplident(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relreplident() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relrewrite(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relrewrite() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relrewrite = m.Relrewrite
	err = crud.Update_Relrewrite(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relrewrite() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Relrowsecurity(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relrowsecurity() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Relrowsecurity = m.Relrowsecurity
	err = crud.Update_Relrowsecurity(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Relrowsecurity() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Reltablespace(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reltablespace() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Reltablespace = m.Reltablespace
	err = crud.Update_Reltablespace(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reltablespace() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Reltoastrelid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reltoastrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Reltoastrelid = m.Reltoastrelid
	err = crud.Update_Reltoastrelid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reltoastrelid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Reltuples(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reltuples() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Reltuples = m.Reltuples
	err = crud.Update_Reltuples(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reltuples() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Reltype(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reltype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.Reltype = m.Reltype
	err = crud.Update_Reltype(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Reltype() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_VersionID(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_class.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = m.Oid
	Otvet.VersionID = m.VersionID
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Update() error: ", err)
	}
}
