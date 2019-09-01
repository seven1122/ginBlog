package main

import (
	"fmt"
	"net/http"

	"seven1122/ginBlog/models"
	"seven1122/ginBlog/pkg/setting"
	"seven1122/ginBlog/routers"
)

// @title ginBlog APIS
// @version 1.0
// @description include all blog apis
// @termsOfService http://swagger.io/terms/

// @contact.name seven007
// @contact.url http://www.swagger.io/support
// @contact.email 931880645@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath api/v1
func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	models.CloseDB()
}