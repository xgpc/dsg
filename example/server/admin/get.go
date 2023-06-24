/**
 * @Author: smono
 * @Description:
 * @File:  get
 * @Version: 1.0.0
 * @Date: 2022/9/25 12:49
 */

package admin

import "company/admin/models"

func getUser(userID, SubjectID uint32) models.ShoppingUser {
	var info models.ShoppingUser
	db().Model(&info).
		Where(models.ShoppingUserColumns.UserID, userID).
		Where(models.ShoppingUserColumns.SubjectID, SubjectID).
		First(&info)
	return info
}
