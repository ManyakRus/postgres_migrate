//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_description"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"context"
	"encoding/json"
	"github.com/ManyakRus/starter/log"
	"time"
)

// ReadObject - возвращает модель из БД, и все модели имеющие foreign key
func (crud Crud_GRPC) ReadObject(m *object_pg_description.ObjectPgDescription) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = m.GetStructVersion()

	Request := &grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = m.Classoid
	Request.Int64_2 = m.Objoid
	Request.Int32_1 = m.Objsubid
	Request.Int64_3 = m.VersionID

	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	var Response *grpc_proto.Response
	if grpc_nrpc.NeedNRPC == true {
		//Response, err = nrpc_client.Client.PgDescription_ReadObject(Request)
	} else {
		Response, err = grpc_client_func.Client.PgDescription_ReadObject(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	// ответ
	sModel := Response.ModelString
	err = json.Unmarshal([]byte(sModel), m)
	if err != nil {
		return err
	}

	return err
}
