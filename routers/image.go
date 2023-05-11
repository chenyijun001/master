package routers

import (
	"GinBlog/api"
)

var image = api.App.ImageApi

func (r RouterGroup) ImageRouter() {
	r.POST("imageUpload", image.ImageUploadView)
	r.GET("imageUpload", image.ImageListView)
	r.GET("imageUpload/list", image.ImageNameListView)
	r.DELETE("imageUpload/delete", image.DeleteImageByIdView)
	r.PUT("imageUpload/rename", image.ImageRenameView)
}
