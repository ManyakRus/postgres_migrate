//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package grpc_postgres_migrate_pg_sequence

import (
	"context"
	"github.com/ManyakRus/postgres_migrate/api/grpc_proto"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client_func"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc_nrpc"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"github.com/ManyakRus/starter/log"
	"time"
)

// UpdateManyFields - обновляет несколько полей в базе данных, по ИД
func (crud Crud_GRPC) UpdateManyFields(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence, MassNeedUpdateFields []string) error {
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
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_UpdateManyFields(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_UpdateManyFields(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Seqcache - изменяет 1 поле Seqcache в базе данных
func (crud Crud_GRPC) Update_Seqcache(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Seqrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Seqcache
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_Update_Seqcache(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_Update_Seqcache(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Seqcycle - изменяет 1 поле Seqcycle в базе данных
func (crud Crud_GRPC) Update_Seqcycle(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Bool{}
	Request.Int64_1 = m.Seqrelid
	Request.Int64_2 = m.VersionID

	Request.Bool_1 = m.Seqcycle
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_Update_Seqcycle(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_Update_Seqcycle(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Seqincrement - изменяет 1 поле Seqincrement в базе данных
func (crud Crud_GRPC) Update_Seqincrement(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Seqrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Seqincrement
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_Update_Seqincrement(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_Update_Seqincrement(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Seqmax - изменяет 1 поле Seqmax в базе данных
func (crud Crud_GRPC) Update_Seqmax(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Seqrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Seqmax
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_Update_Seqmax(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_Update_Seqmax(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Seqmin - изменяет 1 поле Seqmin в базе данных
func (crud Crud_GRPC) Update_Seqmin(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Seqrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Seqmin
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_Update_Seqmin(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_Update_Seqmin(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Seqrelid - изменяет 1 поле Seqrelid в базе данных
func (crud Crud_GRPC) Update_Seqrelid(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = m.Seqrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_1 = m.Seqrelid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_Update_Seqrelid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_Update_Seqrelid(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Seqstart - изменяет 1 поле Seqstart в базе данных
func (crud Crud_GRPC) Update_Seqstart(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Seqrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Seqstart
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_Update_Seqstart(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_Update_Seqstart(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}

// Update_Seqtypid - изменяет 1 поле Seqtypid в базе данных
func (crud Crud_GRPC) Update_Seqtypid(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64_Int64{}
	Request.Int64_1 = m.Seqrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_3 = m.Seqtypid
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_Update_Seqtypid(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_Update_Seqtypid(ctx, Request)
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
func (crud Crud_GRPC) Update_VersionID(m *postgres_migrate_pg_sequence.PostgresMigratePgSequence) error {
	var err error

	// подключение
	grpc_client_func.Func_Connect_GRPC_NRPC.Connect_GRPC_NRPC()

	// подготовка запроса
	var versionModel = crud.GetVersionModel()

	Request := &grpc_proto.Request_Int64_Int64{}
	Request.Int64_1 = m.Seqrelid
	Request.Int64_2 = m.VersionID

	Request.Int64_2 = m.VersionID
	Request.VersionModel = versionModel

	ctxMain := context.Background()
	ctx, ctxCancelFunc := context.WithTimeout(ctxMain, time.Second*time.Duration(grpc_constants.GetTimeoutSeconds()))
	defer ctxCancelFunc()

	// запрос
	if grpc_nrpc.NeedNRPC == true {
		//_, err = nrpc_client.Client.PostgresMigratePgSequence_Update_VersionID(Request)
	} else {
		_, err = grpc_client_func.Client.PostgresMigratePgSequence_Update_VersionID(ctx, Request)
	}
	if err != nil {
		if grpc_client_func.IsErrorModelVersion(err) == true {
			log.Panic(err)
		}
		return err
	}

	return err
}
