package roles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	roles, err := GetRolesServices()

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
		Data: roles,
	})

}
