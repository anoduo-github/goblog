package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//BlogList 博客列表
func BlogList(c *gin.Context) {
	/*
		查询操作，结果返回到页面上
	*/
	c.HTML(http.StatusOK, "blogList.html", gin.H{})
}

//BlogEdit 博客新增
func BlogEdit(c *gin.Context) {
	/*
		新增操作，结果返回到页面上
	*/
	c.HTML(http.StatusOK, "blogEdit.html", gin.H{})
}

//BlogDetails 博客详情
func BlogDetails(c *gin.Context) {
	/*
		博客详情，结果返回到页面上
	*/
	c.HTML(http.StatusOK, "blogDetails.html", gin.H{})
}
