package patients

import (
	"fmt"
	"github.com/macrusal/orthopedic-clinic-api/datasources/mongodb/patients_db"
	"github.com/macrusal/orthopedic-clinic-api/utils/date_utils"
	"github.com/macrusal/orthopedic-clinic-api/utils/errors"
	"gopkg.in/mgo.v2/bson"
)

const (
	databaseName = "orthopedic_clinic_db"
	COLLECTION = "patients"
)

var (
	patientsDB = make(map[int64]*Patient)
)

func (patient *Patient) Get() *errors.RestErr {
	if err := patients_db.GetSession().Ping(); err != nil {
		panic(err)
	}
	db := patients_db.GetSession().DB(databaseName)
	err := db.C(COLLECTION).Find(bson.M{"patientname":patient.PatientName}).One(&patient)
	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("Patient %d not found", patient.PatientName))
	}
	return nil
}

func (patient *Patient) Save() *errors.RestErr {
	db := patients_db.GetSession().DB(databaseName)
	if db == nil {
		return errors.NewInternalServerError(fmt.Sprintf("Database %s not found", db.Name))
	}
	defer patients_db.GetSession().Close()

	patient.DateCreated = date_utils.GetNowString()

	err := db.C(COLLECTION).Insert(&patient)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save patients: %s", err.Error()))
	}

	return nil
}
