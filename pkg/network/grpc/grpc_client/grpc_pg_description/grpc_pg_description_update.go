//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"context"
	"time"
	"github.com/ManyakRus/starter/log"
)

// UpdateManyFields - обновляет несколько полей в базе данных, по ИД
func (crud Crud_GRPC) UpdateManyFields(m *pg_description.PgDescription, MassNeedUpdateFields []string) error {
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
		//_, err = nrpc_client.Client.PgDescription_UpdateManyFields(Request)
	} else {
		_, err = grpc_client_func.Client.PgDescription_UpdateManyFields(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Classoid - изменяет 1 поле Classoid в базе данных
func (crud Crud_GRPC) Update_Classoid(m *pg_description.PgDescription) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = m.Classoid
	Request.Int64_2 = m.Objoid
	Request.Int32_1 = m.Objsubid
	Request.Int64_3 = m.VersionID

	Request.Int64_1 = m.Classoid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PgDescription_Update_Classoid(Request)
	} else {
		_, err = grpc_client_func.Client.PgDescription_Update_Classoid(ctx, Request)
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
func (crud Crud_GRPC) Update_Description(m *pg_description.PgDescription) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32_Int64_String{}
	Request.Int64_1 = m.Classoid
	Request.Int64_2 = m.Objoid
	Request.Int32_1 = m.Objsubid
	Request.Int64_3 = m.VersionID

	Request.String_1 = m.Description
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PgDescription_Update_Description(Request)
	} else {
		_, err = grpc_client_func.Client.PgDescription_Update_Description(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Objoid - изменяет 1 поле Objoid в базе данных
func (crud Crud_GRPC) Update_Objoid(m *pg_description.PgDescription) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = m.Classoid
	Request.Int64_2 = m.Objoid
	Request.Int32_1 = m.Objsubid
	Request.Int64_3 = m.VersionID

	Request.Int64_2 = m.Objoid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PgDescription_Update_Objoid(Request)
	} else {
		_, err = grpc_client_func.Client.PgDescription_Update_Objoid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Objsubid - изменяет 1 поле Objsubid в базе данных
func (crud Crud_GRPC) Update_Objsubid(m *pg_description.PgDescription) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = m.Classoid
	Request.Int64_2 = m.Objoid
	Request.Int32_1 = m.Objsubid
	Request.Int64_3 = m.VersionID

	Request.Int32_1 = m.Objsubid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PgDescription_Update_Objsubid(Request)
	} else {
		_, err = grpc_client_func.Client.PgDescription_Update_Objsubid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_VersionID - изменяет 1 поле VersionID в базе данных
func (crud Crud_GRPC) Update_VersionID(m *pg_description.PgDescription) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32_Int64{}
	Request.Int64_1 = m.Classoid
	Request.Int64_2 = m.Objoid
	Request.Int32_1 = m.Objsubid
	Request.Int64_3 = m.VersionID

	Request.Int64_3 = m.VersionID
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PgDescription_Update_VersionID(Request)
	} else {
		_, err = grpc_client_func.Client.PgDescription_Update_VersionID(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}
