package crud_starter_pg_constraint

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_pg_constraint"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/pg_constraint"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_pg_constraint.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if pg_constraint.Crud_manual_PgConstraint != crud {
		t.Errorf("Expected pg_constraint.Crud_manual_PgConstraint to be set to crud, but got %+v", pg_constraint.Crud_manual_PgConstraint)
	}
}
