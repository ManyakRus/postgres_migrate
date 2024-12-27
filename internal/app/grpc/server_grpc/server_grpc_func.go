package server_grpc

import (
	"github.com/ManyakRus/postgres_migrate/pkg/network/grpc/grpc_constants"
	"errors"
	"github.com/ManyakRus/starter/micro"
)

// ErrorModelVersion
func ErrorModelVersion(Model interface{}) error {
	var err error

	TypeName := micro.GetType(Model)

	s := grpc_constants.TEXT_ERROR_MODEL_VERSION + " " + TypeName
	err = errors.New(s)
	return err
}
