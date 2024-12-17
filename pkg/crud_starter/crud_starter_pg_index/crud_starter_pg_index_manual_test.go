package crud_starter_pg_index

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_index"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_index"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_pg_index.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if pg_index.Crud_manual_PgIndex != crud {
		t.Errorf("Expected pg_index.Crud_manual_PgIndex to be set to crud, but got %+v", pg_index.Crud_manual_PgIndex)
	}
}
