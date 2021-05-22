package utils

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/zengqiang96/mattermostclone/model"
)

func RenderWebAppError(config *model.Config, w http.ResponseWriter, r *http.Request, err *model.AppError) {
	RenderWebError(config, w, r, err.StatusCode, url.Values{
		"message": []string{err.Message},
	})
	fmt.Printf("错误: %v", err)
}

func RenderWebError(config *model.Config, w http.ResponseWriter, r *http.Request, status int, params url.Values) {
}
