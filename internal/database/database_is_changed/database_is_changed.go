package database_is_changed

import (
	"fmt"
	"github.com/ManyakRus/starter/log"
	"strconv"
)

func IsChanged_any() (string, error) {
	Otvet := ""
	var err error

	//IsChanged := false

	//attribute
	Count_Attribute, err := IsChanged_Attribute()
	if err != nil {
		err = fmt.Errorf("IsChanged_Attribute() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//class
	Count_Class, err := IsChanged_Class()
	if err != nil {
		err = fmt.Errorf("IsChanged_Class() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//constraint
	Count_Constraint, err := IsChanged_Constraint()
	if err != nil {
		err = fmt.Errorf("IsChanged_Constraint() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//description
	Count_Description, err := IsChanged_Description()
	if err != nil {
		err = fmt.Errorf("IsChanged_Description() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//index
	Count_Index, err := IsChanged_Index()
	if err != nil {
		err = fmt.Errorf("IsChanged_Index() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//namespace
	Count_Namespace, err := IsChanged_Namespace()
	if err != nil {
		err = fmt.Errorf("IsChanged_Namespace() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//sequence
	Count_Sequence, err := IsChanged_Sequence()
	if err != nil {
		err = fmt.Errorf("IsChanged_Sequence() error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//
	if Count_Attribute == 0 && Count_Class == 0 && Count_Constraint == 0 && Count_Description == 0 && Count_Index == 0 && Count_Namespace == 0 && Count_Sequence == 0 {
		return Otvet, err
	}

	//есть изменения
	Otvet = strconv.Itoa(Count_Attribute) + " + " + strconv.Itoa(Count_Class) + " + " + strconv.Itoa(Count_Constraint) + " + " + strconv.Itoa(Count_Description) + " + " + strconv.Itoa(Count_Index) + " + " + strconv.Itoa(Count_Namespace) + " changes"

	return Otvet, err
}
