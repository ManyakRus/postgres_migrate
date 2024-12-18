package crud_starter_postgres_migrate_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_namespace"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_postgres_migrate_pg_namespace.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if postgres_migrate_pg_namespace.Crud_manual_PostgresMigratePgNamespace != crud {
		t.Errorf("Expected postgres_migrate_pg_namespace.Crud_manual_PostgresMigratePgNamespace to be set to crud, but got %+v", postgres_migrate_pg_namespace.Crud_manual_PostgresMigratePgNamespace)
	}
}
