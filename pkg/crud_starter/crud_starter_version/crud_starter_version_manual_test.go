package crud_starter_version

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/crud/crud_version"
	"github.com/ManyakRus/postgres_migrate/pkg/object_model/entities/version"
	"testing"
)

func TestSetCrudManualInterface(t *testing.T) {

	crud := crud_version.Crud_DB{}

	SetCrudManualInterface(crud)

	// Test that the crud variable is set correctly
	if version.Crud_manual_Version != crud {
		t.Errorf("Expected version.Crud_manual_Version to be set to crud, but got %+v", version.Crud_manual_Version)
	}
}
