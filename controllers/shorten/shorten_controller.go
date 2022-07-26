package shorten

import (
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/roadsigns/url-shortner/repositories"
	"github.com/roadsigns/url-shortner/services"
	"github.com/roadsigns/url-shortner/shortener"
	"net/http"
)

func Execute(c *gin.Context) {
	var request shortenRequestData

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is not allowed to be empty"})
		return
	}

	service := services.StoreUrl{
		Shortener: shortener.Base62Shortener{},
		Repository: repositories.CassandraUrlRepository{
			Cluster: gocql.NewCluster("host.docker.internal"),
		},
	}

	shortUrl, err := service.Save(request.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": shortUrl})
	return
}

type shortenRequestData struct {
	Url string
}
