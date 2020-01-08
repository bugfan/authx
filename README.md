## 精简的授权服务程序

### 功能
1. 提供注册 登录 退出 改密码等接口 
2. 使用cookie保存用户身份
3. JWT 保存会话的cookie value采用JWT加解密（代码tuils集成了JWT相关函数）
4. 数据库采用mysql 相关配置从settings里面查看
5. 用户密码经过加密存储
6. 默认使用sqlite3数据库,无其他依赖

### 使用方法
1. 执行 
```
    go get -u github.com/bugfan/authx
``` 
2. 单独使用: 切换到此目录执行 `go test` ,如果需要为此进程设置环境变量需要在命令行加入,或export出来;window用户配置类似;或者直接`AUTHX_PORT=8080 go test`,则把默认的端口9993改成了8080

3. 插件使用: 在你的代码里面引用此库，例如:
```
    import (
	    "github.com/bugfan/authx"
    )
    func main() {
        authx.Run("127.0.0.1:8080") //运行authx 
    }
```

### 环境变量
```
"db_obj":      "sqlite3",
"db_user":     "root",
"db_password": "",
"db_host":     "127.0.0.1:3306",
"db_name":     "authx",
"db_log":      "xorm.log",
"authx_host":  "localhost",
"authx_port":  "9993",
"jwt_secret":  "", 
```
#### 注意事项
1. 默认使用sqlite数据库,`authx.db`文件生成到同级目录,如果需要使用mysql,环境变量需要`export DB_OBJ=mysql`
2. 使用此库需要配置以上环境变量（左边的）,需要大写 例如 `export DB_USER=root` ,如果不配置则使用以上默认值,
3. 如果JWT加解密密钥不设置,即jwt_secret为空则使用随机字符串当作JWT加解密密钥