package pelaksana

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPelaksana(c *gin.Context) {
	pelaksana, err := GetPelaksanaServices()

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
		Data: pelaksana,
	})
}

func GetPelaksanaID(c *gin.Context) {
	id := c.Param("id")

	pemda, err := GetPelaksanaServicesID(id)

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
		Data: pemda,
	})
}

func CreatePelaksana(c *gin.Context) {

}

func UpdatePelaksana(c *gin.Context) {

}

func DeletePelaksana(c *gin.Context) {

}