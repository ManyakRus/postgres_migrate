//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_sequence"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/postgres_gorm"
	"time"
)

// PostgresMigratePgSequence_UpdateManyFields - изменяет только нужные колонки в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_UpdateManyFields(ctx context.Context, Request *grpc_proto.Request_Model_MassString) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_sequence.UpdateManyFields_ctx(ctx, db, &m, Request.MassNames)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgSequence_Update_Seqcache - изменяет колонку Seqcache в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_Update_Seqcache(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Seqcache := Request.Int64_3
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	m.Seqcache = Seqcache
	err = crud_postgres_migrate_pg_sequence.Update_Seqcache_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgSequence_Update_Seqcycle - изменяет колонку Seqcycle в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_Update_Seqcycle(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Seqcycle := Request.Bool_1
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	m.Seqcycle = Seqcycle
	err = crud_postgres_migrate_pg_sequence.Update_Seqcycle_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgSequence_Update_Seqincrement - изменяет колонку Seqincrement в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_Update_Seqincrement(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Seqincrement := Request.Int64_3
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	m.Seqincrement = Seqincrement
	err = crud_postgres_migrate_pg_sequence.Update_Seqincrement_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgSequence_Update_Seqmax - изменяет колонку Seqmax в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_Update_Seqmax(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Seqmax := Request.Int64_3
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	m.Seqmax = Seqmax
	err = crud_postgres_migrate_pg_sequence.Update_Seqmax_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgSequence_Update_Seqmin - изменяет колонку Seqmin в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_Update_Seqmin(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Seqmin := Request.Int64_3
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	m.Seqmin = Seqmin
	err = crud_postgres_migrate_pg_sequence.Update_Seqmin_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgSequence_Update_Seqrelid - изменяет колонку Seqrelid в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_Update_Seqrelid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	m.Seqrelid = Seqrelid
	err = crud_postgres_migrate_pg_sequence.Update_Seqrelid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgSequence_Update_Seqstart - изменяет колонку Seqstart в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_Update_Seqstart(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Seqstart := Request.Int64_3
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	m.Seqstart = Seqstart
	err = crud_postgres_migrate_pg_sequence.Update_Seqstart_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgSequence_Update_Seqtypid - изменяет колонку Seqtypid в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_Update_Seqtypid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Seqtypid := Request.Int64_3
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	m.Seqtypid = Seqtypid
	err = crud_postgres_migrate_pg_sequence.Update_Seqtypid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgSequence_Update_VersionID - изменяет колонку VersionID в базе данных
func (s *ServerGRPC) PostgresMigratePgSequence_Update_VersionID(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_sequence.PostgresMigratePgSequence{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_sequence.PostgresMigratePgSequence{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	m.VersionID = VersionID
	err = crud_postgres_migrate_pg_sequence.Update_VersionID_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}
