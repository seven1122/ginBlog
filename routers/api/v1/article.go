package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"seven1122/ginBlog/models"
	"seven1122/ginBlog/pkg/erorrs"
	"seven1122/ginBlog/pkg/logging"
	"seven1122/ginBlog/pkg/setting"
	"seven1122/ginBlog/pkg/utils"
)

//获取单个文章
func GetArticle(c *gin.Context)  {
	id, _ := com.StrTo(c.Param("id")).Int()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := erorrs.INVALID_PARAMS
	var data interface{}
	if ! valid.HasErrors(){
		if models.ExistArticleByID(id){
			data = models.GetArticle(id)
			code = erorrs.SUCCESS
		}else{
			code = erorrs.ERROR_NOT_EXIST_ARTICLE
		}
	}else{
		for _, err := range valid.Errors{
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": erorrs.GetMsg(code),
		"data": data,
	})

}

//获取多篇文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != ""{
		state, _ = com.StrTo(arg).Int()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态值只能是0或者1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != ""{
		tagId, _ = com.StrTo(arg).Int()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := erorrs.INVALID_PARAMS
	if ! valid.HasErrors(){
		code = erorrs.SUCCESS
		data["list"] = models.GetArticles(utils.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	}else{
		for _, err := range valid.Errors{
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":erorrs.GetMsg(code),
		"data": data,
	})

}

//新增文章
func AddArticle(c *gin.Context){
	tagId, _ := com.StrTo(c.PostForm("tag_id")).Int()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	createdBy := c.PostForm("created_by")
	state, _ := com.StrTo(c.DefaultPostForm("state", "0")).Int()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题是必填")
	valid.Required(desc, "desc").Message("简述是必填")
	valid.Required(content, "content").Message("内容是必填")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0,1, "state").Message("状态值只能是0或者1")

	code := erorrs.INVALID_PARAMS
	if !valid.HasErrors(){
		if models.ExistTagByID(tagId){
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			models.AddArticle(data)
			code = erorrs.SUCCESS
		}else{
			code = erorrs.ERROR_NOT_EXIST_TAG
		}
	}else{
		for _, err := range valid.Errors{
			logging.Error(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": erorrs.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//修改文章
func EditArticle(c *gin.Context)  {
	valid := validation.Validation{}
	id,_ := com.StrTo(c.Param("id")).Int()
	tagID, _ := com.StrTo(c.PostForm("tag_id")).Int()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")

	var state int = -1
	if arg := c.PostForm("state"); arg != ""{
		state, _ =  com.StrTo(arg).Int()
		valid.Range(state, 0, 1, "state").Message("状态值只能是0或者1")
	}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长允许100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	code := erorrs.INVALID_PARAMS
	if ! valid.HasErrors(){
		if models.ExistArticleByID(id){
			if models.ExistTagByID(tagID){
				data := make(map[string]interface{})
				data["tag_id"]  = tagID
				if title != ""{
					data["title"] = title
				}
				if desc != ""{
					data["desc"] = desc
				}
				if content != ""{
					data["content"] = content
				}
				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = erorrs.SUCCESS
			}else{
				code = erorrs.ERROR_NOT_EXIST_TAG
			}
		}else{
			code = erorrs.ERROR_NOT_EXIST_ARTICLE
		}
	}else{
		for _, err := range valid.Errors{
			logging.Error(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": erorrs.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

//删除文章
func DeleteArticle(c *gin.Context)  {
	id, _ := com.StrTo(c.Param("id")).Int()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := erorrs.INVALID_PARAMS
	if !valid.HasErrors(){
		if models.ExistArticleByID(id){
			models.DeleteArticle(id)
			code = erorrs.SUCCESS
		}else{
			code = erorrs.ERROR_NOT_EXIST_ARTICLE
		}
	}else{
		for _, err := range valid.Errors{
			logging.Error(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":erorrs.GetMsg(code),
		"data": make(map[string]string),
	})

}
