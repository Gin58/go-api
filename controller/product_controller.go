package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	product "../service"
)

type Controller struct {

}

func (pc Controller) Index(c *gin.Context) {
	// Get parameter
	title := c.Query("title")

	var s product.Service
	p, err := s.Search(title)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

