//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
	"context"
	"github.com/ManyakRus/starter/postgres_gorm"
)

// PostgresMigratePgAttribute_Read - читает и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgAttribute_Read(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	err = crud_postgres_migrate_pg_attribute.Read_ctx(ctx, db, m)
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

// PostgresMigratePgAttribute_Delete - записывает в БД is_deleted = true и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgAttribute_Delete(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	err = crud_postgres_migrate_pg_attribute.Delete_ctx(ctx, db, m)
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

// PostgresMigratePgAttribute_Restore - записывает в БД is_deleted = false и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgAttribute_Restore(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	err = crud_postgres_migrate_pg_attribute.Restore_ctx(ctx, db, m)
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

// PostgresMigratePgAttribute_Create - создаёт новую запись в БД
func (s *ServerGRPC) PostgresMigratePgAttribute_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_attribute.Create_ctx(ctx, db, m)
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

// PostgresMigratePgAttribute_Update - обновляет новую запись в БД
func (s *ServerGRPC) PostgresMigratePgAttribute_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	m := &postgres_migrate_pg_attribute.PostgresMigratePgAttribute{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_pg_attribute.Update_ctx(ctx, db, m)
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

// PostgresMigratePgAttribute_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) PostgresMigratePgAttribute_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	err = crud_postgres_migrate_pg_attribute.Save_ctx(ctx, db, &m)
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
