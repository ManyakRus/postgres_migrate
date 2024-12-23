//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
	"github.com/ManyakRus/starter/postgres_gorm"
)

// PostgresMigratePgDescription_Read - читает и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgDescription_Read(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	err = crud_postgres_migrate_pg_description.Read_ctx(ctx, db, m)
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

// PostgresMigratePgDescription_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgDescription_Delete(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	err = crud_postgres_migrate_pg_description.Delete_ctx(ctx, db, m)
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

// PostgresMigratePgDescription_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgDescription_Restore(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	err = crud_postgres_migrate_pg_description.Restore_ctx(ctx, db, m)
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

// PostgresMigratePgDescription_Create - создаёт новую запись в БД
func (s *ServerGRPC) PostgresMigratePgDescription_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	m := &postgres_migrate_pg_description.PostgresMigratePgDescription{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_description.Create_ctx(ctx, db, m)
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

// PostgresMigratePgDescription_Update - обновляет новую запись в БД
func (s *ServerGRPC) PostgresMigratePgDescription_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	m := &postgres_migrate_pg_description.PostgresMigratePgDescription{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_description.Update_ctx(ctx, db, m)
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

// PostgresMigratePgDescription_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) PostgresMigratePgDescription_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	err = crud_postgres_migrate_pg_description.Save_ctx(ctx, db, &m)
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
