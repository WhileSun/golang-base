package service

import (
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"gorm.io/gorm"
)

type UserRoleService struct {
}

func NewUserRole() *UserRoleService {
	return &UserRoleService{}
}

func (s *UserRoleService) Add(userId int,roleIds []string,tx *gorm.DB) bool{
	userRoles := make([]*models.RUserRole,0)
	for _,roleId := range roleIds{
		userRole := &models.RUserRole{}
		userRole.UserId = userId
		userRole.RoleId = gconvert.StrToInt(roleId)
		if userRole.RoleId !=0{
			userRoles = append(userRoles, userRole)
		}
	}
	err := tx.Create(&userRoles).Error
	if err != nil {
		gsys.Logger.Error("添加用户角色失败—>", err.Error())
		return false
	}
	return true
}

func (s *UserRoleService) Update(userId int,roleIds []string,tx *gorm.DB) bool{
	err := tx.Where("user_id = ?",userId).Delete(&models.RUserRole{}).Error
	if err !=nil{
		gsys.Logger.Error("修改用户角色失败—>", err.Error())
		return false
	}
	return NewUserRole().Add(userId,roleIds,tx)
}

