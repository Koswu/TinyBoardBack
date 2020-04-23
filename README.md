# TinyBoardBack
![tinyboard](https://img.shields.io/badge/tinyboard-back-blue) ![license](https://img.shields.io/github/license/Koswu/TinyBoardBack)  ![last-commit](https://img.shields.io/github/last-commit/Koswu/TinyBoardFront)

一个使用gin框架写的留言板系统的后端


## 使用方法

下载后修改conf/app.ini，
* 修改JWT_SECRET的值为随机的字符串，
* 修改IS_DEBUG=false
* 修改HTTP端口
* 修改数据库类型和连接信息

最后运行

```shell
go run main.go
```



最后在[这里](https://github.com/Koswu/TinyBoardFront)下载前端部署就完成了，请访问看看吧~