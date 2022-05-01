package v1

import (
	"net/http"
	"yslinear/go-covid19/service/hospital_service"

	"github.com/gin-gonic/gin"
)

func GetHospitals(c *gin.Context) {
	hospitalService := hospital_service.Hospital{
		City: c.Query("city"),
	}
	if c.Query("city") != "" {
		hospitalService.District = c.Query("district")
	}

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

func GetAllHospitalCities(c *gin.Context) {
	cities, err := hospital_service.GetAllCities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    cities,
	})
}

func GetAllHospitalDistricts(c *gin.Context) {
	city := c.Param("city")
	districts, err := hospital_service.GetAllDistricts(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    districts,
	})
}
