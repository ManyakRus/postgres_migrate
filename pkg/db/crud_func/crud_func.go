package crud_func

import (
	"github.com/ManyakRus/postgres_migrate/pkg/db/db_constants"
	"strings"
)

// IsRecordNotFound - возвращает true если ошибка = "record not found"
func IsRecordNotFound(err error) bool {
	Otvet := false

	if err == nil {
		return Otvet
	}

	TextErr := err.Error()
	pos1 := strings.Index(TextErr, db_constants.TEXT_RECORD_NOT_FOUND)
	if pos1 >= 0 {
		Otvet = true
	}

	return Otvet
}
