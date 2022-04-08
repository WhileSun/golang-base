package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/e"
	"github.com/whilesun/go-admin/pkg/gcaptcha"
	"github.com/whilesun/go-admin/pkg/gconf"
	"github.com/whilesun/go-admin/pkg/gcrypto"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
	"gorm.io/gorm"
)

type UserService struct {
}

func NewUser() *UserService {
	return &UserService{}
}

func checkUsernameExist(userModel *models.SUser) error {
	if id := userModel.CheckUsernameExist(); id > 0 {
		return errors.New(fmt.Sprintf("用户账号[%s]已经存在，请更换！", userModel.Username))
	}
	return nil
}

func checkRealnameExist(userModel *models.SUser) error {
	if id := userModel.CheckRealnameExist(); id > 0 {
		return errors.New(fmt.Sprintf("用户名称[%s]已经存在，请更换！", userModel.Realname))
	}
	return nil
}

func (s *UserService) Add(params dto.AddUser) error {
	err := gsys.Db.Transaction(func(tx *gorm.DB) error {
		userModel := models.NewUser()
		gconvert.StructCopy(params, userModel)
		if err := checkUsernameExist(userModel); err != nil {
			return err
		}
		if err := checkRealnameExist(userModel); err != nil {
			return err
		}
		initPwd := gconf.Config.GetString("app.initPwd")
		userModel.Password = gcrypto.PwdEncode(gtools.StringDefault(initPwd, "a123456!"))
		err := tx.Create(userModel).Error
		if err != nil && userModel.Id == 0 {
			gsys.Logger.Error("添加用户失败—>", err.Error())
			return errors.New("添加用户失败！")
		}
		//添加用户角色关联
		ok := NewUserRole().Add(userModel.Id, params.RoleIds, tx)
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

func (s *UserService) Update(params dto.UpdateUser) error {
	err := gsys.Db.Transaction(func(tx *gorm.DB) error {
		//已入库信息
		oldUserModel := models.NewUser()
		oldUserModel.GetInfo(params.Id)
		if oldUserModel.Id == 0 {
			return errors.New("更新的用户不存在，请确认！")
		}
		superName := gtools.StringDefault(gconf.Config.GetString("app.userSuperName"), "system")
		if oldUserModel.Username == superName {
			return errors.New("超级用户信息不支持修改！")
		}
		//生成参数model
		userModel := models.NewUser()
		gconvert.StructCopy(params, userModel)
		if oldUserModel.Username != userModel.Username {
			if err := checkUsernameExist(userModel); err != nil {
				return err
			}
		}
		if oldUserModel.Realname != userModel.Realname {
			if err := checkRealnameExist(userModel); err != nil {
				return err
			}
		}
		err := tx.Model(userModel).Select("username", "realname", "status").Updates(userModel).Error
		if err != nil {
			gsys.Logger.Error("修改用户失败—>", err.Error())
			return errors.New("修改用户失败！")
		}
		//修改用户角色
		ok := NewUserRole().Update(userModel.Id, params.RoleIds, tx)
		if !ok {
			return errors.New("修改用户角色失败！")
		}
		return nil
	})
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

//CheckLogin 监测用户登录
func (s *UserService) CheckLogin(params *dto.LoginUser) (string, uint) {
	if gcaptcha.Verify(params.CaptchaId, params.Captcha) == false {
		return "", e.ERROR_CAPTCHA_VERIFY
	}
	pwd := gcrypto.PwdEncode(params.Password)
	user := models.NewUser().CheckExist(params.Username, pwd)
	if user.Id == 0 {
		gsys.Logger.Error("用户登录失败 -> 账号密码错误")
		return "", e.ERROR_ACCOUNT_LOGIN
	}
	//用户禁止登录
	if user.Status == false {
		return "", e.ERROR_ACCOUNT_CLOSE
	}
	token := NewUserAuth().CreateLoginToken(user.Id)
	NewUserAuth().SetSession(user, token)
	return token, e.SUCCESS
}

func (s *UserService) UpdatePasswd(params dto.UpdateUserPasswd,req *gin.Context) error{
	ok := gtools.VerifyPasswdV4(params.NewPasswd)
	if ok{
		userId := req.GetInt("userId")
		userToken := req.GetString("userToken")
		userModel := models.NewUser()
		oldPasswd := userModel.GetPasswd(userId)
		if gcrypto.PwdEncode(params.OldPasswd) != oldPasswd {
			return errors.New("旧密码不正确，请重新输入")
		}
		userModel.Id = userId
		userModel.Password = gcrypto.PwdEncode(params.NewPasswd)
		err:=userModel.UpdatePasswd()
		if err != nil{
			gsys.Logger.Error("更新用户新密码失败", err.Error())
			return errors.New("更新用户新密码失败")
		}
		NewUserAuth().DelSession(userId,userToken)
	}else{
		return errors.New("密码需要符合规则[大小写字母+数字+特殊符号+6位以上]")
	}
	return nil
}