package v1

import (
	"net/http"
	"yslinear/go-covid19/service/hospital_service"

	"github.com/gin-gonic/gin"
)

func GetHospitals(c *gin.Context) {
	hospitalService := hospital_service.Hospital{}

	total, err := hospitalService.Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	hospitals, err := hospitalService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := make(map[string]interface{})
	data["lists"] = hospitals
	data["total"] = total

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
	})
}
