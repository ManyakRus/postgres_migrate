//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter

import (
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_version"

	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_version"

	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_version"
)

// InitCrudTransport_DB - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_DB() { 
	initCrudTransport_manual_DB()
	InitCrudTransport_ReadObject_DB()

	crud_starter_pg_attribute.SetCrudInterface(crud_pg_attribute.Crud_DB{})
	crud_starter_pg_class.SetCrudInterface(crud_pg_class.Crud_DB{})
	crud_starter_pg_constraint.SetCrudInterface(crud_pg_constraint.Crud_DB{})
	crud_starter_pg_description.SetCrudInterface(crud_pg_description.Crud_DB{})
	crud_starter_pg_index.SetCrudInterface(crud_pg_index.Crud_DB{})
	crud_starter_pg_namespace.SetCrudInterface(crud_pg_namespace.Crud_DB{})
	crud_starter_version.SetCrudInterface(crud_version.Crud_DB{})
}

// InitCrudTransport_GRPC - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_GRPC() {
	initCrudTransport_manual_GRPC()
	InitCrudTransport_ReadObject_GRPC()

	crud_starter_pg_attribute.SetCrudInterface(grpc_pg_attribute.Crud_GRPC{})
	crud_starter_pg_class.SetCrudInterface(grpc_pg_class.Crud_GRPC{})
	crud_starter_pg_constraint.SetCrudInterface(grpc_pg_constraint.Crud_GRPC{})
	crud_starter_pg_description.SetCrudInterface(grpc_pg_description.Crud_GRPC{})
	crud_starter_pg_index.SetCrudInterface(grpc_pg_index.Crud_GRPC{})
	crud_starter_pg_namespace.SetCrudInterface(grpc_pg_namespace.Crud_GRPC{})
	crud_starter_version.SetCrudInterface(grpc_version.Crud_GRPC{})
}