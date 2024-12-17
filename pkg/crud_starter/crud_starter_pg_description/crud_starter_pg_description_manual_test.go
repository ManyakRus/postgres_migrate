package crud_starter_pg_description

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_description"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_description"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_pg_description.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if pg_description.Crud_manual_PgDescription != crud {
		t.Errorf("Expected pg_description.Crud_manual_PgDescription to be set to crud, but got %+v", pg_description.Crud_manual_PgDescription)
	}
}
