# 结构目录
gin
```
|—app               应用目录
    |—controller    控制器
    |—dto           数据传输对象
    |—models        模型目录
    |—po            数据库表对象
    |—service       服务层 逻辑抒写
    |—vo            输出对象
|—config            配置文件
|—middware          中间项
    |—gsys          系统公用调用插件
    |—utils         普通插件
|—pkg               工具类
|—routers           路由
|—runtime           运行日志
|—main.go           入口文件
```

# sql 转化成struct工具
- https://github.com/WhileSun/db2go
- pgsqltrans.exe -host=127.0.0.1 -port=5432 -uname=postgres -pwd=lingweiqu -dbname=go_oms

# swagger  生成API文档
- go get -u github.com/swaggo/swag/cmd/swag
- swag init
