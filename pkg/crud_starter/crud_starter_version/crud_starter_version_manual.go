package crud_starter_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
)

// SetCrudManualInterface - заполняет интерфейс crud: DB, GRPC, NRPC
func SetCrudManualInterface(crud version.ICrud_manual_Version) {
	version.Crud_manual_Version = crud

	return
}
