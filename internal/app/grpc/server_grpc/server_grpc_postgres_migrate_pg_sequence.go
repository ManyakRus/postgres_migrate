//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_sequence"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/starter/postgres_gorm"
)

// PostgresMigratePgSequence_Read - читает и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgSequence_Read(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	err = crud_postgres_migrate_pg_sequence.Read_ctx(ctx, db, m)
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

// PostgresMigratePgSequence_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgSequence_Delete(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	err = crud_postgres_migrate_pg_sequence.Delete_ctx(ctx, db, m)
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

// PostgresMigratePgSequence_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgSequence_Restore(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Seqrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	m.Seqrelid = Seqrelid
	m.VersionID = VersionID

	err = crud_postgres_migrate_pg_sequence.Restore_ctx(ctx, db, m)
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

// PostgresMigratePgSequence_Create - создаёт новую запись в БД
func (s *ServerGRPC) PostgresMigratePgSequence_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_sequence.Create_ctx(ctx, db, m)
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

// PostgresMigratePgSequence_Update - обновляет новую запись в БД
func (s *ServerGRPC) PostgresMigratePgSequence_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	m := &postgres_migrate_pg_sequence.PostgresMigratePgSequence{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_sequence.Update_ctx(ctx, db, m)
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

// PostgresMigratePgSequence_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) PostgresMigratePgSequence_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	err = crud_postgres_migrate_pg_sequence.Save_ctx(ctx, db, &m)
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
