//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
	"context"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/postgres_gorm"
	"time"
)

// PostgresMigratePgDescription_UpdateManyFields - изменяет только нужные колонки в базе данных
func (s *ServerGRPC) PostgresMigratePgDescription_UpdateManyFields(ctx context.Context, Request *grpc_proto.Request_Model_MassString) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_description.PostgresMigratePgDescription{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_description.UpdateManyFields_ctx(ctx, db, &m, Request.MassNames)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgDescription_Update_Classoid - изменяет колонку Classoid в базе данных
func (s *ServerGRPC) PostgresMigratePgDescription_Update_Classoid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_description.PostgresMigratePgDescription{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Classoid := Request.Int64_1
	Objoid := Request.Int64_2
	Objsubid := Request.Int32_1
	VersionID := Request.Int64_3
	m := &postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = Classoid
	m.Objoid = Objoid
	m.Objsubid = Objsubid
	m.VersionID = VersionID

	m.Classoid = Classoid
	err = crud_postgres_migrate_pg_description.Update_Classoid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgDescription_Update_Description - изменяет колонку Description в базе данных
func (s *ServerGRPC) PostgresMigratePgDescription_Update_Description(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32_Int64_String) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_description.PostgresMigratePgDescription{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Classoid := Request.Int64_1
	Objoid := Request.Int64_2
	Objsubid := Request.Int32_1
	VersionID := Request.Int64_3
	Description := Request.String_1
	m := &postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = Classoid
	m.Objoid = Objoid
	m.Objsubid = Objsubid
	m.VersionID = VersionID

	m.Description = Description
	err = crud_postgres_migrate_pg_description.Update_Description_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgDescription_Update_Objoid - изменяет колонку Objoid в базе данных
func (s *ServerGRPC) PostgresMigratePgDescription_Update_Objoid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_description.PostgresMigratePgDescription{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Classoid := Request.Int64_1
	Objoid := Request.Int64_2
	Objsubid := Request.Int32_1
	VersionID := Request.Int64_3
	m := &postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = Classoid
	m.Objoid = Objoid
	m.Objsubid = Objsubid
	m.VersionID = VersionID

	m.Objoid = Objoid
	err = crud_postgres_migrate_pg_description.Update_Objoid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgDescription_Update_Objsubid - изменяет колонку Objsubid в базе данных
func (s *ServerGRPC) PostgresMigratePgDescription_Update_Objsubid(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_description.PostgresMigratePgDescription{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Classoid := Request.Int64_1
	Objoid := Request.Int64_2
	Objsubid := Request.Int32_1
	VersionID := Request.Int64_3
	m := &postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = Classoid
	m.Objoid = Objoid
	m.Objsubid = Objsubid
	m.VersionID = VersionID

	m.Objsubid = Objsubid
	err = crud_postgres_migrate_pg_description.Update_Objsubid_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}

// PostgresMigratePgDescription_Update_VersionID - изменяет колонку VersionID в базе данных
func (s *ServerGRPC) PostgresMigratePgDescription_Update_VersionID(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32_Int64) (*grpc_proto.ResponseEmpty, error) {
	var Otvet grpc_proto.ResponseEmpty
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_pg_description.PostgresMigratePgDescription{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_pg_description.PostgresMigratePgDescription{})
		return &Otvet, err
	}

	ctxMain := contextmain.GetContext()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(db_constants.TIMEOUT_DB_SECONDS))
	defer ctxCancelFunc()

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Classoid := Request.Int64_1
	Objoid := Request.Int64_2
	Objsubid := Request.Int32_1
	VersionID := Request.Int64_3
	m := &postgres_migrate_pg_description.PostgresMigratePgDescription{}
	m.Classoid = Classoid
	m.Objoid = Objoid
	m.Objsubid = Objsubid
	m.VersionID = VersionID

	m.VersionID = VersionID
	err = crud_postgres_migrate_pg_description.Update_VersionID_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	return &Otvet, err
}
