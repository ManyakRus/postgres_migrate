//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_sequence"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestCrud_GRPC_UpdateManyFields(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	//прочитаем из БД
	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = SEQRELID_Test
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

	if (Otvet.Seqrelid == 0) || (Otvet.VersionID == 0) {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ID =0")
	}
}

func TestCrud_GRPC_Update_Seqcache(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqcache() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = m.Seqrelid
	Otvet.VersionID = m.VersionID
	Otvet.Seqcache = m.Seqcache
	err = crud.Update_Seqcache(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqcache() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Seqcycle(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqcycle() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = m.Seqrelid
	Otvet.VersionID = m.VersionID
	Otvet.Seqcycle = m.Seqcycle
	err = crud.Update_Seqcycle(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqcycle() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Seqincrement(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqincrement() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = m.Seqrelid
	Otvet.VersionID = m.VersionID
	Otvet.Seqincrement = m.Seqincrement
	err = crud.Update_Seqincrement(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqincrement() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Seqmax(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqmax() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = m.Seqrelid
	Otvet.VersionID = m.VersionID
	Otvet.Seqmax = m.Seqmax
	err = crud.Update_Seqmax(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqmax() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Seqmin(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqmin() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = m.Seqrelid
	Otvet.VersionID = m.VersionID
	Otvet.Seqmin = m.Seqmin
	err = crud.Update_Seqmin(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqmin() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Seqrelid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = m.Seqrelid
	Otvet.VersionID = m.VersionID
	Otvet.Seqrelid = m.Seqrelid
	err = crud.Update_Seqrelid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqrelid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Seqstart(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqstart() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = m.Seqrelid
	Otvet.VersionID = m.VersionID
	Otvet.Seqstart = m.Seqstart
	err = crud.Update_Seqstart(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqstart() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Seqtypid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqtypid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = m.Seqrelid
	Otvet.VersionID = m.VersionID
	Otvet.Seqtypid = m.Seqtypid
	err = crud.Update_Seqtypid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Seqtypid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_sequence.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = SEQRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	Otvet.Seqrelid = m.Seqrelid
	Otvet.VersionID = m.VersionID
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Update() error: ", err)
	}
}
