package models

type DbInit struct {

}

func NewDbInit() *DbInit{
	return &DbInit{}
}

func (m *DbInit) InitSuper(){
	//db.Transaction(func(tx *gorm.DB) error{
	//	role := &SRole{RoleName: "超级管理员",RoleSub: "super_admin",Sort: 0,Status: true}
	//	if err:= db.Create(&role).Error; err != nil{
	//		return err
	//	}
	//	user := &SUser{Username: "system",Password: gcrypto.PwdEncode("a123456!"),Realname: "超级用户",Status: true}
	//	if err := db.Create(&user).Error; err !=nil{
	//		return err
	//	}
	//	//添加rbac中的权限
	//	NewRolePolicy().AddUserRolesPolicy("system",[]string{"super_admin"})
	//	policys := make([]map[string]interface{},0)
	//	policys = append(policys,map[string]interface{}{"data_perms":"*"})
	//	NewRolePolicy().AddRolePolicy("super_admin",policys)
	//	return nil
	//})
}