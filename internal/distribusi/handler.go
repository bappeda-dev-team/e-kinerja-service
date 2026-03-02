package distribusi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDistribusi(c *gin.Context) {
	distribusi, err := GetDistribusiServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, APIResponse{
		Code: 200,
		Success: true,
		Message: "Berhasil mengambil data",
		Data: distribusi,
	})
}

func GetDistribusiId(c *gin.Context) {

}

func CreateDistribusi(c *gin.Context) {

}

func UpdateDistribusi(c *gin.Context) {

}

func DeleteDistribusi(c *gin.Context) {

}