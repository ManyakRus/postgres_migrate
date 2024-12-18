package crud_starter_postgres_migrate_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_class"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_postgres_migrate_pg_class.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if postgres_migrate_pg_class.Crud_manual_PostgresMigratePgClass != crud {
		t.Errorf("Expected postgres_migrate_pg_class.Crud_manual_PostgresMigratePgClass to be set to crud, but got %+v", postgres_migrate_pg_class.Crud_manual_PostgresMigratePgClass)
	}
}
