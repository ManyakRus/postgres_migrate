//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"context"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/postgres_gorm"
	"time"
)

// PostgresMigratePgIndex_UpdateManyFields - изменяет только нужные колонки в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_UpdateManyFields(ctx context.Context, Request *grpc_proto.Request_Model_MassString) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_index.UpdateManyFields_ctx(ctx, db, &m, Request.MassNames)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indcheckxmin - изменяет колонку Indcheckxmin в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indcheckxmin(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indcheckxmin := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indcheckxmin = Indcheckxmin
	err = crud_postgres_migrate_pg_index.Update_Indcheckxmin_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indclass - изменяет колонку Indclass в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indclass(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indclass := Request.String_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indclass = Indclass
	err = crud_postgres_migrate_pg_index.Update_Indclass_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indcollation - изменяет колонку Indcollation в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indcollation(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indcollation := Request.String_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indcollation = Indcollation
	err = crud_postgres_migrate_pg_index.Update_Indcollation_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indexprs - изменяет колонку Indexprs в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indexprs(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indexprs := Request.String_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indexprs = Indexprs
	err = crud_postgres_migrate_pg_index.Update_Indexprs_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indexrelid - изменяет колонку Indexrelid в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indexrelid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indexrelid = Indexrelid
	err = crud_postgres_migrate_pg_index.Update_Indexrelid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indimmediate - изменяет колонку Indimmediate в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indimmediate(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indimmediate := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indimmediate = Indimmediate
	err = crud_postgres_migrate_pg_index.Update_Indimmediate_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indisclustered - изменяет колонку Indisclustered в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indisclustered(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indisclustered := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indisclustered = Indisclustered
	err = crud_postgres_migrate_pg_index.Update_Indisclustered_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indisexclusion - изменяет колонку Indisexclusion в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indisexclusion(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indisexclusion := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indisexclusion = Indisexclusion
	err = crud_postgres_migrate_pg_index.Update_Indisexclusion_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indislive - изменяет колонку Indislive в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indislive(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indislive := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indislive = Indislive
	err = crud_postgres_migrate_pg_index.Update_Indislive_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indisprimary - изменяет колонку Indisprimary в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indisprimary(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indisprimary := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indisprimary = Indisprimary
	err = crud_postgres_migrate_pg_index.Update_Indisprimary_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indisready - изменяет колонку Indisready в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indisready(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indisready := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indisready = Indisready
	err = crud_postgres_migrate_pg_index.Update_Indisready_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indisreplident - изменяет колонку Indisreplident в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indisreplident(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indisreplident := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indisreplident = Indisreplident
	err = crud_postgres_migrate_pg_index.Update_Indisreplident_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indisunique - изменяет колонку Indisunique в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indisunique(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indisunique := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indisunique = Indisunique
	err = crud_postgres_migrate_pg_index.Update_Indisunique_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indisvalid - изменяет колонку Indisvalid в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indisvalid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Bool) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indisvalid := Request.Bool_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indisvalid = Indisvalid
	err = crud_postgres_migrate_pg_index.Update_Indisvalid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indkey - изменяет колонку Indkey в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indkey(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indkey := Request.String_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indkey = Indkey
	err = crud_postgres_migrate_pg_index.Update_Indkey_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indnatts - изменяет колонку Indnatts в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indnatts(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indnatts := Request.Int32_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indnatts = Indnatts
	err = crud_postgres_migrate_pg_index.Update_Indnatts_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indnkeyatts - изменяет колонку Indnkeyatts в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indnkeyatts(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indnkeyatts := Request.Int32_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indnkeyatts = Indnkeyatts
	err = crud_postgres_migrate_pg_index.Update_Indnkeyatts_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indoption - изменяет колонку Indoption в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indoption(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indoption := Request.String_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indoption = Indoption
	err = crud_postgres_migrate_pg_index.Update_Indoption_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indpred - изменяет колонку Indpred в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indpred(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indpred := Request.String_1
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indpred = Indpred
	err = crud_postgres_migrate_pg_index.Update_Indpred_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_Indrelid - изменяет колонку Indrelid в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_Indrelid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	Indrelid := Request.Int64_3
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.Indrelid = Indrelid
	err = crud_postgres_migrate_pg_index.Update_Indrelid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgIndex_Update_VersionID - изменяет колонку VersionID в базе данных
func (s *ServerGRPC) PostgresMigratePgIndex_Update_VersionID(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_index.PostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_index.PostgresMigratePgIndex{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_index.PostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.VersionID = VersionID

	m.VersionID = VersionID
	err = crud_postgres_migrate_pg_index.Update_VersionID_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}
