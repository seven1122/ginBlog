package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"seven1122/ginBlog/models"
	"seven1122/ginBlog/pkg/erorrs"
	"seven1122/ginBlog/pkg/logging"
	"seven1122/ginBlog/pkg/setting"
	"seven1122/ginBlog/pkg/utils"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		maps["state"] = state
	}

	data["list"] = models.GetTags(utils.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagCount(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": erorrs.SUCCESS,
		"msg":  erorrs.GetMsg(erorrs.SUCCESS),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	name := c.PostForm("name")
	state, _ := com.StrTo(c.DefaultPostForm("state", "0")).Int()
	createdBy := c.PostForm("created_by")
	// 校验参数
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100")
	valid.Range(state, 0, 1, "state").Message("状态值只能是0或者1")

	code := erorrs.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = erorrs.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = erorrs.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  erorrs.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTag(c *gin.Context) {
	id, _ := com.StrTo(c.Param("id")).Int()
	name := c.PostForm("name")
	modifiedBy := c.PostForm("modified_by")
	valid := validation.Validation{}

	var state int = -1
	if arg := c.PostForm("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		valid.Range(state, 0, 1, "state").Message("状态只允许是0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人不能超过100长度")
	valid.MaxSize(name, 100, "name").Message("名称最长不要超过100")
	code := erorrs.INVALID_PARAMS
	if !valid.HasErrors() {
		code = erorrs.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = erorrs.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  erorrs.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteTag(c *gin.Context) {
	id, _ := com.StrTo(c.Param("id")).Int()
	valid := validation.Validation{}
	valid.Min(id, 1, "").Message("ID必须大于0")
	code := erorrs.INVALID_PARAMS
	if !valid.HasErrors() {
		code = erorrs.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = erorrs.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  erorrs.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}
