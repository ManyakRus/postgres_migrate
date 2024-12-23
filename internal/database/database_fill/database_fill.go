package database_fill

import (
	"fmt"
	"github.com/ManyakRus/starter/log"
)

// Fill_All - заполняет все таблицы postgres_migrate
func Fill_All(VersionID int64) error {
	var err error

	//namespace
	err = Fill_namespace(VersionID)
	if err != nil {
		err = fmt.Errorf("Fill_namespace() error: %w", err)
		log.Error(err)
		return err
	}

	//attribute
	err = Fill_attribute(VersionID)
	if err != nil {
		err = fmt.Errorf("Fill_attribute() error: %w", err)
		log.Error(err)
		return err
	}

	//class
	err = Fill_class(VersionID)
	if err != nil {
		err = fmt.Errorf("Fill_class() error: %w", err)
		log.Error(err)
		return err
	}

	//constraint
	err = Fill_constraint(VersionID)
	if err != nil {
		err = fmt.Errorf("Fill_constraint() error: %w", err)
		log.Error(err)
		return err
	}

	//description
	err = Fill_description(VersionID)
	if err != nil {
		err = fmt.Errorf("Fill_description() error: %w", err)
		log.Error(err)
		return err
	}

	//index
	err = Fill_index(VersionID)
	if err != nil {
		err = fmt.Errorf("Fill_index() error: %w", err)
		log.Error(err)
		return err
	}

	return err
}
