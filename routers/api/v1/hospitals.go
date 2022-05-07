package v1

import (
	"net/http"
	"strconv"
	"yslinear/go-covid19/service/fst_service"
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

func GetHospital(c *gin.Context) {
	code, err := strconv.ParseUint(c.Param("code"), 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	hospital_service := hospital_service.Hospital{
		Code: int(code),
	}
	data := make(map[string]interface{})
	data["info"], err = hospital_service.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fst_service := fst_service.Fst{
		HospitalCode: int(code),
	}

	data["fst"], err = fst_service.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
	})
}
