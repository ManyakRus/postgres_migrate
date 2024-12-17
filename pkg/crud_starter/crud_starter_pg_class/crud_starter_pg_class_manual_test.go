package crud_starter_pg_class

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_class"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_class"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_pg_class.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if pg_class.Crud_manual_PgClass != crud {
		t.Errorf("Expected pg_class.Crud_manual_PgClass to be set to crud, but got %+v", pg_class.Crud_manual_PgClass)
	}
}
