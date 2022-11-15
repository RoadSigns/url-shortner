package redirect

import (
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/roadsigns/url-shortner/repositories"
	"github.com/roadsigns/url-shortner/services"
	"net/http"
)

func Execute(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	service := services.GetUrl{
		Repository: repositories.CassandraUrlRepository{
			Cluster: gocql.NewCluster("host.docker.internal"),
		},
	}

	urlData, err := service.Repository.Get(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(302, urlData.LongUrl)
	return
}
