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

// initCrudTransport_manual_DB - заполняет объекты crud для работы с БД напрямую
func initCrudTransport_manual_DB() {
	crud_starter_pg_attribute.SetCrudManualInterface(crud_pg_attribute.Crud_DB{})
	crud_starter_pg_class.SetCrudManualInterface(crud_pg_class.Crud_DB{})
	crud_starter_pg_constraint.SetCrudManualInterface(crud_pg_constraint.Crud_DB{})
	crud_starter_pg_description.SetCrudManualInterface(crud_pg_description.Crud_DB{})
	crud_starter_pg_index.SetCrudManualInterface(crud_pg_index.Crud_DB{})
	crud_starter_pg_namespace.SetCrudManualInterface(crud_pg_namespace.Crud_DB{})
	crud_starter_version.SetCrudManualInterface(crud_version.Crud_DB{})
}

// initCrudTransport_manual_GRPC - заполняет объекты crud для работы с БД через протокол GRPC
func initCrudTransport_manual_GRPC() {
	crud_starter_pg_attribute.SetCrudManualInterface(grpc_pg_attribute.Crud_GRPC{})
	crud_starter_pg_class.SetCrudManualInterface(grpc_pg_class.Crud_GRPC{})
	crud_starter_pg_constraint.SetCrudManualInterface(grpc_pg_constraint.Crud_GRPC{})
	crud_starter_pg_description.SetCrudManualInterface(grpc_pg_description.Crud_GRPC{})
	crud_starter_pg_index.SetCrudManualInterface(grpc_pg_index.Crud_GRPC{})
	crud_starter_pg_namespace.SetCrudManualInterface(grpc_pg_namespace.Crud_GRPC{})
	crud_starter_version.SetCrudManualInterface(grpc_version.Crud_GRPC{})
}