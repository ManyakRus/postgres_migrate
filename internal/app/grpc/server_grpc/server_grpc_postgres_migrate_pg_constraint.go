//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_constraint"
	"context"
	"github.com/ManyakRus/starter/postgres_gorm"
)

// PostgresMigratePgConstraint_Read - читает и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgConstraint_Read(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	err = crud_postgres_migrate_pg_constraint.Read_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	//заполяем ответ
	ModelString, err := m.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// PostgresMigratePgConstraint_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgConstraint_Delete(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	err = crud_postgres_migrate_pg_constraint.Delete_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	//заполяем ответ
	ModelString, err := m.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// PostgresMigratePgConstraint_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgConstraint_Restore(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Oid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	m.Oid = Oid
	m.VersionID = VersionID

	err = crud_postgres_migrate_pg_constraint.Restore_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	//заполяем ответ
	ModelString, err := m.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// PostgresMigratePgConstraint_Create - создаёт новую запись в БД
func (s *ServerGRPC) PostgresMigratePgConstraint_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_constraint.Create_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	//заполяем ответ
	ModelString, err := m.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// PostgresMigratePgConstraint_Update - обновляет новую запись в БД
func (s *ServerGRPC) PostgresMigratePgConstraint_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	m := &postgres_migrate_pg_constraint.PostgresMigratePgConstraint{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_constraint.Update_ctx(ctx, db, m)
	if err != nil {
		return &Otvet, err
	}

	//заполяем ответ
	ModelString, err := m.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}

// PostgresMigratePgConstraint_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) PostgresMigratePgConstraint_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	err = crud_postgres_migrate_pg_constraint.Save_ctx(ctx, db, &m)
	if err != nil {
		return &Otvet, err
	}

	//заполяем ответ
	ModelString, err := m.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}
