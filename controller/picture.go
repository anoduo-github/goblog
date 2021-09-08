package controller

import (
	"goblog/model"
	"goblog/module/log"
	"goblog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Pictures 照片管理列表页
func Pictures(c *gin.Context) {
	//获取参数
	index := c.Param("index")
	if index == "" {
		index = "0"
	}
	pageIndex, _ := strconv.Atoi(index)

	//获取分页照片
	pictures, err := service.GetPictureByPage(pageIndex)
	if err != nil {
		log.Logger.Error("获取所有照片失败: ", err)
		c.HTML(http.StatusOK, "admin/pictures.html", gin.H{"error": err.Error()})
		return
	}

	//获取照片总数
	count, err := service.GetAllPictureCount()
	if err != nil {
		log.Logger.Error("获取照片总数失败: ", err)
		c.HTML(http.StatusOK, "admin/pictures.html", gin.H{"error": err.Error()})
		return
	}

	//计算总页数
	total := 0
	if count%10 == 0 {
		total = count / 10
		if total == 0 {
			total = 1
		}
	} else {
		total = count/10 + 1
	}

	c.HTML(http.StatusOK, "admin/pictures.html", gin.H{"pictures": pictures, "count": count, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1})
}

//EditPicture 编辑照片
func EditPicture(c *gin.Context) {
	//获取id
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	//跳到新增
	if id == 0 {
		c.HTML(http.StatusOK, "admin/pictures-input.html", gin.H{})
		return
	}

	//根据id获取link
	picture, err := service.GetPictureById(id)
	if err != nil {
		log.Logger.Error("获取照片失败: ", err)
		c.HTML(http.StatusOK, "admin/pictures-input.html", gin.H{"error": err.Error()})
		return
	}

	//跳到修改
	c.HTML(http.StatusOK, "admin/pictures-input.html", gin.H{"picture": picture})
}

//AddPicture 添加或修改照片
func AddPicture(c *gin.Context) {
	//获取表单数据
	idStr := c.PostForm("id")
	name := c.PostForm("name")
	time_place := c.PostForm("time_place")
	description := c.PostForm("description")
	address := c.PostForm("address")

	//创建picture
	var picture model.Picture
	picture.Name = name
	picture.Description = description
	picture.TimePlace = time_place
	picture.Address = address

	//判断idStr,为空表示新增，否则为修改
	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		picture.Id = id
		err := service.UpdatePicture(picture)
		if err != nil {
			log.Logger.Error("更新照片失败: ", err)
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
	} else {
		//查重
		p, err := service.GetPictureByName(name)
		if err != nil && err.Error() != "record not found" {
			log.Logger.Error("获取照片失败: ", err)
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
		if p.Name != "" {
			c.JSON(http.StatusOK, gin.H{"msg": "照片名称重复", "status": statusErr})
			return
		}
		err = service.AddPicture(picture)
		if err != nil {
			log.Logger.Error("新增照片失败: ", err)
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"msg": "提交成功", "status": statusOk})
}

//DeletePicture 删除照片
func DeletePicture(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := service.DeletePicture(id)
	if err != nil {
		log.Logger.Error("删除照片失败: ", err)
	}
	c.Redirect(302, "/picture/list/0")
}

//PicturePage 照片墙
func PicturePage(c *gin.Context) {
	pictures, err := service.GetAllPicture()
	if err != nil {
		log.Logger.Error("获取所有照片失败: ", err)
		c.HTML(http.StatusOK, "user/picture.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "user/picture.html", gin.H{"pictures": pictures})
}
