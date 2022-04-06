# 结构目录
gin
- config 配置文件
- middware 中间项
- models 模型数据库交互
- pkg 工具类
- routers 路由
- runtime 运行日志
- service 服务层 逻辑抒写
- controller 控制器
- main.go 入口文件

# sql 转化成struct工具
- https://github.com/WhileSun/db2go
- pgsqltrans.exe -host=127.0.0.1 -port=5432 -uname=postgres -pwd=lingweiqu -dbname=go_oms

# swagger  生成API文档
- go get -u github.com/swaggo/swag/cmd/swag
- swag init
