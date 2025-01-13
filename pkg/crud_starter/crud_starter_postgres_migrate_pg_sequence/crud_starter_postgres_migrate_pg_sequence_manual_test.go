package crud_starter_postgres_migrate_pg_sequence

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_postgres_migrate_pg_sequence"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/postgres_migrate_pg_sequence"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_postgres_migrate_pg_sequence.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if postgres_migrate_pg_sequence.Crud_manual_PostgresMigratePgSequence != crud {
		t.Errorf("Expected postgres_migrate_pg_sequence.Crud_manual_PostgresMigratePgSequence to be set to crud, but got %+v", postgres_migrate_pg_sequence.Crud_manual_PostgresMigratePgSequence)
	}
}
