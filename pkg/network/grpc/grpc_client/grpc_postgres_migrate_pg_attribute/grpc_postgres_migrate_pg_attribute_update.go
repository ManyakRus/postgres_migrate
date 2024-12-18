//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_postgres_migrate_pg_attribute

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_attribute"
	"github.com/ManyakRus/starter/log"
	"time"
)

// UpdateManyFields - обновляет несколько полей в базе данных, по ИД
func (crud Crud_GRPC) UpdateManyFields(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute, MassNeedUpdateFields []string) error {
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
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_UpdateManyFields(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_UpdateManyFields(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attalign - изменяет 1 поле Attalign в базе данных
func (crud Crud_GRPC) Update_Attalign(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_String{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.String_2 = m.Attalign
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attalign(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attalign(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attbyval - изменяет 1 поле Attbyval в базе данных
func (crud Crud_GRPC) Update_Attbyval(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Bool{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Attbyval
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attbyval(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attbyval(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attcacheoff - изменяет 1 поле Attcacheoff в базе данных
func (crud Crud_GRPC) Update_Attcacheoff(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Int32{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Attcacheoff
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attcacheoff(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attcacheoff(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attcollation - изменяет 1 поле Attcollation в базе данных
func (crud Crud_GRPC) Update_Attcollation(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Int64{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Attcollation
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attcollation(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attcollation(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attgenerated - изменяет 1 поле Attgenerated в базе данных
func (crud Crud_GRPC) Update_Attgenerated(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_String{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.String_2 = m.Attgenerated
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attgenerated(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attgenerated(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Atthasdef - изменяет 1 поле Atthasdef в базе данных
func (crud Crud_GRPC) Update_Atthasdef(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Bool{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Atthasdef
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Atthasdef(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Atthasdef(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Atthasmissing - изменяет 1 поле Atthasmissing в базе данных
func (crud Crud_GRPC) Update_Atthasmissing(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Bool{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Atthasmissing
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Atthasmissing(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Atthasmissing(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attidentity - изменяет 1 поле Attidentity в базе данных
func (crud Crud_GRPC) Update_Attidentity(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_String{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.String_2 = m.Attidentity
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attidentity(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attidentity(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attinhcount - изменяет 1 поле Attinhcount в базе данных
func (crud Crud_GRPC) Update_Attinhcount(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Int32{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Attinhcount
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attinhcount(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attinhcount(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attisdropped - изменяет 1 поле Attisdropped в базе данных
func (crud Crud_GRPC) Update_Attisdropped(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Bool{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Attisdropped
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attisdropped(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attisdropped(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attislocal - изменяет 1 поле Attislocal в базе данных
func (crud Crud_GRPC) Update_Attislocal(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Bool{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Attislocal
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attislocal(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attislocal(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attlen - изменяет 1 поле Attlen в базе данных
func (crud Crud_GRPC) Update_Attlen(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Int32{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Attlen
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attlen(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attlen(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attname - изменяет 1 поле Attname в базе данных
func (crud Crud_GRPC) Update_Attname(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Attname
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attname(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attname(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attndims - изменяет 1 поле Attndims в базе данных
func (crud Crud_GRPC) Update_Attndims(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Int32{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Attndims
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attndims(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attndims(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attnotnull - изменяет 1 поле Attnotnull в базе данных
func (crud Crud_GRPC) Update_Attnotnull(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Bool{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Attnotnull
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attnotnull(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attnotnull(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attnum - изменяет 1 поле Attnum в базе данных
func (crud Crud_GRPC) Update_Attnum(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Int32{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Attnum
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attnum(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attnum(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attrelid - изменяет 1 поле Attrelid в базе данных
func (crud Crud_GRPC) Update_Attrelid(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_1 = m.Attrelid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attrelid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attrelid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attstattarget - изменяет 1 поле Attstattarget в базе данных
func (crud Crud_GRPC) Update_Attstattarget(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Int32{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Attstattarget
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attstattarget(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attstattarget(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Attstorage - изменяет 1 поле Attstorage в базе данных
func (crud Crud_GRPC) Update_Attstorage(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_String{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.String_2 = m.Attstorage
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Attstorage(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Attstorage(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Atttypid - изменяет 1 поле Atttypid в базе данных
func (crud Crud_GRPC) Update_Atttypid(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Int64{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Atttypid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Atttypid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Atttypid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Atttypmod - изменяет 1 поле Atttypmod в базе данных
func (crud Crud_GRPC) Update_Atttypmod(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64_Int32{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Atttypmod
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_Atttypmod(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_Atttypmod(ctx, Request)
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
func (crud Crud_GRPC) Update_VersionID(m *postgres_migrate_pg_attribute.PostgresMigratePgAttribute) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_String_Int64_Int64{}
	Request.String_1 = m.Attname
	Request.Int64_1 = m.Attrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_2 = m.VersionID
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgAttribute_Update_VersionID(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgAttribute_Update_VersionID(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}
