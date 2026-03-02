package permintaan

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPermintaan(c *gin.Context) {
	permintaan, err := GetPermintaanServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, APIResponse{
		Code   : 200,
		Success: true,
		Message: "Berhasil mengambil data",
		Data: permintaan,
	})
}

func GetPermintaanId(c *gin.Context) {

}

func CreatePermintaan(c *gin.Context) {

}

func UpdatePermintaan(c *gin.Context) {

}

func DeletePermintaan(c *gin.Context) {

}