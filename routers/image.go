package routers

import "GinBlog/api"

var image = api.App.ImageApi

func (r RouterGroup) ImageRouter() {
	r.POST("imageUpload/", image.ImageUploadView)
}
