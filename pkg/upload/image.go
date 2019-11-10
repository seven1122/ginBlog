package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"seven1122/ginBlog/pkg/file"
	"seven1122/ginBlog/pkg/logging"
	"seven1122/ginBlog/pkg/setting"
	"seven1122/ginBlog/pkg/utils"
	"strings"
)

func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name

}


func GetImagePath() string{
	return setting.AppSetting.ImageSavePath
}

func GetImageName(name string ) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)

	return fileName + ext
}

func GetImageFullPath() string {
	return setting.AppSetting.RunTimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExt{
		if strings.ToUpper(allowExt) == strings.ToUpper(ext){
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size , err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}
	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := file.CheckPermission(src)
	if perm == true{
		return fmt.Errorf("file.CheckPermission denied src: %s", src)
	}
	return nil

}