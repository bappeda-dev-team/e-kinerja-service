package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func Auth(c *gin.Context) {

}

func Register(c *gin.Context) {
	register, err := GetRegisterServices()

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
		Data: register,
	})
}

func CreateRegister(c *gin.Context) {

}

func Logout(c *gin.Context) {

}