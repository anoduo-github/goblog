package controller

import (
	"goblog/model"
	"goblog/service"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

//BlogList 博客列表
func BlogList(c *gin.Context) {
	/*
		查询操作，结果返回到页面上
	*/
	c.HTML(http.StatusOK, "blogList.html", gin.H{})
}

//BlogEdit 博客编辑
func BlogEdit(c *gin.Context) {
	/*
		新增操作，结果返回到页面上
	*/
	c.HTML(http.StatusOK, "blogEdit.html", gin.H{})
}

//BlogAdd 博客新增
func BlogAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "blogEdit.html", gin.H{})
}

//BlogSubmit 提交博客
func BlogSubmit(c *gin.Context) {
	//获取参数
	//标题
	title := c.PostForm("title")
	//类型
	kind := c.PostForm("kind")
	//md内容
	str := c.PostForm("text")

	//创建博客对象
	article := model.Article{}
	article.ArticleContext = str
	article.ArticleTitle = title
	article.Author = "user"
	article.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	switch kind {
	case "0":
		article.ArticleType = "go"
	case "1":
		article.ArticleType = "mysql"
	case "2":
		article.ArticleType = "linux"
	case "3":
		article.ArticleType = "http"
	default:
		article.ArticleType = "other"
	}

	//insert
	err := service.InsertArticle(article)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "提交成功", "status": statusOk})
}

//BlogDetails 博客详情
func BlogDetails(c *gin.Context) {
	/*
		博客详情，结果返回到页面上
	*/
	c.HTML(http.StatusOK, "blogDetails.html", gin.H{})
}

//UploadPicture 上传图片
func UploadPicture(c *gin.Context) {
	file, err := c.FormFile("editormd-image-file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": 0, "url": ""})
		return
	}
	picturePath := "static/upload/image"
	err = os.MkdirAll(picturePath, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": 0, "url": ""})
		return
	}
	filePath := picturePath + "/" + file.Filename
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "success": 0, "url": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "success": 1, "url": "../" + filePath})
}
