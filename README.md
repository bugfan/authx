## 精简的授权服务程序

### 功能
1. 提供注册 登录 退出 改密码等接口 
2. 使用cookie保存用户身份
3. JWT 保存会话的cookie value采用JWT加解密（代码tuils集成了JWT相关函数）
4. 数据库采用mysql 相关配置从settings里面查看
5. 用户密码经过加密存储

### 使用方法
 `go run main.go` or `当作库引进来调用mian里面的api即可`

## 特别说明 
1. 近期会一直更新，直到全部完成