//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestCrud_GRPC_UpdateManyFields(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	//прочитаем из БД
	crud := grpc_pg_attribute.Crud_GRPC{}
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = ATTNAME_Test
	Otvet.Attrelid = ATTRELID_Test
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

	if (Otvet.Attname == "") ||  (Otvet.Attrelid == 0) ||  (Otvet.VersionID == 0) {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ID =0")
	}
}

func TestCrud_GRPC_Update_Attalign(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attalign() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attalign = m.Attalign
	err = crud.Update_Attalign(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attalign() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attbyval(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attbyval() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attbyval = m.Attbyval
	err = crud.Update_Attbyval(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attbyval() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attcacheoff(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attcacheoff() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attcacheoff = m.Attcacheoff
	err = crud.Update_Attcacheoff(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attcacheoff() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attcollation(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attcollation() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attcollation = m.Attcollation
	err = crud.Update_Attcollation(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attcollation() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attgenerated(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attgenerated() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attgenerated = m.Attgenerated
	err = crud.Update_Attgenerated(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attgenerated() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Atthasdef(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Atthasdef() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Atthasdef = m.Atthasdef
	err = crud.Update_Atthasdef(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Atthasdef() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Atthasmissing(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Atthasmissing() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Atthasmissing = m.Atthasmissing
	err = crud.Update_Atthasmissing(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Atthasmissing() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attidentity(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attidentity() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attidentity = m.Attidentity
	err = crud.Update_Attidentity(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attidentity() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attinhcount(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attinhcount() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attinhcount = m.Attinhcount
	err = crud.Update_Attinhcount(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attinhcount() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attisdropped(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attisdropped() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attisdropped = m.Attisdropped
	err = crud.Update_Attisdropped(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attisdropped() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attislocal(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attislocal() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attislocal = m.Attislocal
	err = crud.Update_Attislocal(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attislocal() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attlen(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attlen() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attlen = m.Attlen
	err = crud.Update_Attlen(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attlen() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attname(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attname() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attname = m.Attname
	err = crud.Update_Attname(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attname() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attndims(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attndims() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attndims = m.Attndims
	err = crud.Update_Attndims(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attndims() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attnotnull(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attnotnull() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attnotnull = m.Attnotnull
	err = crud.Update_Attnotnull(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attnotnull() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attnum(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attnum() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attnum = m.Attnum
	err = crud.Update_Attnum(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attnum() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attrelid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attrelid = m.Attrelid
	err = crud.Update_Attrelid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attrelid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attstattarget(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attstattarget() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attstattarget = m.Attstattarget
	err = crud.Update_Attstattarget(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attstattarget() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Attstorage(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attstorage() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Attstorage = m.Attstorage
	err = crud.Update_Attstorage(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Attstorage() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Atttypid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Atttypid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Atttypid = m.Atttypid
	err = crud.Update_Atttypid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Atttypid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Atttypmod(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Atttypmod() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.Atttypmod = m.Atttypmod
	err = crud.Update_Atttypmod(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Atttypmod() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_pg_attribute.Crud_GRPC{}

	//прочитаем из БД
	m := pg_attribute.PgAttribute{}
	m.Attname = ATTNAME_Test
	m.Attrelid = ATTRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_attribute.PgAttribute{}
	Otvet.Attname = m.Attname
	Otvet.Attrelid = m.Attrelid
	Otvet.VersionID = m.VersionID
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Update() error: ", err)
	}
}
