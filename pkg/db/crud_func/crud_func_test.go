package crud_func

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"errors"
	"testing"
)

func TestIsRecordNotFound(t *testing.T) {

	err := errors.New(db_constants.TEXT_RECORD_NOT_FOUND + " !")
	Otvet := IsRecordNotFound(err)
	if Otvet != true {
		t.Error("TestIsRecordNotFound() error: false")
	}

	err = errors.New("rpc error: code = Unknown desc = record not found")
	Otvet = IsRecordNotFound(err)
	if Otvet != true {
		t.Error("TestIsRecordNotFound() error: false")
	}

}
