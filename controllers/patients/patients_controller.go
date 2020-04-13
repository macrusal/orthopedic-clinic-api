package patients

import (
	"github.com/gin-gonic/gin"
	"github.com/macrusal/orthopedic-clinic-api/domain/patients"
	"github.com/macrusal/orthopedic-clinic-api/services"
	"github.com/macrusal/orthopedic-clinic-api/utils/errors"
	"net/http"
)

func CreatePatient(c *gin.Context) {
	var patient patients.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreatePatient(patient)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetPatient(c *gin.Context) {
	patient, getErr := services.GetPatient(c.Param("patientName"))
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, patient)
}