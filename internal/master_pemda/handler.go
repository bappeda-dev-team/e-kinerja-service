package master_pemda

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPemda(c *gin.Context) {
	pemda, err := GetAllMasterPemda()

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, pemda)
}

func GetPemdaID(c *gin.Context) {

}

func CreatePemda(c *gin.Context) {

}

func UpdatePemda(c *gin.Context) {

}

func DeletePemda(c *gin.Context) {

}