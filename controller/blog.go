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
		if total == 0 {
			total = 1
		}
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
		temp, err := service.GetBlogByTitle(title)
		if err != nil && err.Error() != "record not found" {
			log.Logger.Error("获取博客失败: ", err)
			c.JSON(http.StatusOK, gin.H{"msg": err.Error(), "status": statusErr})
			return
		}
		if temp.Title != "" {
			c.JSON(http.StatusOK, gin.H{"msg": "标题重复", "status": statusErr})
			return
		}

		blog.CreateDate = blog.UpdateDate
		//insert
		err = service.AddBlog(blog)
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
	//获取参数
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	index := c.DefaultQuery("index", "0")
	pageIndex, _ := strconv.Atoi(index)

	//查询博客
	blog, err := service.FindBlogById(id)
	if err != nil {
		c.HTML(http.StatusOK, "user/blog.html", gin.H{"error": err.Error()})
		return
	}

	//查询绑定的评论
	//首先获取10条父评论（一页的标准）
	parentComments, err := service.GetParentCommentsByPage(id, pageIndex)
	if err != nil {
		log.Logger.Error("获取留言板信息失败: ", err)
		c.HTML(http.StatusOK, "user/blog.html", gin.H{"error": err})
		return
	}
	//获取构造好的结果
	commentDtos, err := service.CreateCommentDtos(parentComments)
	if err != nil {
		log.Logger.Error("构造留言板信息失败: ", err)
		c.HTML(http.StatusOK, "user/blog.html", gin.H{"error": err})
		return
	}
	//获取所有该博客的评论数量
	count, err := service.CountComment(id, -2)
	if err != nil {
		log.Logger.Error("获取留言板信息总数失败: ", err)
		c.HTML(http.StatusOK, "user/blog.html", gin.H{"error": err})
		return
	}
	//获取所有父评论数量
	parent_count, err := service.CountComment(id, -1)
	if err != nil {
		log.Logger.Error("获取留言板父评论信息总数失败: ", err)
		c.HTML(http.StatusOK, "user/blog.html", gin.H{"error": err})
		return
	}
	//计算页数
	//因为分页标准是按照每页父评论数量10条规定的
	//所以用所有父评论的数量来计算页数
	total := 0
	if parent_count%10 == 0 {
		total = parent_count / 10
		//防止显示总页数为0，因为不管有没有评论，总要展示第一页，要是总页数为0，岂不是显示成1/0,应该是1/1
		if total == 0 {
			total = 1
		}
	} else {
		total = parent_count/10 + 1
	}

	//更新浏览数
	blog.ViewCount++
	err = service.UpdateBlog(blog)
	if err != nil {
		log.Logger.Error("更新博客失败: ", err)
		c.HTML(http.StatusOK, "user/blog.html", gin.H{"error": err})
		return
	}

	//判断当前用户是否点过赞
	isLike := false
	username := c.GetString("username")
	var user_id int
	if username != "" {
		user_id = c.GetInt("user_id")
		isLike, err = service.CheckLike(blog.Id, user_id)
		if err != nil {
			log.Logger.Error("判断用户是否点赞过此文章失败: ", err)
			c.HTML(http.StatusOK, "user/blog.html", gin.H{"error": err})
			return
		}
	}

	c.HTML(http.StatusOK, "user/blog.html", gin.H{"isLike": isLike, "blog": blog, "comments": commentDtos, "userId": user_id, "username": username, "count": count, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1})
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

//TypeBlogList 分类的博客列表
func TypeBlogList(c *gin.Context) {
	//获取参数
	typeStr := c.Query("type")
	index := c.Query("index")
	if index == "" {
		index = "0"
	}
	pageIndex, _ := strconv.Atoi(index)

	//查询博客
	blogs, err := service.GetBlog("", typeStr, pageIndex)
	if err != nil {
		log.Logger.Error("获取数据库的博客数据失败: ", err)
		c.HTML(http.StatusOK, "user/types.html", gin.H{"error": err.Error()})
		return
	}
	//查询总条数
	count, err := service.GetAllBlogCount("", typeStr)
	if err != nil {
		log.Logger.Error("获取数据库的博客总条数失败: ", err)
		c.HTML(http.StatusOK, "user/types.html", gin.H{"error": err.Error()})
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

	//查询type数量
	typeCounts := make(map[string]int, 0)
	allBlog, err := service.GetAllBlog()
	if err != nil {
		log.Logger.Error("获取数据库的博客数据失败: ", err)
		c.HTML(http.StatusOK, "user/types.html", gin.H{"error": err.Error()})
		return
	}
	for _, v := range allBlog {
		if _, ok := typeCounts[v.Type]; ok {
			typeCounts[v.Type]++
		} else {
			typeCounts[v.Type] = 1
		}
	}

	c.HTML(http.StatusOK, "user/types.html", gin.H{"blogs": blogs, "count": count, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1, "type": typeStr, "typeCounts": typeCounts})
}

//SearchBlog 搜索指定标题的博客
func SearchBlog(c *gin.Context) {
	//获取标题
	title := c.Query("title")
	//获取当前页数
	index := c.Query("index")
	if index == "" {
		index = "0"
	}
	pageIndex, _ := strconv.Atoi(index)

	//查询博客
	var blogs []model.Blog
	var count int
	var err error
	if title == "" {
		blogs, err = service.GetBlog("", "", pageIndex)
		if err != nil {
			log.Logger.Error("获取数据库的博客数据失败: ", err)
			c.HTML(http.StatusOK, "user/search.html", gin.H{"error": err.Error()})
			return
		}
		count, err = service.GetAllBlogCount("", "")
		if err != nil {
			log.Logger.Error("获取数据库的博客总条数失败: ", err)
			c.HTML(http.StatusOK, "user/search.html", gin.H{"error": err.Error()})
			return
		}
	} else {
		blogs, err = service.GetBlogLikeTitle(title, pageIndex)
		if err != nil {
			log.Logger.Error("获取数据库的博客数据失败: ", err)
			c.HTML(http.StatusOK, "user/search.html", gin.H{"error": err.Error()})
			return
		}
		count, err = service.CountBlogLikeTitle(title)
		if err != nil {
			log.Logger.Error("获取数据库的模糊查询的博客总条数失败: ", err)
			c.HTML(http.StatusOK, "user/search.html", gin.H{"error": err.Error()})
			return
		}
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

	c.HTML(http.StatusOK, "user/search.html", gin.H{"blogs": blogs, "count": count, "total": total, "pre": pageIndex - 1, "page": pageIndex + 1, "title": title})
}

//Archives 时间轴
func Archives(c *gin.Context) {
	blogs, err := service.GetArchivesBlog()
	if err != nil {
		log.Logger.Error("获取数据库的时间轴博客数据失败: ", err)
		c.HTML(http.StatusOK, "user/archives.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "user/archives.html", gin.H{"blogs": blogs})
}
