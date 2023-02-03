### 简介
- pgk为开发中常用包,每个包前加g用于防止和原生插件包重复
- config为配置文件，定义包参数

### 目录结构
```html
- e
- extend
- gcaptcha
- gcasbin
- gconf
- gcrypto
- gdb
- gjwt
- glog
- gredis
- gsession
- gsys**
- gvalidator
- utils
```

go test -v -run TestDb go_test.go
go clean -testcache


### gconf-读取配置包
- 采用viper插件包，目前自动读取根目录下config/config.yaml文件
- 使用方法
  - gconf.Config.UnmarshalKey()

### glog-日志
- 采用logrus插件包 目前只支持file储存，可扩展
- 参考网址
  - github地址 https://github.com/sirupsen/logrus
- 使用方法
  - glog.Get() 调用
  
### gdb-数据库
- 采用grom插件包  目前支持数据库[mysql,postgre]
- 参考网址
  - 官网地址 https://gorm.io/zh_CN/docs/create.html
- 使用方法
  - gdb.Get() 调用
- 注意
  - 日志输出依赖glog包
  
### gcasbin-权限管理
- 采用casbin插件包
- 参考网址
  - 官网地址 https://casbin.org/docs/zh-CN/overview
- 使用方法
  - gcasbin.Get() 调用
- 注意
  - 依赖数据库gdb
  
### gjwt-无状态登录
- 采用jwt-go插件包
- 参考网址
  - github https://github.com/golang-jwt/jwt
- 使用方法
  - gjwt.CreateToken(jwt.MapClaims{"123":123})
  - gjwt.ParseToken(token)
  
### gcrypto-加密
- 采用crypto插件包，支持[md5,sha256]加密

### gcaptcha-验证码


  

