//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_postgres_migrate_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_attribute"
	"github.com/ManyakRus/starter/postgres_gorm"
)

// PostgresMigratePgAttribute_ReadObject - читает и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgAttribute_ReadObject(ctx context.Context, Request *grpc_proto.Request_String_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := object_postgres_migrate_pg_attribute.ObjectPostgresMigratePgAttribute{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(object_postgres_migrate_pg_attribute.ObjectPostgresMigratePgAttribute{})
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Attname := Request.String_1
	Attrelid := Request.Int64_1
	VersionID := Request.Int64_2
	m := &object_postgres_migrate_pg_attribute.ObjectPostgresMigratePgAttribute{}
	m.Attname = Attname
	m.Attrelid = Attrelid
	m.VersionID = VersionID

	err = crud_object_postgres_migrate_pg_attribute.ReadObject_ctx(ctx, db, m)
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
