### Go Web开发较通用的脚手架模板

go run main.go 成功执行后

访问swagger接口文档：http://localhost:8000/swagger/index.html#/

项目结构中: DAO 的含义是 Data Access Objects

#### 压力测试

1. 访问：http://127.0.0.1:8080/debug/pprof/

2. 启动压力测试：`go-wrk -n 50000 http://127.0.0.1:8080/version`

3. 点击 web 页面的 profile 按钮，收集信息，30 s 后会自动下载 profile 文件

4. 拿到 profile 文件后：`go tool pprof -http=":8081" profile`