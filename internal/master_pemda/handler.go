package master_pemda

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPemda(c *gin.Context) {
	pemda, err := GetMasterPemdaServices()

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

func GetPemdaID(c *gin.Context) {
	id := c.Param("id")

	pemda, err := GetMasterPemdaServicesID(id)

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

func CreatePemda(c *gin.Context) {

}

func UpdatePemda(c *gin.Context) {

}

func DeletePemda(c *gin.Context) {

}