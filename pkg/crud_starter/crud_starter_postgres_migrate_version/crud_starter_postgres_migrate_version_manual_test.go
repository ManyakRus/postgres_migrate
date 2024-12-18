package crud_starter_postgres_migrate_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_version"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_version"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_postgres_migrate_version.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if postgres_migrate_version.Crud_manual_PostgresMigrateVersion != crud {
		t.Errorf("Expected postgres_migrate_version.Crud_manual_PostgresMigrateVersion to be set to crud, but got %+v", postgres_migrate_version.Crud_manual_PostgresMigrateVersion)
	}
}
