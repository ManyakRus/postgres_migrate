//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package server_grpc

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/postgres_gorm"
	"time"
)

// PostgresMigratePgIndex_ReadFromCache - читает и возвращает модель из кеша или БД
func (s *ServerGRPC) PostgresMigratePgIndex_ReadFromCache(ctx context.Context, Request *grpc_proto.Request_Int64_Int64) (*grpc_proto.Response, error) {
	var Otvet grpc_proto.Response
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
	Model := postgres_migrate_pg_index.PostgresMigratePgIndex{}
	Model, err = crud_postgres_migrate_pg_index.ReadFromCache_ctx(ctx, db, Indexrelid, VersionID)
	if err != nil {
		return &Otvet, err
	}

	//заполяем ответ
	ModelString, err := Model.GetJSON()
	if err != nil {
		return &Otvet, err
	}
	Otvet.ModelString = ModelString

	return &Otvet, err
}
