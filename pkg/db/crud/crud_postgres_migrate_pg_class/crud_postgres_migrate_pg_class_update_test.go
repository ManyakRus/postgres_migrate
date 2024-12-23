//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_postgres_migrate_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestUpdateManyFields(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName_SingularTableName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
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

func TestUpdate_Oid(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Oid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Oid = m.Oid
	err = crud.Update_Oid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Oid() Update() error: ", err)
	}
}

func TestUpdate_Relallvisible(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relallvisible() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relallvisible = m.Relallvisible
	err = crud.Update_Relallvisible(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relallvisible() Update() error: ", err)
	}
}

func TestUpdate_Relam(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relam() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relam = m.Relam
	err = crud.Update_Relam(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relam() Update() error: ", err)
	}
}

func TestUpdate_Relchecks(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relchecks() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relchecks = m.Relchecks
	err = crud.Update_Relchecks(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relchecks() Update() error: ", err)
	}
}

func TestUpdate_Relfilenode(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relfilenode() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relfilenode = m.Relfilenode
	err = crud.Update_Relfilenode(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relfilenode() Update() error: ", err)
	}
}

func TestUpdate_Relforcerowsecurity(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relforcerowsecurity() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relforcerowsecurity = m.Relforcerowsecurity
	err = crud.Update_Relforcerowsecurity(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relforcerowsecurity() Update() error: ", err)
	}
}

func TestUpdate_Relfrozenxid(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relfrozenxid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relfrozenxid = m.Relfrozenxid
	err = crud.Update_Relfrozenxid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relfrozenxid() Update() error: ", err)
	}
}

func TestUpdate_Relhasindex(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relhasindex() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relhasindex = m.Relhasindex
	err = crud.Update_Relhasindex(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relhasindex() Update() error: ", err)
	}
}

func TestUpdate_Relhasrules(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relhasrules() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relhasrules = m.Relhasrules
	err = crud.Update_Relhasrules(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relhasrules() Update() error: ", err)
	}
}

func TestUpdate_Relhassubclass(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relhassubclass() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relhassubclass = m.Relhassubclass
	err = crud.Update_Relhassubclass(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relhassubclass() Update() error: ", err)
	}
}

func TestUpdate_Relhastriggers(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relhastriggers() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relhastriggers = m.Relhastriggers
	err = crud.Update_Relhastriggers(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relhastriggers() Update() error: ", err)
	}
}

func TestUpdate_Relispartition(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relispartition() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relispartition = m.Relispartition
	err = crud.Update_Relispartition(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relispartition() Update() error: ", err)
	}
}

func TestUpdate_Relispopulated(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relispopulated() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relispopulated = m.Relispopulated
	err = crud.Update_Relispopulated(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relispopulated() Update() error: ", err)
	}
}

func TestUpdate_Relisshared(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relisshared() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relisshared = m.Relisshared
	err = crud.Update_Relisshared(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relisshared() Update() error: ", err)
	}
}

func TestUpdate_Relkind(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relkind() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relkind = m.Relkind
	err = crud.Update_Relkind(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relkind() Update() error: ", err)
	}
}

func TestUpdate_Relminmxid(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relminmxid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relminmxid = m.Relminmxid
	err = crud.Update_Relminmxid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relminmxid() Update() error: ", err)
	}
}

func TestUpdate_Relname(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relname() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relname = m.Relname
	err = crud.Update_Relname(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relname() Update() error: ", err)
	}
}

func TestUpdate_Relnamespace(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relnamespace() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relnamespace = m.Relnamespace
	err = crud.Update_Relnamespace(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relnamespace() Update() error: ", err)
	}
}

func TestUpdate_Relnatts(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relnatts() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relnatts = m.Relnatts
	err = crud.Update_Relnatts(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relnatts() Update() error: ", err)
	}
}

func TestUpdate_Reloftype(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Reloftype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Reloftype = m.Reloftype
	err = crud.Update_Reloftype(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Reloftype() Update() error: ", err)
	}
}

func TestUpdate_Relowner(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relowner() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relowner = m.Relowner
	err = crud.Update_Relowner(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relowner() Update() error: ", err)
	}
}

func TestUpdate_Relpages(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relpages() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relpages = m.Relpages
	err = crud.Update_Relpages(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relpages() Update() error: ", err)
	}
}

func TestUpdate_Relpersistence(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relpersistence() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relpersistence = m.Relpersistence
	err = crud.Update_Relpersistence(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relpersistence() Update() error: ", err)
	}
}

func TestUpdate_Relreplident(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relreplident() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relreplident = m.Relreplident
	err = crud.Update_Relreplident(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relreplident() Update() error: ", err)
	}
}

func TestUpdate_Relrewrite(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relrewrite() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relrewrite = m.Relrewrite
	err = crud.Update_Relrewrite(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relrewrite() Update() error: ", err)
	}
}

func TestUpdate_Relrowsecurity(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Relrowsecurity() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Relrowsecurity = m.Relrowsecurity
	err = crud.Update_Relrowsecurity(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Relrowsecurity() Update() error: ", err)
	}
}

func TestUpdate_Reltablespace(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Reltablespace() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Reltablespace = m.Reltablespace
	err = crud.Update_Reltablespace(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Reltablespace() Update() error: ", err)
	}
}

func TestUpdate_Reltoastrelid(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Reltoastrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Reltoastrelid = m.Reltoastrelid
	err = crud.Update_Reltoastrelid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Reltoastrelid() Update() error: ", err)
	}
}

func TestUpdate_Reltuples(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Reltuples() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Reltuples = m.Reltuples
	err = crud.Update_Reltuples(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Reltuples() Update() error: ", err)
	}
}

func TestUpdate_Reltype(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Reltype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Reltype = m.Reltype
	err = crud.Update_Reltype(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Reltype() Update() error: ", err)
	}
}

func TestUpdate_VersionID(t *testing.T) {
	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := postgres_migrate_pg_class.PostgresMigratePgClass{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestUpdate_VersionID() Update() error: ", err)
	}
}
