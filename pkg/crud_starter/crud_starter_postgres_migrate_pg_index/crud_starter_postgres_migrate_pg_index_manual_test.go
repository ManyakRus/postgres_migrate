package crud_starter_postgres_migrate_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_index"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_postgres_migrate_pg_index.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if postgres_migrate_pg_index.Crud_manual_PostgresMigratePgIndex != crud {
		t.Errorf("Expected postgres_migrate_pg_index.Crud_manual_PostgresMigratePgIndex to be set to crud, but got %+v", postgres_migrate_pg_index.Crud_manual_PostgresMigratePgIndex)
	}
}
