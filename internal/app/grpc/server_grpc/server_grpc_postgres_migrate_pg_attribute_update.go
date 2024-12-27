//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
	"context"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/postgres_gorm"
	"time"
)

// PostgresMigratePgAttribute_UpdateManyFields - изменяет только нужные колонки в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_UpdateManyFields(ctx context.Context, Request *grpc_proto.Request_Model_MassString) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_attribute.UpdateManyFields_ctx(ctx, db, &m, Request.MassNames)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attalign - изменяет колонку Attalign в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attalign(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attalign := Request.String_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attalign = Attalign
	err = crud_postgres_migrate_pg_attribute.Update_Attalign_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attbyval - изменяет колонку Attbyval в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attbyval(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attbyval := Request.Bool_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attbyval = Attbyval
	err = crud_postgres_migrate_pg_attribute.Update_Attbyval_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attcacheoff - изменяет колонку Attcacheoff в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attcacheoff(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attcacheoff := Request.Int32_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attcacheoff = Attcacheoff
	err = crud_postgres_migrate_pg_attribute.Update_Attcacheoff_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attcollation - изменяет колонку Attcollation в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attcollation(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attcollation := Request.Int64_3
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attcollation = Attcollation
	err = crud_postgres_migrate_pg_attribute.Update_Attcollation_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attgenerated - изменяет колонку Attgenerated в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attgenerated(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attgenerated := Request.String_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attgenerated = Attgenerated
	err = crud_postgres_migrate_pg_attribute.Update_Attgenerated_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Atthasdef - изменяет колонку Atthasdef в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Atthasdef(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Atthasdef := Request.Bool_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Atthasdef = Atthasdef
	err = crud_postgres_migrate_pg_attribute.Update_Atthasdef_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Atthasmissing - изменяет колонку Atthasmissing в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Atthasmissing(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Atthasmissing := Request.Bool_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Atthasmissing = Atthasmissing
	err = crud_postgres_migrate_pg_attribute.Update_Atthasmissing_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attidentity - изменяет колонку Attidentity в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attidentity(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attidentity := Request.String_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attidentity = Attidentity
	err = crud_postgres_migrate_pg_attribute.Update_Attidentity_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attinhcount - изменяет колонку Attinhcount в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attinhcount(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attinhcount := Request.Int32_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attinhcount = Attinhcount
	err = crud_postgres_migrate_pg_attribute.Update_Attinhcount_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attisdropped - изменяет колонку Attisdropped в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attisdropped(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attisdropped := Request.Bool_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attisdropped = Attisdropped
	err = crud_postgres_migrate_pg_attribute.Update_Attisdropped_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attislocal - изменяет колонку Attislocal в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attislocal(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attislocal := Request.Bool_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attislocal = Attislocal
	err = crud_postgres_migrate_pg_attribute.Update_Attislocal_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attlen - изменяет колонку Attlen в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attlen(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attlen := Request.Int32_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attlen = Attlen
	err = crud_postgres_migrate_pg_attribute.Update_Attlen_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attmissingval - изменяет колонку Attmissingval в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attmissingval(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attmissingval := Request.String_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attmissingval = Attmissingval
	err = crud_postgres_migrate_pg_attribute.Update_Attmissingval_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attname - изменяет колонку Attname в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attname(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attname = Attname
	err = crud_postgres_migrate_pg_attribute.Update_Attname_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attndims - изменяет колонку Attndims в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attndims(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attndims := Request.Int32_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attndims = Attndims
	err = crud_postgres_migrate_pg_attribute.Update_Attndims_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attnotnull - изменяет колонку Attnotnull в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attnotnull(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attnotnull := Request.Bool_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attnotnull = Attnotnull
	err = crud_postgres_migrate_pg_attribute.Update_Attnotnull_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attnum - изменяет колонку Attnum в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attnum(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attnum := Request.Int32_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attnum = Attnum
	err = crud_postgres_migrate_pg_attribute.Update_Attnum_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attrelid - изменяет колонку Attrelid в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attrelid(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attrelid = Attrelid
	err = crud_postgres_migrate_pg_attribute.Update_Attrelid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attstattarget - изменяет колонку Attstattarget в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attstattarget(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attstattarget := Request.Int32_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attstattarget = Attstattarget
	err = crud_postgres_migrate_pg_attribute.Update_Attstattarget_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Attstorage - изменяет колонку Attstorage в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Attstorage(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Attstorage := Request.String_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Attstorage = Attstorage
	err = crud_postgres_migrate_pg_attribute.Update_Attstorage_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Atttypid - изменяет колонку Atttypid в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Atttypid(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Atttypid := Request.Int64_3
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Atttypid = Atttypid
	err = crud_postgres_migrate_pg_attribute.Update_Atttypid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_Atttypmod - изменяет колонку Atttypmod в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_Atttypmod(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Atttypmod := Request.Int32_1
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.Atttypmod = Atttypmod
	err = crud_postgres_migrate_pg_attribute.Update_Atttypmod_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgAttribute_Update_VersionID - изменяет колонку VersionID в базе данных
func (s *ServerGRPC) PostgresMigratePgAttribute_Update_VersionID(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_attribute.PostgresMigratePgAttribute{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	m.VersionID = VersionID
	err = crud_postgres_migrate_pg_attribute.Update_VersionID_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}
