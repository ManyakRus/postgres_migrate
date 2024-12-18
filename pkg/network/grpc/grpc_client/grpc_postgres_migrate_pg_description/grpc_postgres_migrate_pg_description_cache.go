//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_postgres_migrate_pg_description

import (
	"context"
	"encoding/json"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
	"github.com/ManyakRus/starter/log"
	"time"
)

// ReadFromCache - возвращает модель из БД
func (crud Crud_GRPC) ReadFromCache(Classoid int64, Objoid int64, Objsubid int32, VersionID int64) (postgres_migrate_pg_description.PostgresMigratePgDescription, error) {
	Otvet := postgres_migrate_pg_description.PostgresMigratePgDescription{}
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = Classoid
	Request.Int64_2 = Objoid
	Request.Int32_1 = Objsubid
	Request.Int64_3 = VersionID

	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	var Response *grpc_proto.Response
	if grpc_nrpc.NeedNRPC == true {
		//Response, err = nrpc_client.Client.PostgresMigratePgDescription_ReadFromCache(Request)
	} else {
		Response, err = grpc_client_func.Client.PostgresMigratePgDescription_ReadFromCache(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return Otvet, err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), &Otvet)
	if err != nil {
		return Otvet, err
	}

	return Otvet, err
}
