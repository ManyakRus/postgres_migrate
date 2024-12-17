//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_index"
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
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
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

func TestUpdate_Indcheckxmin(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indcheckxmin() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indcheckxmin = m.Indcheckxmin
	err = crud.Update_Indcheckxmin(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indcheckxmin() Update() error: ", err)
	}
}

func TestUpdate_Indclass(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indclass() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indclass = m.Indclass
	err = crud.Update_Indclass(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indclass() Update() error: ", err)
	}
}

func TestUpdate_Indcollation(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indcollation() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indcollation = m.Indcollation
	err = crud.Update_Indcollation(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indcollation() Update() error: ", err)
	}
}

func TestUpdate_Indexprs(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indexprs() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indexprs = m.Indexprs
	err = crud.Update_Indexprs(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indexprs() Update() error: ", err)
	}
}

func TestUpdate_Indexrelid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indexrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indexrelid = m.Indexrelid
	err = crud.Update_Indexrelid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indexrelid() Update() error: ", err)
	}
}

func TestUpdate_Indimmediate(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indimmediate() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indimmediate = m.Indimmediate
	err = crud.Update_Indimmediate(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indimmediate() Update() error: ", err)
	}
}

func TestUpdate_Indisclustered(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indisclustered() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indisclustered = m.Indisclustered
	err = crud.Update_Indisclustered(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indisclustered() Update() error: ", err)
	}
}

func TestUpdate_Indisexclusion(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indisexclusion() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indisexclusion = m.Indisexclusion
	err = crud.Update_Indisexclusion(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indisexclusion() Update() error: ", err)
	}
}

func TestUpdate_Indislive(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indislive() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indislive = m.Indislive
	err = crud.Update_Indislive(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indislive() Update() error: ", err)
	}
}

func TestUpdate_Indisprimary(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indisprimary() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indisprimary = m.Indisprimary
	err = crud.Update_Indisprimary(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indisprimary() Update() error: ", err)
	}
}

func TestUpdate_Indisready(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indisready() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indisready = m.Indisready
	err = crud.Update_Indisready(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indisready() Update() error: ", err)
	}
}

func TestUpdate_Indisreplident(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indisreplident() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indisreplident = m.Indisreplident
	err = crud.Update_Indisreplident(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indisreplident() Update() error: ", err)
	}
}

func TestUpdate_Indisunique(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indisunique() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indisunique = m.Indisunique
	err = crud.Update_Indisunique(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indisunique() Update() error: ", err)
	}
}

func TestUpdate_Indisvalid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indisvalid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indisvalid = m.Indisvalid
	err = crud.Update_Indisvalid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indisvalid() Update() error: ", err)
	}
}

func TestUpdate_Indkey(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indkey() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indkey = m.Indkey
	err = crud.Update_Indkey(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indkey() Update() error: ", err)
	}
}

func TestUpdate_Indnatts(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indnatts() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indnatts = m.Indnatts
	err = crud.Update_Indnatts(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indnatts() Update() error: ", err)
	}
}

func TestUpdate_Indnkeyatts(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indnkeyatts() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indnkeyatts = m.Indnkeyatts
	err = crud.Update_Indnkeyatts(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indnkeyatts() Update() error: ", err)
	}
}

func TestUpdate_Indoption(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indoption() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indoption = m.Indoption
	err = crud.Update_Indoption(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indoption() Update() error: ", err)
	}
}

func TestUpdate_Indpred(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indpred() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indpred = m.Indpred
	err = crud.Update_Indpred(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indpred() Update() error: ", err)
	}
}

func TestUpdate_Indrelid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Indrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Indrelid = m.Indrelid
	err = crud.Update_Indrelid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Indrelid() Update() error: ", err)
	}
}

func TestUpdate_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_index.PgIndex{}
	m.Indexrelid = INDEXRELID_Test
	m.Indrelid = INDRELID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_index.PgIndex{}
	Otvet.Indexrelid = INDEXRELID_Test
	Otvet.Indrelid = INDRELID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestUpdate_VersionID() Update() error: ", err)
	}
}
