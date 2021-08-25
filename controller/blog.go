package controller

import (
	"goblog/model"
	"goblog/module/log"
	"goblog/service"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//BlogList 博客列表
func BlogList(c *gin.Context) {
	//解析各参数
	title := c.Query("title")
	typeStr := c.Query("type")
	index := c.Query("index")
	if index == "" {
		index = "0"
	}
	pageIndex, _ := strconv.Atoi(index)

	//查询指定页面的数据
	//默认每页10条数据，写死的，不支持用户更改
	blogs, err := service.GetBlog(title, typeStr, pageIndex)
	if err != nil {
		log.Logger.Error("获取数据库的博客数据失败: ", err)
		c.HTML(http.StatusOK, "admin/blogs.html", gin.H{"error": err.Error()})
		return
	}
	//查询总条数
	count, err := service.GetAllBlogCount(title, typeStr)
	if err != nil {
		log.Logger.Error("获取数据库的博客总条数失败: ", err)
		c.HTML(http.StatusOK, "admin/blogs.html", gin.H{"error": err.Error()})
		return
	}
	//查询分类列表
	types, err := service.GetAllType()
	if err != nil {
		log.Logger.Error("获取数据库的分类总条数失败: ", err)
		c.HTML(http.StatusOK, "admin/blogs.html", gin.H{"error": err.Error()})
		return
	}
	//计算总页数
	total := 0
	if count%10 == 0 {
		total = count / 10
	} else {
		total = count/10 + 1
	}

	//博客数据、分类数据、博客总数、页面总数、前一页、当前页数、特定标题、特定类型
	c.HTML(http.StatusOK, "admin/blogs.html", gin.H{"blogs": blogs, "types": types, "count": count, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1, "title": title, "type": typeStr})
}

//BlogEdit 博客编辑
func BlogEdit(c *gin.Context) {
	//首先获取所有分类
	types, err := service.GetAllType()
	if err != nil {
		log.Logger.Error("get blog type error: ", err)
		c.HTML(http.StatusOK, "admin/blog-input.html", gin.H{"error": err.Error()})
		return
	}
	//其次获取url参数
	idStr := c.Param("id")
	if idStr == "" {
		//表示新增
		c.HTML(http.StatusOK, "admin/blog-input.html", gin.H{"types": types})
		return
	}
	id, _ := strconv.Atoi(idStr)
	//表示修改
	blog, err := service.GetBlogById(id)
	if err != nil {
		log.Logger.Error("edit blog error: ", err)
		c.HTML(http.StatusOK, "admin/blog-input.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "admin/blog-input.html", gin.H{"blog": blog, "types": types})
}

//BlogSubmit 提交博客
func BlogSubmit(c *gin.Context) {
	//获取参数
	//标签
	tag := c.PostForm("tag")
	//标题
	title := c.PostForm("title")
	//分类
	typeStr := c.PostForm("type")
	//md内容
	str := c.PostForm("content")
	//背景图
	pic := c.PostForm("firstPicture")
	//描述
	des := c.PostForm("description")
	//获取id，如果有则表示修改博客，没有就是创建
	idStr := c.PostForm("id")

	//创建博客对象
	blog := model.Blog{}
	blog.Context = str
	blog.Title = title
	blog.Author = "admin"
	blog.Tag = tag
	blog.Description = des
	blog.Pircture = pic
	blog.Type = typeStr
	blog.UpdateDate = time.Now().Format("2006-01-02 15:04:05")
	//判断
	if idStr == "" {
		//根据标题查重
		temp, _ := service.GetBlogByTitle(title)
		if temp.Title != "" {
			c.JSON(http.StatusOK, gin.H{"msg": "标题重复", "status": statusErr})
			return
		}
		blog.CreateDate = blog.UpdateDate
		//insert
		err := service.AddBlog(blog)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
	} else {
		id, _ := strconv.Atoi(idStr)
		blog.Id = id
		//update
		err := service.UpdateBlog(blog)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"msg": "提交成功", "status": statusOk})
}

//BlogDetails 博客详情
func BlogDetails(c *gin.Context) {
	//查询
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	blog, err := service.FindBlogById(id)
	if err != nil {
		c.HTML(http.StatusOK, "user/blog.html", gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "user/blog.html", gin.H{"blog": blog})
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
	c.JSON(http.StatusOK, gin.H{"message": "ok", "success": 1, "url": "/" + filePath})
}

//DeleteBlog 删除博客
func DeleteBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := service.DeleteBlog(id)
	if err != nil {
		log.Logger.Error("delete blog error: ", err)
	}
	c.Redirect(302, "/blog/list")
}
