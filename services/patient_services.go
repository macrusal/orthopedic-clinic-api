package services

import (
	"github.com/macrusal/orthopedic-clinic-api/domain/patients"
	"github.com/macrusal/orthopedic-clinic-api/utils/errors"
)

func GetPatient(patientName string) (*patients.Patient, *errors.RestErr) {

	result := &patients.Patient{PatientName:patientName}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreatePatient(patient patients.Patient) (*patients.Patient, *errors.RestErr) {
	if err := patient.Validate(); err != nil {
		return nil, err
	}
	if err := patient.Save(); err != nil {
		return nil, err
	}
	return &patient, nil
}