//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_index"
	"github.com/ManyakRus/starter/postgres_gorm"
)

// PostgresMigratePgIndex_ReadObject - читает и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgIndex_ReadObject(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := object_postgres_migrate_pg_index.ObjectPostgresMigratePgIndex{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(object_postgres_migrate_pg_index.ObjectPostgresMigratePgIndex{})
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Indexrelid := Request.Int64_1
	Indrelid := Request.Int64_2
	VersionID := Request.Int64_3
	m := &object_postgres_migrate_pg_index.ObjectPostgresMigratePgIndex{}
	m.Indexrelid = Indexrelid
	m.Indrelid = Indrelid
	m.VersionID = VersionID

	err = crud_object_postgres_migrate_pg_index.ReadObject_ctx(ctx, db, m)
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
