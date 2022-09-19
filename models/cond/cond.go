// Package cond
// @Author:        asus
// @Description:   $
// @File:          cond
// @Data:          2022/4/129:54
//
package cond

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

// Paginate 分页
func Page(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := page
		pageSize := pageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// PagebyURL 分页
func PagebyURL(ctx iris.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := ctx.URLParamIntDefault("Page", 1)

		pageSize := ctx.URLParamIntDefault("PageSize", 10)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// PageByParams 分页
func PageByParams(ctx iris.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := ctx.Params().GetIntDefault("Page", 1)

		pageSize := ctx.Params().GetIntDefault("PageSize", 10)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Eq(column string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" = (?)", args)
	}

}

func NotEq(column string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" <> (?)", args)
	}

}

func Gt(column string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" > (?)", args)
	}

}

func Gte(column string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" >= (?)", args)
	}

}

func Lt(column string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" < (?)", args)
	}

}

func Lte(column string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" <= (?)", args)
	}

}

func Like(column string, str string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" LIKE (?)", "%"+str+"%")
	}
}

func Starting(column string, str string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" LIKE (?)", str+"%")
	}
}

func Ending(column string, str string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" LIKE (?)", "%"+str)
	}
}

func In(column string, params interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" in (?) ", params)

	}
}

func NotIn(column string, params interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" not in (?) ", params)
	}
}
