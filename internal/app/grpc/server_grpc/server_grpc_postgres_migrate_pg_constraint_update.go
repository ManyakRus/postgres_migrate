//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_constraint"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/postgres_gorm"
	"time"
)

// PostgresMigratePgConstraint_UpdateManyFields - изменяет только нужные колонки в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_UpdateManyFields(ctx context.Context, Request *grpc_proto.Request_Model_MassString) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_constraint.UpdateManyFields_ctx(ctx, db, &m, Request.MassNames)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Condeferrable - изменяет колонку Condeferrable в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Condeferrable(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Condeferrable := Request.Bool_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Condeferrable = Condeferrable
	err = crud_postgres_migrate_pg_constraint.Update_Condeferrable_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Condeferred - изменяет колонку Condeferred в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Condeferred(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Condeferred := Request.Bool_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Condeferred = Condeferred
	err = crud_postgres_migrate_pg_constraint.Update_Condeferred_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conexclop - изменяет колонку Conexclop в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conexclop(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conexclop := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conexclop = Conexclop
	err = crud_postgres_migrate_pg_constraint.Update_Conexclop_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Confdeltype - изменяет колонку Confdeltype в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Confdeltype(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Confdeltype := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Confdeltype = Confdeltype
	err = crud_postgres_migrate_pg_constraint.Update_Confdeltype_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conffeqop - изменяет колонку Conffeqop в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conffeqop(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conffeqop := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conffeqop = Conffeqop
	err = crud_postgres_migrate_pg_constraint.Update_Conffeqop_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Confkey - изменяет колонку Confkey в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Confkey(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Confkey := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Confkey = Confkey
	err = crud_postgres_migrate_pg_constraint.Update_Confkey_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Confmatchtype - изменяет колонку Confmatchtype в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Confmatchtype(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Confmatchtype := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Confmatchtype = Confmatchtype
	err = crud_postgres_migrate_pg_constraint.Update_Confmatchtype_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Confrelid - изменяет колонку Confrelid в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Confrelid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Confrelid := Request.Int64_3
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Confrelid = Confrelid
	err = crud_postgres_migrate_pg_constraint.Update_Confrelid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Confupdtype - изменяет колонку Confupdtype в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Confupdtype(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Confupdtype := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Confupdtype = Confupdtype
	err = crud_postgres_migrate_pg_constraint.Update_Confupdtype_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conindid - изменяет колонку Conindid в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conindid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conindid := Request.Int64_3
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conindid = Conindid
	err = crud_postgres_migrate_pg_constraint.Update_Conindid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Coninhcount - изменяет колонку Coninhcount в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Coninhcount(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Coninhcount := Request.Int32_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Coninhcount = Coninhcount
	err = crud_postgres_migrate_pg_constraint.Update_Coninhcount_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conislocal - изменяет колонку Conislocal в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conislocal(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conislocal := Request.Bool_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conislocal = Conislocal
	err = crud_postgres_migrate_pg_constraint.Update_Conislocal_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conkey - изменяет колонку Conkey в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conkey(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conkey := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conkey = Conkey
	err = crud_postgres_migrate_pg_constraint.Update_Conkey_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conname - изменяет колонку Conname в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conname(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conname := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conname = Conname
	err = crud_postgres_migrate_pg_constraint.Update_Conname_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Connamespace - изменяет колонку Connamespace в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Connamespace(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Connamespace := Request.Int64_3
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Connamespace = Connamespace
	err = crud_postgres_migrate_pg_constraint.Update_Connamespace_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Connoinherit - изменяет колонку Connoinherit в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Connoinherit(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Connoinherit := Request.Bool_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Connoinherit = Connoinherit
	err = crud_postgres_migrate_pg_constraint.Update_Connoinherit_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conparentid - изменяет колонку Conparentid в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conparentid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conparentid := Request.Int64_3
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conparentid = Conparentid
	err = crud_postgres_migrate_pg_constraint.Update_Conparentid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conpfeqop - изменяет колонку Conpfeqop в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conpfeqop(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conpfeqop := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conpfeqop = Conpfeqop
	err = crud_postgres_migrate_pg_constraint.Update_Conpfeqop_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conppeqop - изменяет колонку Conppeqop в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conppeqop(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conppeqop := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conppeqop = Conppeqop
	err = crud_postgres_migrate_pg_constraint.Update_Conppeqop_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Conrelid - изменяет колонку Conrelid в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Conrelid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Conrelid := Request.Int64_3
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Conrelid = Conrelid
	err = crud_postgres_migrate_pg_constraint.Update_Conrelid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Contype - изменяет колонку Contype в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Contype(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Contype := Request.String_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Contype = Contype
	err = crud_postgres_migrate_pg_constraint.Update_Contype_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Contypid - изменяет колонку Contypid в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Contypid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Contypid := Request.Int64_3
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Contypid = Contypid
	err = crud_postgres_migrate_pg_constraint.Update_Contypid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Convalidated - изменяет колонку Convalidated в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Convalidated(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Convalidated := Request.Bool_1
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Convalidated = Convalidated
	err = crud_postgres_migrate_pg_constraint.Update_Convalidated_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_Oid - изменяет колонку Oid в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_Oid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Oid = Oid
	err = crud_postgres_migrate_pg_constraint.Update_Oid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update_VersionID - изменяет колонку VersionID в базе данных
func (s *ServerGRPC) PostgresMigratePgConstraint_Update_VersionID(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_constraint.PostgresMigratePgConstraint{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.VersionID = VersionID
	err = crud_postgres_migrate_pg_constraint.Update_VersionID_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}
