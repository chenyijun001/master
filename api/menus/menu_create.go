package menus

import (
	"GinBlog/global"
	"GinBlog/models"
	"GinBlog/models/ctype"
	"GinBlog/models/res"
	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	MenuTitle     string      `json:"menu_title" binding:"required" msg:"请完善菜单名称"`        //中文标题
	MenuTitleEn   string      `json:"menu_title_en" binding:"required" msg:"请完善菜单名称英文"` //英文标题
	Slogan        string      `json:"slogan"`                                                    //标语
	Abstract      ctype.Array `json:"abstract"`                                                  //简介
	AbstractTime  int         `json:"abstract_time"`                                             //简介的切换时间
	BannerTime    int         `json:"banner_time"`                                               //菜单的切换时间
	Sort          int         `json:"sort" binding:"required" msg:"请完善菜单顺序"`              //菜单的顺序
	ImageSortList []ImageSort `json:"image_sort_list"`                                           //具体的图片顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var mr MenuRequest
	err := c.ShouldBindJSON(&mr)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	//判断重复值

	//创建banner入库
	model := models.MenuModel{
		MenuTitle:    mr.MenuTitle,
		MenuTitleEn:  mr.MenuTitleEn,
		Slogan:       mr.Slogan,
		Abstract:     mr.Abstract,
		AbstractTime: mr.AbstractTime,
		BannerTime:   mr.BannerTime,
		Sort:         mr.Sort,
	}
	err = global.DB.Create(&model).Error
	if err != nil {
		global.Logs.Error("创建banner入库错误，", err)
		res.FailWithMessage("菜单保存错误", c)
		return
	}

	//给关联表入库
	if len(mr.ImageSortList) == 0 {
		res.OKWithMessage("菜单保存成功", c)
		return
	}

	var menuBannerList []models.MenuBannerModel

	for _, sort := range mr.ImageSortList {
		//判断图片ID
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   model.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}

	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Logs.Error(err)
		res.FailWithMessage("菜单保存失败", c)
		return
	}

}
