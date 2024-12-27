//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_postgres_migrate_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"context"
	"time"
	"github.com/ManyakRus/starter/log"
)

// UpdateManyFields - обновляет несколько полей в базе данных, по ИД
func (crud Crud_GRPC) UpdateManyFields(m *postgres_migrate_pg_class.PostgresMigratePgClass, MassNeedUpdateFields []string) error {
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
		//_, err = nrpc_client.Client.PostgresMigratePgClass_UpdateManyFields(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_UpdateManyFields(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Oid - изменяет 1 поле Oid в базе данных
func (crud Crud_GRPC) Update_Oid(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_1 = m.Oid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Oid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Oid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relallvisible - изменяет 1 поле Relallvisible в базе данных
func (crud Crud_GRPC) Update_Relallvisible(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Relallvisible
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relallvisible(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relallvisible(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relam - изменяет 1 поле Relam в базе данных
func (crud Crud_GRPC) Update_Relam(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Relam
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relam(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relam(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relchecks - изменяет 1 поле Relchecks в базе данных
func (crud Crud_GRPC) Update_Relchecks(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Relchecks
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relchecks(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relchecks(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relfilenode - изменяет 1 поле Relfilenode в базе данных
func (crud Crud_GRPC) Update_Relfilenode(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Relfilenode
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relfilenode(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relfilenode(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relforcerowsecurity - изменяет 1 поле Relforcerowsecurity в базе данных
func (crud Crud_GRPC) Update_Relforcerowsecurity(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Relforcerowsecurity
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relforcerowsecurity(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relforcerowsecurity(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relfrozenxid - изменяет 1 поле Relfrozenxid в базе данных
func (crud Crud_GRPC) Update_Relfrozenxid(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Relfrozenxid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relfrozenxid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relfrozenxid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relhasindex - изменяет 1 поле Relhasindex в базе данных
func (crud Crud_GRPC) Update_Relhasindex(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Relhasindex
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relhasindex(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relhasindex(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relhasrules - изменяет 1 поле Relhasrules в базе данных
func (crud Crud_GRPC) Update_Relhasrules(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Relhasrules
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relhasrules(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relhasrules(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relhassubclass - изменяет 1 поле Relhassubclass в базе данных
func (crud Crud_GRPC) Update_Relhassubclass(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Relhassubclass
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relhassubclass(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relhassubclass(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relhastriggers - изменяет 1 поле Relhastriggers в базе данных
func (crud Crud_GRPC) Update_Relhastriggers(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Relhastriggers
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relhastriggers(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relhastriggers(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relispartition - изменяет 1 поле Relispartition в базе данных
func (crud Crud_GRPC) Update_Relispartition(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Relispartition
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relispartition(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relispartition(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relispopulated - изменяет 1 поле Relispopulated в базе данных
func (crud Crud_GRPC) Update_Relispopulated(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Relispopulated
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relispopulated(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relispopulated(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relisshared - изменяет 1 поле Relisshared в базе данных
func (crud Crud_GRPC) Update_Relisshared(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Relisshared
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relisshared(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relisshared(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relkind - изменяет 1 поле Relkind в базе данных
func (crud Crud_GRPC) Update_Relkind(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Relkind
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relkind(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relkind(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relminmxid - изменяет 1 поле Relminmxid в базе данных
func (crud Crud_GRPC) Update_Relminmxid(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Relminmxid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relminmxid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relminmxid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relname - изменяет 1 поле Relname в базе данных
func (crud Crud_GRPC) Update_Relname(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Relname
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relname(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relname(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relnamespace - изменяет 1 поле Relnamespace в базе данных
func (crud Crud_GRPC) Update_Relnamespace(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Relnamespace
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relnamespace(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relnamespace(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relnatts - изменяет 1 поле Relnatts в базе данных
func (crud Crud_GRPC) Update_Relnatts(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Relnatts
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relnatts(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relnatts(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Reloftype - изменяет 1 поле Reloftype в базе данных
func (crud Crud_GRPC) Update_Reloftype(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Reloftype
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Reloftype(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Reloftype(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relowner - изменяет 1 поле Relowner в базе данных
func (crud Crud_GRPC) Update_Relowner(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Relowner
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relowner(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relowner(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relpages - изменяет 1 поле Relpages в базе данных
func (crud Crud_GRPC) Update_Relpages(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Relpages
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relpages(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relpages(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relpersistence - изменяет 1 поле Relpersistence в базе данных
func (crud Crud_GRPC) Update_Relpersistence(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Relpersistence
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relpersistence(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relpersistence(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relreplident - изменяет 1 поле Relreplident в базе данных
func (crud Crud_GRPC) Update_Relreplident(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Relreplident
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relreplident(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relreplident(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relrewrite - изменяет 1 поле Relrewrite в базе данных
func (crud Crud_GRPC) Update_Relrewrite(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Relrewrite
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relrewrite(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relrewrite(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Relrowsecurity - изменяет 1 поле Relrowsecurity в базе данных
func (crud Crud_GRPC) Update_Relrowsecurity(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Relrowsecurity
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Relrowsecurity(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Relrowsecurity(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Reltablespace - изменяет 1 поле Reltablespace в базе данных
func (crud Crud_GRPC) Update_Reltablespace(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Reltablespace
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Reltablespace(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Reltablespace(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Reltoastrelid - изменяет 1 поле Reltoastrelid в базе данных
func (crud Crud_GRPC) Update_Reltoastrelid(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Reltoastrelid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Reltoastrelid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Reltoastrelid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Reltuples - изменяет 1 поле Reltuples в базе данных
func (crud Crud_GRPC) Update_Reltuples(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Float32{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Float32_1 = m.Reltuples
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Reltuples(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Reltuples(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Reltype - изменяет 1 поле Reltype в базе данных
func (crud Crud_GRPC) Update_Reltype(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Reltype
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_Reltype(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_Reltype(ctx, Request)
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
func (crud Crud_GRPC) Update_VersionID(m *postgres_migrate_pg_class.PostgresMigratePgClass) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_2 = m.VersionID
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgClass_Update_VersionID(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgClass_Update_VersionID(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}
