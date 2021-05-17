package api

import (
	"bluebell/pkg/errno"
	"bluebell/repository"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


var communityService *service.CommunityService

func init() {
	communityService, _ = (service.NewCommunityService(repository.NewCommunityRepository(""))).(*service.CommunityService)
}

func CommunityHandler(c *gin.Context)  {
	list, _ := communityService.SelectCommunityList()
	c.JSON(http.StatusOK, errno.OK.WithData(list))
}
//
 //func CommunityDetailHandler(c *gin.Context) {
//	postID := c.Params.Get("id")
//	id, err := strconv.ParseInt(postID, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusOK, errno.ErrParam)
//		return
//	}
//
//
//}