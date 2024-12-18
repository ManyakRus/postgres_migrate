//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/postgres_gorm"
	"time"
)

// PostgresMigratePgClass_UpdateManyFields - изменяет только нужные колонки в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_UpdateManyFields(ctx context.Context, Request *grpc_proto.Request_Model_MassString) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := postgres_migrate_pg_class.PostgresMigratePgClass{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_class.UpdateManyFields_ctx(ctx, db, &m, Request.MassNames)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Oid - изменяет колонку Oid в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Oid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Oid = Oid
	err = crud_postgres_migrate_pg_class.Update_Oid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relallvisible - изменяет колонку Relallvisible в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relallvisible(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relallvisible := Request.Int32_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relallvisible = Relallvisible
	err = crud_postgres_migrate_pg_class.Update_Relallvisible_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relam - изменяет колонку Relam в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relam(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relam := Request.Int64_3
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relam = Relam
	err = crud_postgres_migrate_pg_class.Update_Relam_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relchecks - изменяет колонку Relchecks в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relchecks(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relchecks := Request.Int32_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relchecks = Relchecks
	err = crud_postgres_migrate_pg_class.Update_Relchecks_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relfilenode - изменяет колонку Relfilenode в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relfilenode(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relfilenode := Request.Int64_3
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relfilenode = Relfilenode
	err = crud_postgres_migrate_pg_class.Update_Relfilenode_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relforcerowsecurity - изменяет колонку Relforcerowsecurity в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relforcerowsecurity(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relforcerowsecurity := Request.Bool_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relforcerowsecurity = Relforcerowsecurity
	err = crud_postgres_migrate_pg_class.Update_Relforcerowsecurity_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relfrozenxid - изменяет колонку Relfrozenxid в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relfrozenxid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relfrozenxid := Request.Int32_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relfrozenxid = Relfrozenxid
	err = crud_postgres_migrate_pg_class.Update_Relfrozenxid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relhasindex - изменяет колонку Relhasindex в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relhasindex(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relhasindex := Request.Bool_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relhasindex = Relhasindex
	err = crud_postgres_migrate_pg_class.Update_Relhasindex_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relhasrules - изменяет колонку Relhasrules в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relhasrules(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relhasrules := Request.Bool_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relhasrules = Relhasrules
	err = crud_postgres_migrate_pg_class.Update_Relhasrules_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relhassubclass - изменяет колонку Relhassubclass в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relhassubclass(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relhassubclass := Request.Bool_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relhassubclass = Relhassubclass
	err = crud_postgres_migrate_pg_class.Update_Relhassubclass_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relhastriggers - изменяет колонку Relhastriggers в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relhastriggers(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relhastriggers := Request.Bool_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relhastriggers = Relhastriggers
	err = crud_postgres_migrate_pg_class.Update_Relhastriggers_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relispartition - изменяет колонку Relispartition в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relispartition(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relispartition := Request.Bool_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relispartition = Relispartition
	err = crud_postgres_migrate_pg_class.Update_Relispartition_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relispopulated - изменяет колонку Relispopulated в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relispopulated(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relispopulated := Request.Bool_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relispopulated = Relispopulated
	err = crud_postgres_migrate_pg_class.Update_Relispopulated_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relisshared - изменяет колонку Relisshared в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relisshared(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relisshared := Request.Bool_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relisshared = Relisshared
	err = crud_postgres_migrate_pg_class.Update_Relisshared_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relkind - изменяет колонку Relkind в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relkind(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relkind := Request.String_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relkind = Relkind
	err = crud_postgres_migrate_pg_class.Update_Relkind_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relminmxid - изменяет колонку Relminmxid в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relminmxid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relminmxid := Request.Int32_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relminmxid = Relminmxid
	err = crud_postgres_migrate_pg_class.Update_Relminmxid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relname - изменяет колонку Relname в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relname(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relname := Request.String_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relname = Relname
	err = crud_postgres_migrate_pg_class.Update_Relname_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relnamespace - изменяет колонку Relnamespace в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relnamespace(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relnamespace := Request.Int64_3
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relnamespace = Relnamespace
	err = crud_postgres_migrate_pg_class.Update_Relnamespace_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relnatts - изменяет колонку Relnatts в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relnatts(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relnatts := Request.Int32_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relnatts = Relnatts
	err = crud_postgres_migrate_pg_class.Update_Relnatts_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Reloftype - изменяет колонку Reloftype в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Reloftype(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Reloftype := Request.Int64_3
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Reloftype = Reloftype
	err = crud_postgres_migrate_pg_class.Update_Reloftype_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relowner - изменяет колонку Relowner в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relowner(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relowner := Request.Int64_3
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relowner = Relowner
	err = crud_postgres_migrate_pg_class.Update_Relowner_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relpages - изменяет колонку Relpages в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relpages(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relpages := Request.Int32_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relpages = Relpages
	err = crud_postgres_migrate_pg_class.Update_Relpages_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relpersistence - изменяет колонку Relpersistence в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relpersistence(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relpersistence := Request.String_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relpersistence = Relpersistence
	err = crud_postgres_migrate_pg_class.Update_Relpersistence_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relreplident - изменяет колонку Relreplident в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relreplident(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relreplident := Request.String_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relreplident = Relreplident
	err = crud_postgres_migrate_pg_class.Update_Relreplident_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relrewrite - изменяет колонку Relrewrite в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relrewrite(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relrewrite := Request.Int64_3
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relrewrite = Relrewrite
	err = crud_postgres_migrate_pg_class.Update_Relrewrite_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Relrowsecurity - изменяет колонку Relrowsecurity в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Relrowsecurity(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Relrowsecurity := Request.Bool_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Relrowsecurity = Relrowsecurity
	err = crud_postgres_migrate_pg_class.Update_Relrowsecurity_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Reltablespace - изменяет колонку Reltablespace в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Reltablespace(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Reltablespace := Request.Int64_3
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Reltablespace = Reltablespace
	err = crud_postgres_migrate_pg_class.Update_Reltablespace_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Reltoastrelid - изменяет колонку Reltoastrelid в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Reltoastrelid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Reltoastrelid := Request.Int64_3
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Reltoastrelid = Reltoastrelid
	err = crud_postgres_migrate_pg_class.Update_Reltoastrelid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Reltuples - изменяет колонку Reltuples в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Reltuples(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Float32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Reltuples := Request.Float32_1
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Reltuples = Reltuples
	err = crud_postgres_migrate_pg_class.Update_Reltuples_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_Reltype - изменяет колонку Reltype в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_Reltype(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Reltype := Request.Int64_3
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Reltype = Reltype
	err = crud_postgres_migrate_pg_class.Update_Reltype_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgClass_Update_VersionID - изменяет колонку VersionID в базе данных
func (s *ServerGRPC) PostgresMigratePgClass_Update_VersionID(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_class.PostgresMigratePgClass{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_class.PostgresMigratePgClass{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_class.PostgresMigratePgClass{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.VersionID = VersionID
	err = crud_postgres_migrate_pg_class.Update_VersionID_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}
