/**
 * @Author: smono
 * @Description:
 * @File:  hand
 * @Version: 1.0.0
 * @Date: 2022/9/25 12:47
 */

package admin

import "gorm.io/gorm"

var (
	_db *gorm.DB
)

func Init(db *gorm.DB) {
	_db = db
}

func db() *gorm.DB {
	return _db
}
