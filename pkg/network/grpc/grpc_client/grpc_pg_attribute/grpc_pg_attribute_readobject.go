//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/objects/object_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"context"
	"encoding/json"
	"github.com/ManyakRus/starter/log"
	"time"
)

// ReadObject - возвращает модель из БД, и все модели имеющие foreign key
func (crud Crud_GRPC) ReadObject(m *object_pg_attribute.ObjectPgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = m.GetStructVersion()

	Request := &grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	var Response *grpc_proto.Response
	if grpc_nrpc.NeedNRPC == true {
		//Response, err = nrpc_client.Client.PgAttribute_ReadObject(Request)
	} else {
		Response, err = grpc_client_func.Client.PgAttribute_ReadObject(ctx, Request)
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
