package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	user := CurrentUser(c)
	service := service.CreateVideoService{}
	//c.ShouldBind(&service)将前端的数据绑定到结构体内
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListVideo 视频列表详情接口
func ListVideo(c *gin.Context) {
	service := service.ListVideoService{}
	//c.ShouldBind(&service)将前端的数据绑定到结构体内
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateVideo 更新视频的接口
func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	//c.ShouldBind(&service)将前端的数据绑定到结构体内
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo 删除视频的接口
func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

