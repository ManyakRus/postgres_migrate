//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"context"
	"time"
	"github.com/ManyakRus/starter/log"
)

// UpdateManyFields - обновляет несколько полей в базе данных, по ИД
func (crud Crud_GRPC) UpdateManyFields(m *postgres_migrate_version.PostgresMigrateVersion, MassNeedUpdateFields []string) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var VersionModel = crud.GetVersionModel()

	ModelString, err := m.GetJSON()
	if err != nil {
		return err
	}
	Request := &grpc_proto.Request_Model_MassString{}
	Request.VersionModel = VersionModel
	Request.ModelString = ModelString
	Request.MassNames = MassNeedUpdateFields

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigrateVersion_UpdateManyFields(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigrateVersion_UpdateManyFields(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Description - изменяет 1 поле Description в базе данных
func (crud Crud_GRPC) Update_Description(m *postgres_migrate_version.PostgresMigrateVersion) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_String{}
	Request.Int64_1 = m.ID

	Request.String_1 = m.Description
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigrateVersion_Update_Description(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigrateVersion_Update_Description(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Name - изменяет 1 поле Name в базе данных
func (crud Crud_GRPC) Update_Name(m *postgres_migrate_version.PostgresMigrateVersion) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_String{}
	Request.Int64_1 = m.ID

	Request.String_1 = m.Name
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigrateVersion_Update_Name(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigrateVersion_Update_Name(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}
