package crud_starter_pg_attribute

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_attribute"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_attribute"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_pg_attribute.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if pg_attribute.Crud_manual_PgAttribute != crud {
		t.Errorf("Expected pg_attribute.Crud_manual_PgAttribute to be set to crud, but got %+v", pg_attribute.Crud_manual_PgAttribute)
	}
}
