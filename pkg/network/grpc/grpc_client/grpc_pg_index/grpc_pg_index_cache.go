//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_index"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"context"
	"encoding/json"
	"github.com/ManyakRus/starter/log"
	"time"
)

// ReadFromCache - возвращает модель из БД
func (crud Crud_GRPC) ReadFromCache(Indexrelid int64, Indrelid int64, VersionID int64) (pg_index.PgIndex, error) {
	Otvet := pg_index.PgIndex{}
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = Indexrelid
	Request.Int64_2 = Indrelid
	Request.Int64_3 = VersionID

	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	var Response *grpc_proto.Response
	if grpc_nrpc.NeedNRPC == true {
		//Response, err = nrpc_client.Client.PgIndex_ReadFromCache(Request)
	} else {
		Response, err = grpc_client_func.Client.PgIndex_ReadFromCache(ctx, Request)
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
