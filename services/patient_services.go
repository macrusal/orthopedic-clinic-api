package services

import (
	"github.com/macrusal/orthopedic-clinic-api/domain/patients"
	"github.com/macrusal/orthopedic-clinic-api/utils/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func FindPatient(patientName string) (patients.Patient, error) {
	session, err := mgo.Dial("127.0.0.1")
	p := patients.Patient{}
	if err != nil {
		return p, err
	}
	defer session.Close()
	c := session.DB("orthopedic_clinic_db").C(patients.COLLECTION)
	if err = c.Find(bson.M{"PatientName":patientName}).One(&p); err != nil {
		return p, err
	}
	return p, nil
}

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