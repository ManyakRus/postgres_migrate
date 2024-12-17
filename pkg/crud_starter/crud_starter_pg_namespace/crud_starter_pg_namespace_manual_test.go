package crud_starter_pg_namespace

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_namespace"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_namespace"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_pg_namespace.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if pg_namespace.Crud_manual_PgNamespace != crud {
		t.Errorf("Expected pg_namespace.Crud_manual_PgNamespace to be set to crud, but got %+v", pg_namespace.Crud_manual_PgNamespace)
	}
}
