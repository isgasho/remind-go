## 一款由微信公众号引发的指定时间消息提醒事件


### 对此项目介绍

[一段骚操作 我又搞出了啥](https://learnku.com/articles/45034)

### 环境
**go环境自行安装**

**项目未使用框架**

### 增加配置
根目录下增加 config.json 文件 
```json
{
  "Wechat": {
    "AppID": "xxx",
    "AppSecret": "xxxx",
    "Token": "xxx",
    "EncodingAESKey": "xxxx"
  },
  "Db": {
    "Driver": "mysql",
    "Address": "192.168.10.10:3306",
    "Database": "remind",
    "User": "homestead",
    "Password": "secret"
  },
  "Email": {
    "User": "1185079673@qq.com",
    "Pass": "xxx",
    "Host": "smtp.qq.com",
    "Port": 25
  },
  "Teng": {
    "SECRETID": "xxx",
    "SecretKey": "xxx",
    "SDKAppID": "xxx",
    "TemplateID": "xxx"
  }
}
```

**数据库**

**数据库文件放在根目录下**



### 运行
**在项目根目录下运行**
```go
go run main.go
```

**或者根目录下已将应用程序打包成二进制可执行文件,可以在根目录下直接执行**
```go
./remind-go
```






