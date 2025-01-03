//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_version"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
	"context"
	"github.com/ManyakRus/starter/postgres_gorm"
)

// PostgresMigrateVersion_Read - читает и возвращает модель из БД
func (s *ServerGRPC) PostgresMigrateVersion_Read(ctx context.Context, Request *grpc_proto.Request_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_version.PostgresMigrateVersion{})
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	ID := Request.Int64_1
	m := &postgres_migrate_version.PostgresMigrateVersion{}
	m.ID = ID

	err = crud_postgres_migrate_version.Read_ctx(ctx, db, m)
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

// PostgresMigrateVersion_Create - создаёт новую запись в БД
func (s *ServerGRPC) PostgresMigrateVersion_Create(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_version.PostgresMigrateVersion{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := &postgres_migrate_version.PostgresMigrateVersion{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_version.Create_ctx(ctx, db, m)
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

// PostgresMigrateVersion_Update - обновляет новую запись в БД
func (s *ServerGRPC) PostgresMigrateVersion_Update(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_version.PostgresMigrateVersion{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := &postgres_migrate_version.PostgresMigrateVersion{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_version.Update_ctx(ctx, db, m)
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

// PostgresMigrateVersion_Save - записывает (создаёт или обновляет) запись в БД
func (s *ServerGRPC) PostgresMigrateVersion_Save(ctx context.Context, Request *grpc_proto.RequestModel) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := postgres_migrate_version.PostgresMigrateVersion{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(postgres_migrate_version.PostgresMigrateVersion{})
		return &Otvet, err
	}

	//получим модель из строки JSON
	m := postgres_migrate_version.PostgresMigrateVersion{}
	err = m.GetModelFromJSON(Request.ModelString)
	if err != nil {
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	err = crud_postgres_migrate_version.Save_ctx(ctx, db, &m)
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
