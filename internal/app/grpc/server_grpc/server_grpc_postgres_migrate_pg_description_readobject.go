//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_postgres_migrate_pg_description"
	"context"
	"github.com/ManyakRus/starter/postgres_gorm"
)

// PostgresMigratePgDescription_ReadObject - читает и возвращает модель из БД
func (s *ServerGRPC) PostgresMigratePgDescription_ReadObject(ctx context.Context, Request *grpc_proto.Request_Int64_Int64_Int32_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
	var err error

	//проверим контекст уже отменён
	if ctx.Err() != nil {
		err = ctx.Err()
		return &Otvet, err
	}

	//проверим совпадения версии модели
	VersionServer := object_postgres_migrate_pg_description.ObjectPostgresMigratePgDescription{}.GetStructVersion()
	VersionClient := Request.VersionModel
	if VersionServer != VersionClient {
		err = ErrorModelVersion(object_postgres_migrate_pg_description.ObjectPostgresMigratePgDescription{})
		return &Otvet, err
	}

	//запрос в БД
	db := postgres_gorm.GetConnection()
	Classoid := Request.Int64_1
	Objoid := Request.Int64_2
	Objsubid := Request.Int32_1
	VersionID := Request.Int64_3
	m := &object_postgres_migrate_pg_description.ObjectPostgresMigratePgDescription{}
	m.Classoid = Classoid
	m.Objoid = Objoid
	m.Objsubid = Objsubid
	m.VersionID = VersionID

	err = crud_object_postgres_migrate_pg_description.ReadObject_ctx(ctx, db, m)
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

