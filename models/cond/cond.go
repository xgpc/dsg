// Package cond
// @Author:        asus
// @Description:   $
// @File:          cond
// @Data:          2022/4/129:54
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
func PageByQuery(ctx iris.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := ctx.URLParamIntDefault("page", 1)

		pageSize := ctx.URLParamIntDefault("page_size", 10)
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
		page := ctx.Params().GetIntDefault("page", 1)

		pageSize := ctx.Params().GetIntDefault("page_size", 10)
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

// DefinedOr 自定义gorm方法，or条件并列查询。(实现where条件下的查询 where ( A or B )方法。)
func DefinedOr(column1, column2 string, params1, params2 interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column1+"= ? or "+column2+"= ?", params1, params2)
	}
}

// LikeOr 自定义gorm方法，or条件中并列LIKE查询。（实现where条件下的查询 where ( A like a or B like b)方法。 )
func LikeOr(column1, column2, params1, params2 string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column1+" LIKE (?) or "+column2+" LIKE (?)", "%"+params1+"%", "%"+params2+"%")
	}
}

// Between 范围查询
func Between(column string, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(column+" BETWEEN ? AND ?", args)
	}
}
