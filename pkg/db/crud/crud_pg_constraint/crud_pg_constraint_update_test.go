//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_constraint"
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
	m := pg_constraint.PgConstraint{}
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

func TestUpdate_Condeferrable(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Condeferrable() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Condeferrable = m.Condeferrable
	err = crud.Update_Condeferrable(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Condeferrable() Update() error: ", err)
	}
}

func TestUpdate_Condeferred(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Condeferred() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Condeferred = m.Condeferred
	err = crud.Update_Condeferred(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Condeferred() Update() error: ", err)
	}
}

func TestUpdate_Conexclop(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conexclop() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conexclop = m.Conexclop
	err = crud.Update_Conexclop(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conexclop() Update() error: ", err)
	}
}

func TestUpdate_Confdeltype(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Confdeltype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Confdeltype = m.Confdeltype
	err = crud.Update_Confdeltype(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Confdeltype() Update() error: ", err)
	}
}

func TestUpdate_Conffeqop(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conffeqop() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conffeqop = m.Conffeqop
	err = crud.Update_Conffeqop(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conffeqop() Update() error: ", err)
	}
}

func TestUpdate_Confkey(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Confkey() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Confkey = m.Confkey
	err = crud.Update_Confkey(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Confkey() Update() error: ", err)
	}
}

func TestUpdate_Confmatchtype(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Confmatchtype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Confmatchtype = m.Confmatchtype
	err = crud.Update_Confmatchtype(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Confmatchtype() Update() error: ", err)
	}
}

func TestUpdate_Confrelid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Confrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Confrelid = m.Confrelid
	err = crud.Update_Confrelid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Confrelid() Update() error: ", err)
	}
}

func TestUpdate_Confupdtype(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Confupdtype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Confupdtype = m.Confupdtype
	err = crud.Update_Confupdtype(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Confupdtype() Update() error: ", err)
	}
}

func TestUpdate_Conindid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conindid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conindid = m.Conindid
	err = crud.Update_Conindid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conindid() Update() error: ", err)
	}
}

func TestUpdate_Coninhcount(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Coninhcount() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Coninhcount = m.Coninhcount
	err = crud.Update_Coninhcount(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Coninhcount() Update() error: ", err)
	}
}

func TestUpdate_Conislocal(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conislocal() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conislocal = m.Conislocal
	err = crud.Update_Conislocal(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conislocal() Update() error: ", err)
	}
}

func TestUpdate_Conkey(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conkey() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conkey = m.Conkey
	err = crud.Update_Conkey(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conkey() Update() error: ", err)
	}
}

func TestUpdate_Conname(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conname() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conname = m.Conname
	err = crud.Update_Conname(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conname() Update() error: ", err)
	}
}

func TestUpdate_Connamespace(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Connamespace() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Connamespace = m.Connamespace
	err = crud.Update_Connamespace(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Connamespace() Update() error: ", err)
	}
}

func TestUpdate_Connoinherit(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Connoinherit() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Connoinherit = m.Connoinherit
	err = crud.Update_Connoinherit(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Connoinherit() Update() error: ", err)
	}
}

func TestUpdate_Conparentid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conparentid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conparentid = m.Conparentid
	err = crud.Update_Conparentid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conparentid() Update() error: ", err)
	}
}

func TestUpdate_Conpfeqop(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conpfeqop() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conpfeqop = m.Conpfeqop
	err = crud.Update_Conpfeqop(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conpfeqop() Update() error: ", err)
	}
}

func TestUpdate_Conppeqop(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conppeqop() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conppeqop = m.Conppeqop
	err = crud.Update_Conppeqop(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conppeqop() Update() error: ", err)
	}
}

func TestUpdate_Conrelid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Conrelid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Conrelid = m.Conrelid
	err = crud.Update_Conrelid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Conrelid() Update() error: ", err)
	}
}

func TestUpdate_Contype(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Contype() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Contype = m.Contype
	err = crud.Update_Contype(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Contype() Update() error: ", err)
	}
}

func TestUpdate_Contypid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Contypid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Contypid = m.Contypid
	err = crud.Update_Contypid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Contypid() Update() error: ", err)
	}
}

func TestUpdate_Convalidated(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Convalidated() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Convalidated = m.Convalidated
	err = crud.Update_Convalidated(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Convalidated() Update() error: ", err)
	}
}

func TestUpdate_Oid(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_Oid() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.Oid = m.Oid
	err = crud.Update_Oid(&Otvet)
	if err != nil {
		t.Error("TestUpdate_Oid() Update() error: ", err)
	}
}

func TestUpdate_VersionID(t *testing.T) {
	t.SkipNow() //now rows in DB

	config_main.LoadEnv()

	postgres_gorm.Connect_WithApplicationName(constants.SERVICE_NAME + "_test")
	defer postgres_gorm.CloseConnection()

	//прочитаем из БД
	crud := Crud_DB{}
	m := pg_constraint.PgConstraint{}
	m.Oid = OID_Test
	m.VersionID = VERSIONID_Test
	err := crud.Read(&m)
	if err != nil {
		t.Error("TestUpdate_VersionID() Read() error: ", err)
	}

	//запишем в БД это же значение
	Otvet := pg_constraint.PgConstraint{}
	Otvet.Oid = OID_Test
	Otvet.VersionID = VERSIONID_Test
	Otvet.VersionID = m.VersionID
	err = crud.Update_VersionID(&Otvet)
	if err != nil {
		t.Error("TestUpdate_VersionID() Update() error: ", err)
	}
}
