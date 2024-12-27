//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_postgres_migrate_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"context"
	"time"
	"github.com/ManyakRus/starter/log"
)

// UpdateManyFields - обновляет несколько полей в базе данных, по ИД
func (crud Crud_GRPC) UpdateManyFields(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint, MassNeedUpdateFields []string) error {
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
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_UpdateManyFields(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_UpdateManyFields(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Condeferrable - изменяет 1 поле Condeferrable в базе данных
func (crud Crud_GRPC) Update_Condeferrable(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Condeferrable
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Condeferrable(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Condeferrable(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Condeferred - изменяет 1 поле Condeferred в базе данных
func (crud Crud_GRPC) Update_Condeferred(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Condeferred
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Condeferred(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Condeferred(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conexclop - изменяет 1 поле Conexclop в базе данных
func (crud Crud_GRPC) Update_Conexclop(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Conexclop
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conexclop(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conexclop(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Confdeltype - изменяет 1 поле Confdeltype в базе данных
func (crud Crud_GRPC) Update_Confdeltype(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Confdeltype
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Confdeltype(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Confdeltype(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conffeqop - изменяет 1 поле Conffeqop в базе данных
func (crud Crud_GRPC) Update_Conffeqop(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Conffeqop
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conffeqop(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conffeqop(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Confkey - изменяет 1 поле Confkey в базе данных
func (crud Crud_GRPC) Update_Confkey(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Confkey
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Confkey(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Confkey(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Confmatchtype - изменяет 1 поле Confmatchtype в базе данных
func (crud Crud_GRPC) Update_Confmatchtype(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Confmatchtype
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Confmatchtype(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Confmatchtype(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Confrelid - изменяет 1 поле Confrelid в базе данных
func (crud Crud_GRPC) Update_Confrelid(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Confrelid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Confrelid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Confrelid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Confupdtype - изменяет 1 поле Confupdtype в базе данных
func (crud Crud_GRPC) Update_Confupdtype(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Confupdtype
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Confupdtype(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Confupdtype(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conindid - изменяет 1 поле Conindid в базе данных
func (crud Crud_GRPC) Update_Conindid(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Conindid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conindid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conindid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Coninhcount - изменяет 1 поле Coninhcount в базе данных
func (crud Crud_GRPC) Update_Coninhcount(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Coninhcount
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Coninhcount(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Coninhcount(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conislocal - изменяет 1 поле Conislocal в базе данных
func (crud Crud_GRPC) Update_Conislocal(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Conislocal
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conislocal(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conislocal(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conkey - изменяет 1 поле Conkey в базе данных
func (crud Crud_GRPC) Update_Conkey(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Conkey
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conkey(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conkey(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conname - изменяет 1 поле Conname в базе данных
func (crud Crud_GRPC) Update_Conname(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Conname
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conname(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conname(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Connamespace - изменяет 1 поле Connamespace в базе данных
func (crud Crud_GRPC) Update_Connamespace(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Connamespace
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Connamespace(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Connamespace(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Connoinherit - изменяет 1 поле Connoinherit в базе данных
func (crud Crud_GRPC) Update_Connoinherit(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Connoinherit
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Connoinherit(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Connoinherit(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conparentid - изменяет 1 поле Conparentid в базе данных
func (crud Crud_GRPC) Update_Conparentid(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Conparentid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conparentid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conparentid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conpfeqop - изменяет 1 поле Conpfeqop в базе данных
func (crud Crud_GRPC) Update_Conpfeqop(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Conpfeqop
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conpfeqop(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conpfeqop(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conppeqop - изменяет 1 поле Conppeqop в базе данных
func (crud Crud_GRPC) Update_Conppeqop(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Conppeqop
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conppeqop(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conppeqop(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Conrelid - изменяет 1 поле Conrelid в базе данных
func (crud Crud_GRPC) Update_Conrelid(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Conrelid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Conrelid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Conrelid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Contype - изменяет 1 поле Contype в базе данных
func (crud Crud_GRPC) Update_Contype(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Contype
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Contype(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Contype(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Contypid - изменяет 1 поле Contypid в базе данных
func (crud Crud_GRPC) Update_Contypid(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Contypid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Contypid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Contypid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Convalidated - изменяет 1 поле Convalidated в базе данных
func (crud Crud_GRPC) Update_Convalidated(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Oid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Convalidated
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Convalidated(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Convalidated(ctx, Request)
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
func (crud Crud_GRPC) Update_Oid(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
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
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_Oid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_Oid(ctx, Request)
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
func (crud Crud_GRPC) Update_VersionID(m *postgres_migrate_pg_constraint.PostgresMigratePgConstraint) error {
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
		//_, err = nrpc_client.Client.PostgresMigratePgConstraint_Update_VersionID(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgConstraint_Update_VersionID(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}
