package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"seven1122/ginBlog/models"
	"seven1122/ginBlog/pkg/erorrs"
	"seven1122/ginBlog/pkg/logging"
	"seven1122/ginBlog/pkg/utils"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	valid := validation.Validation{}
	a := auth{Username:username, Password:password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := erorrs.INVALID_PARAMS
	if ok{
		isExist := models.CheckAuth(username, password)
		if isExist{
			token, err := utils.GenerateToken(username, password)
			if err != nil{
				code = erorrs.ERROR_AUTH_TOKEN
				logging.Error(err)
			}else {
				data["token"] = token
				code = erorrs.SUCCESS
			}
		}else{
			code = erorrs.ERROR_AUTH
		}
	}else{
		for _, err := range valid.Errors{
			logging.Error(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": erorrs.GetMsg(code),
		"data": data,
	})


}