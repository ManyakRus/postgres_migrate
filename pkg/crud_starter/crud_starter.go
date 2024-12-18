//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter

import (
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_postgres_migrate_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_postgres_migrate_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_postgres_migrate_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/crud_starter/crud_starter_postgres_migrate_version"

	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_version"

	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_postgres_migrate_version"
)

// InitCrudTransport_DB - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_DB() {
	initCrudTransport_manual_DB()
	InitCrudTransport_ReadObject_DB()

	crud_starter_postgres_migrate_pg_attribute.SetCrudInterface(crud_postgres_migrate_pg_attribute.Crud_DB{})
	crud_starter_postgres_migrate_pg_class.SetCrudInterface(crud_postgres_migrate_pg_class.Crud_DB{})
	crud_starter_postgres_migrate_pg_constraint.SetCrudInterface(crud_postgres_migrate_pg_constraint.Crud_DB{})
	crud_starter_postgres_migrate_pg_description.SetCrudInterface(crud_postgres_migrate_pg_description.Crud_DB{})
	crud_starter_postgres_migrate_pg_index.SetCrudInterface(crud_postgres_migrate_pg_index.Crud_DB{})
	crud_starter_postgres_migrate_pg_namespace.SetCrudInterface(crud_postgres_migrate_pg_namespace.Crud_DB{})
	crud_starter_postgres_migrate_version.SetCrudInterface(crud_postgres_migrate_version.Crud_DB{})
}

// InitCrudTransport_GRPC - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_GRPC() {
	initCrudTransport_manual_GRPC()
	InitCrudTransport_ReadObject_GRPC()

	crud_starter_postgres_migrate_pg_attribute.SetCrudInterface(grpc_postgres_migrate_pg_attribute.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_class.SetCrudInterface(grpc_postgres_migrate_pg_class.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_constraint.SetCrudInterface(grpc_postgres_migrate_pg_constraint.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_description.SetCrudInterface(grpc_postgres_migrate_pg_description.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_index.SetCrudInterface(grpc_postgres_migrate_pg_index.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_namespace.SetCrudInterface(grpc_postgres_migrate_pg_namespace.Crud_GRPC{})
	crud_starter_postgres_migrate_version.SetCrudInterface(grpc_postgres_migrate_version.Crud_GRPC{})
}
