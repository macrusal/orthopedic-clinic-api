package app

import (
	"github.com/macrusal/orthopedic-clinic-api/controllers/patients"
	"github.com/macrusal/orthopedic-clinic-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/patients/:patientName", patients.GetPatient)
	router.POST("/patients", patients.CreatePatient)
}