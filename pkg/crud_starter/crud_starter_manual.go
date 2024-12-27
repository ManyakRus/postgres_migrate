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

// initCrudTransport_manual_DB - заполняет объекты crud для работы с БД напрямую
func initCrudTransport_manual_DB() {
	crud_starter_postgres_migrate_pg_attribute.SetCrudManualInterface(crud_postgres_migrate_pg_attribute.Crud_DB{})
	crud_starter_postgres_migrate_pg_class.SetCrudManualInterface(crud_postgres_migrate_pg_class.Crud_DB{})
	crud_starter_postgres_migrate_pg_constraint.SetCrudManualInterface(crud_postgres_migrate_pg_constraint.Crud_DB{})
	crud_starter_postgres_migrate_pg_description.SetCrudManualInterface(crud_postgres_migrate_pg_description.Crud_DB{})
	crud_starter_postgres_migrate_pg_index.SetCrudManualInterface(crud_postgres_migrate_pg_index.Crud_DB{})
	crud_starter_postgres_migrate_pg_namespace.SetCrudManualInterface(crud_postgres_migrate_pg_namespace.Crud_DB{})
	crud_starter_postgres_migrate_version.SetCrudManualInterface(crud_postgres_migrate_version.Crud_DB{})
}

// initCrudTransport_manual_GRPC - заполняет объекты crud для работы с БД через протокол GRPC
func initCrudTransport_manual_GRPC() {
	crud_starter_postgres_migrate_pg_attribute.SetCrudManualInterface(grpc_postgres_migrate_pg_attribute.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_class.SetCrudManualInterface(grpc_postgres_migrate_pg_class.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_constraint.SetCrudManualInterface(grpc_postgres_migrate_pg_constraint.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_description.SetCrudManualInterface(grpc_postgres_migrate_pg_description.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_index.SetCrudManualInterface(grpc_postgres_migrate_pg_index.Crud_GRPC{})
	crud_starter_postgres_migrate_pg_namespace.SetCrudManualInterface(grpc_postgres_migrate_pg_namespace.Crud_GRPC{})
	crud_starter_postgres_migrate_version.SetCrudManualInterface(grpc_postgres_migrate_version.Crud_GRPC{})
}