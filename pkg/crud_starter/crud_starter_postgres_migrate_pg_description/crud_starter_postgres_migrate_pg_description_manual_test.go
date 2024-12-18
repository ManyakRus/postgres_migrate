package crud_starter_postgres_migrate_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_description"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_postgres_migrate_pg_description.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if postgres_migrate_pg_description.Crud_manual_PostgresMigratePgDescription != crud {
		t.Errorf("Expected postgres_migrate_pg_description.Crud_manual_PostgresMigratePgDescription to be set to crud, but got %+v", postgres_migrate_pg_description.Crud_manual_PostgresMigratePgDescription)
	}
}
