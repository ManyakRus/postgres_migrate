//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"context"
	"time"
	"github.com/ManyakRus/starter/log"
)

// UpdateManyFields - обновляет несколько полей в базе данных, по ИД
func (crud Crud_GRPC) UpdateManyFields(m *postgres_migrate_pg_index.PostgresMigratePgIndex, MassNeedUpdateFields []string) error {
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
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_UpdateManyFields(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_UpdateManyFields(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indcheckxmin - изменяет 1 поле Indcheckxmin в базе данных
func (crud Crud_GRPC) Update_Indcheckxmin(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indcheckxmin
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indcheckxmin(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indcheckxmin(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indclass - изменяет 1 поле Indclass в базе данных
func (crud Crud_GRPC) Update_Indclass(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Indclass
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indclass(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indclass(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indcollation - изменяет 1 поле Indcollation в базе данных
func (crud Crud_GRPC) Update_Indcollation(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Indcollation
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indcollation(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indcollation(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indexprs - изменяет 1 поле Indexprs в базе данных
func (crud Crud_GRPC) Update_Indexprs(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Indexprs
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indexprs(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indexprs(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indexrelid - изменяет 1 поле Indexrelid в базе данных
func (crud Crud_GRPC) Update_Indexrelid(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_1 = m.Indexrelid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indexrelid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indexrelid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indimmediate - изменяет 1 поле Indimmediate в базе данных
func (crud Crud_GRPC) Update_Indimmediate(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indimmediate
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indimmediate(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indimmediate(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indisclustered - изменяет 1 поле Indisclustered в базе данных
func (crud Crud_GRPC) Update_Indisclustered(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indisclustered
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indisclustered(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indisclustered(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indisexclusion - изменяет 1 поле Indisexclusion в базе данных
func (crud Crud_GRPC) Update_Indisexclusion(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indisexclusion
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indisexclusion(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indisexclusion(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indislive - изменяет 1 поле Indislive в базе данных
func (crud Crud_GRPC) Update_Indislive(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indislive
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indislive(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indislive(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indisprimary - изменяет 1 поле Indisprimary в базе данных
func (crud Crud_GRPC) Update_Indisprimary(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indisprimary
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indisprimary(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indisprimary(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indisready - изменяет 1 поле Indisready в базе данных
func (crud Crud_GRPC) Update_Indisready(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indisready
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indisready(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indisready(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indisreplident - изменяет 1 поле Indisreplident в базе данных
func (crud Crud_GRPC) Update_Indisreplident(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indisreplident
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indisreplident(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indisreplident(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indisunique - изменяет 1 поле Indisunique в базе данных
func (crud Crud_GRPC) Update_Indisunique(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indisunique
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indisunique(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indisunique(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indisvalid - изменяет 1 поле Indisvalid в базе данных
func (crud Crud_GRPC) Update_Indisvalid(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Indisvalid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indisvalid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indisvalid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indkey - изменяет 1 поле Indkey в базе данных
func (crud Crud_GRPC) Update_Indkey(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Indkey
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indkey(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indkey(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indnatts - изменяет 1 поле Indnatts в базе данных
func (crud Crud_GRPC) Update_Indnatts(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Indnatts
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indnatts(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indnatts(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indnkeyatts - изменяет 1 поле Indnkeyatts в базе данных
func (crud Crud_GRPC) Update_Indnkeyatts(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int32{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Int32_1 = m.Indnkeyatts
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indnkeyatts(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indnkeyatts(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indoption - изменяет 1 поле Indoption в базе данных
func (crud Crud_GRPC) Update_Indoption(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Indoption
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indoption(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indoption(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indpred - изменяет 1 поле Indpred в базе данных
func (crud Crud_GRPC) Update_Indpred(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_String{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.String_1 = m.Indpred
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indpred(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indpred(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Indrelid - изменяет 1 поле Indrelid в базе данных
func (crud Crud_GRPC) Update_Indrelid(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Indrelid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_Indrelid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_Indrelid(ctx, Request)
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
func (crud Crud_GRPC) Update_VersionID(m *postgres_migrate_pg_index.PostgresMigratePgIndex) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = m.Indexrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_2 = m.VersionID
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgIndex_Update_VersionID(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgIndex_Update_VersionID(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}
