package master_aplikasi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAplikasi(c *gin.Context) {
	aplikasi, err := GetMasterAplikasiServices()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Berhasil mengambil data",
		Data: aplikasi,
	})
}

func GetAplikasiID(c *gin.Context) {

}

func CreateAplikasi(c *gin.Context) {

}

func UpdateAplikasi(c *gin.Context) {

}

func DeleteAplikasi(c *gin.Context) {

}