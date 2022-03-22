package category

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

// CategoryControllerPost 创建分类
type CreateCategoryRequest struct {
	CategoryName string `validate:"required,min=1,max=20" label:"分类名称" json:"categoryName"` // 分类名称
	CategoryImg  string `validate:"-" label:"分类图片" json:"categoryImg" default:""`           // 分类图片地址链接
	CategoryID   uint   `validate:"-" label:"所属分类" json:"categoryID" default:"0"`           // 所属分类
	Sort         uint   `validate:"min=0,max=100" label:"排序" json:"sort" default:"0"`       //排序
	IsFinal      bool   `validate:"-" label:"是否为最终类" json:"isFinal"  default:"false"`       // 是否为最终类
}

// CreateCategory 添加商品分类
// @Summary 添加商品分类
// @Description 后台管理人员添加商品分类
// @Accept json
// @Produce json
// @param root body CreateCategoryRequest true "添加商品分类"
// @Tags 商品分类
// @Success 200 {object} render.Response{data=render.Category}
// @Router /api/category [post]
func CreateCategory(ctx iris.Context) {
	param := &CreateCategoryRequest{}

	this := frame.NewBase(ctx)
	this.Init(param)
}

// GetList 根据分类id获取子集分类
// @Summary 获取子分类列表
// @Description 根据分类id获取子分类列表
// @Produce json
// @param categoryID query uint false "分类id" default(0)
// @param page query uint false "分页" default(1)
// @param pageSize query uint false "分页页数" default(10)
// @Tags 商品分类
// @Success 200 {object} render.Response{data=[]render.Category}
// @Router /api/category [get]
func GetList(ctx iris.Context) {

}
