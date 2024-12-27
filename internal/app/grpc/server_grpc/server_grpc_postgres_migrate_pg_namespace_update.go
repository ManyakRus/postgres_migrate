//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_namespace"
	"context"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/postgres_gorm"
	"time"
)

// PostgresMigratePgNamespace_UpdateManyFields - изменяет только нужные колонки в базе данных
func (s *ServerGRPC) PostgresMigratePgNamespace_UpdateManyFields(ctx context.Context, Request *grpc_proto.Request_Model_MassString) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_namespace.PostgresMigratePgNamespace{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_namespace.UpdateManyFields_ctx(ctx, db, &m, Request.MassNames)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgNamespace_Update_Nspacl - изменяет колонку Nspacl в базе данных
func (s *ServerGRPC) PostgresMigratePgNamespace_Update_Nspacl(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_namespace.PostgresMigratePgNamespace{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Nspacl := Request.String_1
	m := &postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Nspacl = Nspacl
	err = crud_postgres_migrate_pg_namespace.Update_Nspacl_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgNamespace_Update_Nspname - изменяет колонку Nspname в базе данных
func (s *ServerGRPC) PostgresMigratePgNamespace_Update_Nspname(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_namespace.PostgresMigratePgNamespace{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Nspname := Request.String_1
	m := &postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Nspname = Nspname
	err = crud_postgres_migrate_pg_namespace.Update_Nspname_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgNamespace_Update_Nspowner - изменяет колонку Nspowner в базе данных
func (s *ServerGRPC) PostgresMigratePgNamespace_Update_Nspowner(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_namespace.PostgresMigratePgNamespace{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	Nspowner := Request.Int64_3
	m := &postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Nspowner = Nspowner
	err = crud_postgres_migrate_pg_namespace.Update_Nspowner_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgNamespace_Update_Oid - изменяет колонку Oid в базе данных
func (s *ServerGRPC) PostgresMigratePgNamespace_Update_Oid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_namespace.PostgresMigratePgNamespace{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.Oid = Oid
	err = crud_postgres_migrate_pg_namespace.Update_Oid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgNamespace_Update_VersionID - изменяет колонку VersionID в базе данных
func (s *ServerGRPC) PostgresMigratePgNamespace_Update_VersionID(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_namespace.PostgresMigratePgNamespace{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_namespace.PostgresMigratePgNamespace{}
	m.Oid = Oid
	m.VersionID = VersionID

	m.VersionID = VersionID
	err = crud_postgres_migrate_pg_namespace.Update_VersionID_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}
