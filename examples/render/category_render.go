// Package render
// @Author:        asus
// @Description:   $
// @File:          category_render
// @Data:          2022/2/2816:48
//
package render

import (
	"gorm.io/gorm"
)

type Category struct {
	ID           uint           `json:"id,omitempty" example:"1"`
	CreatedAt    uint32         `json:"createdAt,omitempty" example:"1646036184"`
	UpdatedAt    uint32         `json:"updatedAt,omitempty" example:"1646036184"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt,omitempty" swaggertype:"string" example:"2022-02-01 16:51:21"`
	Path         string         `json:"path,omitempty" example:"-"`
	CategoryName string         `json:"categoryName,omitempty" example:"一级分类"`
	CategoryImg  string         `json:"categoryImg,omitempty" example:"http://test/image/1.jpg"`
	IsFinal      bool           `json:"isFinal,omitempty" example:"false"`
	Sort         uint           `json:"sort,omitempty" example:"0"`
	CategoryID   uint           `json:"categoryID,omitempty" example:"0"`
	Level        uint           `json:"level,omitempty" example:"1"`
}
