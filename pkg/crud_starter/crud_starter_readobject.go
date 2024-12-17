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

	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud_objects/crud_object_version"

	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_client/grpc_version"
)

// InitCrudTransport_ReadObject_DB - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_ReadObject_DB() { 

	crud_starter_pg_attribute.SetCrudReadObjectInterface(crud_object_pg_attribute.Crud_DB{})
	crud_starter_pg_class.SetCrudReadObjectInterface(crud_object_pg_class.Crud_DB{})
	crud_starter_pg_constraint.SetCrudReadObjectInterface(crud_object_pg_constraint.Crud_DB{})
	crud_starter_pg_description.SetCrudReadObjectInterface(crud_object_pg_description.Crud_DB{})
	crud_starter_pg_index.SetCrudReadObjectInterface(crud_object_pg_index.Crud_DB{})
	crud_starter_pg_namespace.SetCrudReadObjectInterface(crud_object_pg_namespace.Crud_DB{})
	crud_starter_version.SetCrudReadObjectInterface(crud_object_version.Crud_DB{})
}

// InitCrudTransport_ReadObject_GRPC - заполняет объекты crud для работы с БД напрямую
func InitCrudTransport_ReadObject_GRPC() { 

	crud_starter_pg_attribute.SetCrudReadObjectInterface(grpc_pg_attribute.Crud_GRPC{})
	crud_starter_pg_class.SetCrudReadObjectInterface(grpc_pg_class.Crud_GRPC{})
	crud_starter_pg_constraint.SetCrudReadObjectInterface(grpc_pg_constraint.Crud_GRPC{})
	crud_starter_pg_description.SetCrudReadObjectInterface(grpc_pg_description.Crud_GRPC{})
	crud_starter_pg_index.SetCrudReadObjectInterface(grpc_pg_index.Crud_GRPC{})
	crud_starter_pg_namespace.SetCrudReadObjectInterface(grpc_pg_namespace.Crud_GRPC{})
	crud_starter_version.SetCrudReadObjectInterface(grpc_version.Crud_GRPC{})
}