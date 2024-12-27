//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package tests

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client"
	"github.com/ManyakRus/starter/config_main"
	"testing"
)

func TestCrud_GRPC_UpdateManyFields(t *testing.T) {
	config_main.LoadEnv()
	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	//прочитаем из БД
	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
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

	if (Otvet.Indexrelid == 0) ||  (Otvet.VersionID == 0) {
		t.Error("TestCrud_GRPC_UpdateManyFields() error: ID =0")
	}
}

func TestCrud_GRPC_Update_Indcheckxmin(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indcheckxmin() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indcheckxmin = m.Indcheckxmin
	err = crud.Update_Indcheckxmin(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indcheckxmin() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indclass(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indclass() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indclass = m.Indclass
	err = crud.Update_Indclass(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indclass() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indcollation(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indcollation() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indcollation = m.Indcollation
	err = crud.Update_Indcollation(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indcollation() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indexprs(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indexprs() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indexprs = m.Indexprs
	err = crud.Update_Indexprs(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indexprs() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indexrelid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indexrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indexrelid = m.Indexrelid
	err = crud.Update_Indexrelid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indexrelid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indimmediate(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indimmediate() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indimmediate = m.Indimmediate
	err = crud.Update_Indimmediate(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indimmediate() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indisclustered(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisclustered() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indisclustered = m.Indisclustered
	err = crud.Update_Indisclustered(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisclustered() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indisexclusion(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisexclusion() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indisexclusion = m.Indisexclusion
	err = crud.Update_Indisexclusion(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisexclusion() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indislive(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indislive() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indislive = m.Indislive
	err = crud.Update_Indislive(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indislive() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indisprimary(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisprimary() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indisprimary = m.Indisprimary
	err = crud.Update_Indisprimary(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisprimary() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indisready(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisready() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indisready = m.Indisready
	err = crud.Update_Indisready(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisready() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indisreplident(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisreplident() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indisreplident = m.Indisreplident
	err = crud.Update_Indisreplident(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisreplident() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indisunique(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisunique() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indisunique = m.Indisunique
	err = crud.Update_Indisunique(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisunique() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indisvalid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisvalid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indisvalid = m.Indisvalid
	err = crud.Update_Indisvalid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indisvalid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indkey(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indkey() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indkey = m.Indkey
	err = crud.Update_Indkey(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indkey() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indnatts(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indnatts() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indnatts = m.Indnatts
	err = crud.Update_Indnatts(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indnatts() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indnkeyatts(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indnkeyatts() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indnkeyatts = m.Indnkeyatts
	err = crud.Update_Indnkeyatts(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indnkeyatts() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indoption(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indoption() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indoption = m.Indoption
	err = crud.Update_Indoption(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indoption() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indpred(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indpred() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indpred = m.Indpred
	err = crud.Update_Indpred(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indpred() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_Indrelid(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.Indrelid = m.Indrelid
	err = crud.Update_Indrelid(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_Indrelid() Update() error: ", err)
	}
}

func TestCrud_GRPC_Update_VersionID(t *testing.T) {
	config_main.LoadEnv()

	grpc_client.Connect()
	defer grpc_client.CloseConnection()

	crud := grpc_postgres_migrate_pg_index.Crud_GRPC{}

	//прочитаем из БД
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)

	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Otvet.Indexrelid = m.Indexrelid
	Otvet.VersionID = m.VersionID
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestCrud_GRPC_Update_VersionID() Update() error: ", err)
	}
}
