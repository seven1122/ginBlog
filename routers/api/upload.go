package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seven1122/ginBlog/pkg/erorrs"
	"seven1122/ginBlog/pkg/logging"
	"seven1122/ginBlog/pkg/upload"
)

func UploadImage(c *gin.Context) {
	code := erorrs.SUCCESS
	data := make(map[string] string)
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = erorrs.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg": erorrs.GetMsg(code),
			"data": data,
		})
	}
	if image == nil {
		code = erorrs.INVALID_PARAMS
	}else{
		imageMame := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageMame
		if ! upload.CheckImageExt(imageMame) || !upload.CheckImageSize(file) {
			code = erorrs.ERROR_UPLOAD_CHEXK_IMAGE_FORMAT
		}else{
			err := upload.CheckImage(fullPath)
			if err != nil{
				logging.Warn(err)
				code = erorrs.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			}else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = erorrs.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			}else{
				data["image_url"] = upload.GetImageFullUrl(imageMame)
				data["image_save_url"] = savePath + imageMame
			}
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": erorrs.GetMsg(code),
		"data": data,
	})


}