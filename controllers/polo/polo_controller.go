package polo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const polo = "polo"

func Polo(c *gin.Context) {
	c.JSON(http.StatusOK, polo)
}
