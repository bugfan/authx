## 精简的授权服务程序

### 功能
1. 提供注册 登录 退出 改密码等接口 
2. 使用cookie保存用户身份
3. JWT 保存会话的cookie value采用JWT加解密（代码tuils集成了JWT相关函数）
4. 数据库采用mysql 相关配置从settings里面查看
5. 用户密码经过加密存储

### 使用方法
1. 切换到此目录执行 `go test` 
2. 执行 `go get -u github.com/bugfan/authx` 在你的代码里面引用此库，例如:
`
    import (
	    "github.com/bugfan/authx"
    )
    func main() {
        authx.Run() //运行authx 
    }
`

### 环境变量
`
        "db_user":     "root",
		"db_password": "123456",
		"db_host":     "127.0.0.1:3306",
		"db_name":     "authx",
		"db_log":      "xorm.log",
		"authx_host":  "localhost",
		"authx_port":  "9993",
`
使用此库需要配置以上环境变量（左边的）,需要大写 例如 `export DB_USER=root` ,如果不配值则使用以上默认值
