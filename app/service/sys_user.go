package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/types/dto"
	"github.com/whilesun/go-admin/gctx"
	"github.com/whilesun/go-admin/pkg/utils/e"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"github.com/whilesun/go-admin/pkg/utils/gcrypto"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
	"gorm.io/gorm"
)

type SysUserService struct {
}

var SysUserServiceApp = new(SysUserService)

func (s *SysUserService) checkUsernameExist(username string) error {
	if id := models.NewUser().CheckUsernameExist(username); id > 0 {
		return errors.New(fmt.Sprintf("用户账号[%s]已经存在，请更换！", username))
	}
	return nil
}

func (s *SysUserService) checkRealnameExist(realname string) error {
	if id := models.NewUser().CheckRealnameExist(realname); id > 0 {
		return errors.New(fmt.Sprintf("用户名称[%s]已经存在，请更换！", realname))
	}
	return nil
}

// Add 添加用户
func (s *SysUserService) Add(params dto.AddUser) error {
	err := gctx.Db.Transaction(func(tx *gorm.DB) error {
		userModel := models.NewUser()
		gconvert.StructCopy(params, userModel)
		if err := SysUserServiceApp.checkUsernameExist(userModel.Username); err != nil {
			return err
		}
		if err := SysUserServiceApp.checkRealnameExist(userModel.Realname); err != nil {
			return err
		}
		initPwd := gctx.GConfig.GetString("app.initPwd")
		userModel.Password = gcrypto.PwdEncode(gtools.StringDefault(initPwd, "a123456!"), gctx.GConfig.GetString("app.saltPwd"))
		err := tx.Create(userModel).Error
		if err != nil && userModel.Id == 0 {
			gctx.Logger.Error("添加用户失败—>", err.Error())
			return errors.New("添加用户失败！")
		}
		//添加用户角色关联
		ok := SysUserServiceApp.addUserRole(userModel.Id, params.RoleIds, tx)
		if !ok {
			return errors.New("添加用户角色失败！")
		}
		return nil
	})
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// addUserRole 用户和角色管理表添加
// @param roleIds 角色ID可以多个
func (s *SysUserService) addUserRole(userId int, roleIds []string, tx *gorm.DB) bool {
	userRoles := make([]*models.RUserRole, 0)
	for _, roleId := range roleIds {
		userRole := &models.RUserRole{}
		userRole.UserId = userId
		userRole.RoleId = gconvert.StrToInt(roleId)
		if userRole.RoleId != 0 {
			userRoles = append(userRoles, userRole)
		}
	}
	err := tx.Create(&userRoles).Error
	if err != nil {
		gctx.Logger.Error("添加用户角色失败—>", err.Error())
		return false
	}
	return true
}

// Update 修改用户
func (s *SysUserService) Update(params dto.UpdateUser) error {
	err := gctx.Db.Transaction(func(tx *gorm.DB) error {
		//已入库信息
		oldUserModel := models.NewUser()
		oldUserModel.GetInfo(params.Id)
		if oldUserModel.Id == 0 {
			return errors.New("更新的用户不存在，请确认！")
		}
		superName := gtools.StringDefault(gctx.GConfig.GetString("app.userSuperName"), "system")
		if oldUserModel.Username == superName {
			return errors.New("超级用户信息不支持修改！")
		}
		//生成参数model
		userModel := models.NewUser()
		gconvert.StructCopy(params, userModel)
		if oldUserModel.Username != userModel.Username {
			if err := SysUserServiceApp.checkUsernameExist(userModel.Username); err != nil {
				return err
			}
		}
		if oldUserModel.Realname != userModel.Realname {
			if err := SysUserServiceApp.checkRealnameExist(userModel.Realname); err != nil {
				return err
			}
		}
		err := tx.Model(userModel).Select("username", "realname", "status").Updates(userModel).Error
		if err != nil {
			gctx.Logger.Error("修改用户失败—>", err.Error())
			return errors.New("修改用户失败！")
		}
		oldRoleIds := models.NewUserRole().GetUserRoleIds(oldUserModel.Id)
		equalBool := gtools.StrArrayEquals(oldRoleIds, params.RoleIds, true)
		//修改角色
		if equalBool == false {
			//修改用户角色
			ok := SysUserServiceApp.updateUserRole(userModel.Id, params.RoleIds, tx)
			if !ok {
				return errors.New("修改用户角色失败！")
			}
		}
		return nil
	})
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// updateUserRole 更新用户和角色的关联表
func (s *SysUserService) updateUserRole(userId int, roleIds []string, tx *gorm.DB) bool {
	err := tx.Where("user_id = ?", userId).Delete(&models.RUserRole{}).Error
	if err != nil {
		gctx.Logger.Error("修改用户角色失败—>", err.Error())
		return false
	}
	//重新插入
	return SysUserServiceApp.addUserRole(userId, roleIds, tx)
}

// CheckLogin 检测用户登录
func (s *SysUserService) CheckLogin(params dto.LoginUser) (string, uint) {
	if gctx.GCaptcha.Verify(params.CaptchaId, params.CaptchaCode) == false {
		return "", e.ERROR_CAPTCHA_VERIFY
	}
	pwd := gcrypto.PwdEncode(params.Password, gctx.GConfig.GetString("app.saltPwd"))
	userModel := models.NewUser().GetLoginInfo(params.Username, pwd)
	if userModel.Id == 0 {
		gctx.Logger.Errorf("用户[%s]登录失败 -> 账号密码错误", params.Username)
		return "", e.ERROR_ACCOUNT_LOGIN
	}
	//用户禁止登录
	if userModel.Status == false {
		return "", e.ERROR_ACCOUNT_CLOSE
	}
	token, err := SysUserAuthServiceApp.CreateJwtToken(userModel.Id)
	if err != nil {
		return "", e.FAILED
	}
	//NewUserAuth().SetSession(userModel, sessionKey)
	return token, e.SUCCESS
}

// UpdatePasswd 更改密码
func (s *SysUserService) UpdatePasswd(params dto.UpdateUserPasswd, req *gin.Context) error {
	ok := gtools.VerifyPasswdV4(params.NewPasswd)
	if ok {
		userId := req.GetInt("userId")
		userModel := models.NewUser()
		oldPasswd := userModel.GetPasswd(userId)
		if gcrypto.PwdEncode(params.OldPasswd, gctx.GConfig.GetString("app.saltPwd")) != oldPasswd {
			return errors.New("旧密码不正确，请重新输入")
		}
		userModel.Id = userId
		userModel.Password = gcrypto.PwdEncode(params.NewPasswd, gctx.GConfig.GetString("app.saltPwd"))
		err := userModel.UpdatePasswd()
		if err != nil {
			gctx.Logger.Error("更新用户新密码失败", err.Error())
			return errors.New("更新用户新密码失败")
		}
	} else {
		return errors.New("密码需要符合规则[大小写字母+数字+特殊符号+6位以上]")
	}
	return nil
}
