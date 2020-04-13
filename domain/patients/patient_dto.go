package patients

import (
	"github.com/macrusal/orthopedic-clinic-api/utils/errors"
	"strings"
)

type Patient struct {
	PatientName string
	FullName    string
	Phone       string
	Age         int
	Gender      string
	DateCreated string
}

func (patient *Patient) Validate() *errors.RestErr {
	patient.PatientName = strings.TrimSpace(strings.ToLower(patient.PatientName))
	if patient.PatientName == "" {
		return errors.NewBadRequestError("Invalid empty value to field patients name")
	}

	patient.FullName = strings.TrimSpace(strings.ToLower(patient.FullName))
	if patient.FullName == "" {
		return errors.NewBadRequestError("Invalid empty value to field full name")
	}

	patient.Phone = strings.TrimSpace(strings.ToLower(patient.Phone))
	if patient.Phone == "" {
		return errors.NewBadRequestError("Invalid empty value to field phone")
	}

	if patient.Age < 1 {
		return errors.NewBadRequestError("Invalid age")
	}

	patient.Gender = strings.TrimSpace(strings.ToLower(patient.Gender))
	if patient.Gender == "" {
		return errors.NewBadRequestError("Invalid empty value to field gender")
	}
	return nil
}
