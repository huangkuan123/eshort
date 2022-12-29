package base

import (
	"eshort/pkg/app_file_path"
	"eshort/pkg/array_tool"
	"eshort/pkg/auth"
	"eshort/pkg/csrf_token"
	"eshort/pkg/helpers"
	rsp "eshort/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type CommonController struct {
	BaseController
}

var upload_allow_ext = []string{"jpg", "png", "jpeg", "doc", "xls", "txt", "xlsx", "pdf"}
var max_size int64 = 1024 * 1024 * 4

// Upload 通用上传
func (a *CommonController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	ext := filepath.Ext(file.Filename)
	if b, _ := array_tool.Inarray(upload_allow_ext, ext[1:]); !b {
		rsp.ErrMsg(c, "文件扩展名不在允许范围内")
		return
	}
	if file.Size > max_size {
		rsp.ErrMsg(c, "文件大小超出限制")
		return
	}
	file_name := helpers.RandomString(10) + ext
	path := app_file_path.PUBLIC_PATH + "/uploads/" + file_name
	err := c.SaveUploadedFile(file, path)
	if err != nil {
		rsp.ErrMsg(c, err.Error())
		return
	}
	rsp.RepData(c, gin.H{
		"upload_file_name": file.Filename,
		"file_name":        file_name,
		"path":             path,
		"ext":              ext,
	})
}

// GetCSRFToken 获取csrf_token
func (a CommonController) GetCSRFToken(c *gin.Context) {
	uid := auth.GetUID()
	if uid == "" {
		rsp.ErrMsg(c, "请先登陆")
		return
	}
	csrf_token.Set(c, uid)
	c.AbortWithStatus(http.StatusOK)
}
