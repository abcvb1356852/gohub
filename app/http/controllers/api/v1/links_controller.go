package v1

import (
	"gohub/app/models/link"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	// links := link.All() 不直接访问数据库，改为调用缓存中数据
	response.Data(c, link.AllCached())
}
