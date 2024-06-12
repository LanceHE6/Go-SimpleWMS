<p align="center">
<img alt="logo" src="../imgs/WMS-Logo.png">
</p>
<h1 align="center" style="margin: 30px 0 30px; font-weight: bold;">Go-SimpleWMS</h1>
<div align="center"> 
</div>

<h4 align="center">基于gin+Vue3前后端分离式仓库管理系统</h4>

<h3 align="center">后端文档</h3>

<div align="center">

![Static Badge](https://img.shields.io/badge/Licence-MIT-blue)
![Static Badge](https://img.shields.io/badge/前端-vue-orange)
![Static Badge](https://img.shields.io/badge/后端-gin-green)

</div>

* **[接口文档](api.md)**



# 后端运行

```bash
# 进入项目目录
cd go-server

# 下载依赖
go mod download

# 构建项目
go build -o main .
```

----

# 项目结构

```
Go_SimpleWMS\go-server
├─config.yaml 				// 配置文件
├─go.mod
├─go.sum
├─main.go 					// 程序入口文件
├─utils						// 工具方法文件夹
├─route						//路由文件夹
|   ├─route.go 				// 总路由
|   ├─group					// 路由组
├─logs						// 日志文件夹
├─static					// 静态资源文件夹
├─handler					// 处理器文件夹
├─database					// 数据库相关文件夹
|    ├─myDb					// 数据库连接及数据初始化文件夹
|    ├─model				// 数据表模型定义文件夹
├─config					// 配置加载文件夹
```

----

# 配置文件说明

config.yaml

```yaml
server:
    port: "8080" 			# 服务器监听端口
    mode: "debug" 			# 日志输出模式 debug/release
    log:
        path: "logs"		# 日志保存路径
        max_files: 15		# 最大保留日志文件数
    secretKey: "simple_wms" # 用于token签名的密钥
db:
  mysql:
    host: "localhost"		# mysql数据库url
    port: "3306"			# 数据库端口
    account: "root"			# 账号
    password: "root"		# 密码
    dbname: "simple_wms"	# 数据库名称
```

----



# 数据库建表示例

## 数据库建模

![image-20240606163507642](..\imgs\DBModel.png)

## users表

| 字段 | id             |          uid           |   account    |   password   |  nick_name   |   permission    | created_at | updated_at |     token      |    phone     |
| :--: | -------------- | :--------------------: | :----------: | :----------: | :----------: | :-------------: | :--------: | ---------- | :------------: | :----------: |
| 类型 | int            |      varchar(20)       | varchar(255) | varchar(255) | varchar(255) |       int       |  datetime  | datetime   |  varchar(255)  | varchar(255) |
| 说明 | 数据库内置索引 | 标识+8位唯一索引(主键) |     账号     |     密码     |     昵称     | 权限（1，2，3） |  创建时间  | 更新时间   | 登录生成的凭证 |     电话     |

----

## warehouses表

| 字段 | id             |          wid           |     name     | created_at | updated_at |   comment    |     manager     |  status  |
| :--: | -------------- | :--------------------: | :----------: | :--------: | ---------- | :----------: | :-------------: | :------: |
| 类型 | int            |      varchar(20)       | varchar(255) |  datetime  | datetime   | varchar(255) |  varchar(255)   |   int    |
| 说明 | 数据库内置索引 | 标识+8位唯一索引(主键) |    仓库名    |  创建时间  | 更新时间   |     备注     | 负责人sid(外键) | 仓库状态 |

----

## goods_types表

| 字段 | id             |          gtid          |     name     |  type_code   | created_at | updated_at |
| :--: | -------------- | :--------------------: | :----------: | :----------: | :--------: | ---------- |
| 类型 | int            |      varchar(20)       | varchar(255) | varchar(255) |  datetime  | datetime   |
| 说明 | 数据库内置索引 | 标识+8位唯一索引(主键) |  货品类型名  | 货品类型编码 |  创建时间  | 更新时间   |

----

## departments表

| 字段 | id             |          did           |     name     | created_at | updated_at |
| :--: | -------------- | :--------------------: | :----------: | :--------: | ---------- |
| 类型 | int            |      varchar(20)       | varchar(255) |  datetime  | datetime   |
| 说明 | 数据库内置索引 | 标识+8位唯一索引(主键) |   部门名称   |  创建时间  | datetime   |

----

## staffs表

| 字段 | id             |          sid           |     name     |    phone     |      department      | created_at | updated_at |
| :--: | -------------- | :--------------------: | :----------: | :----------: | :------------------: | :--------: | ---------- |
| 类型 | int            |      varchar(20)       | varchar(255) | varchar(255) |     varchar(255)     |  datetime  | datetime   |
| 说明 | 数据库内置索引 | 标识+8位唯一索引(主键) |   员工名称   |     电话     | 员工所属部门id(外键) |  创建时间  | 更新时间   |

----

## inventory_types表

| 字段 | id             |          itid          |      name      |   type_code    | created_at | updated_at |
| :--: | -------------- | :--------------------: | :------------: | :------------: | :--------: | ---------- |
| 类型 | int            |      varchar(20)       |  varchar(255)  |  varchar(255)  |  datetime  | datetime   |
| 说明 | 数据库内置索引 | 标识+8位唯一索引(主键) | 出入库类型名称 | 出入库类型编码 |  创建时间  | 更新时间   |

## goods表

| 字段 | id             |       gid        |  goods_code  |     name     |    model     |   goods_type   | manufacturer |      unit      | created_at | updated_at |
| :--: | -------------- | :--------------: | :----------: | :----------: | :----------: | :------------: | :----------: | :------------: | ---------- | ---------- |
| 类型 | int            |   varchar(20)    | varchar(255) | varchar(255) | varchar(255) |  varchar(255)  | varchar(255) |  varchar(255)  | datetime   | datetime   |
| 说明 | 数据库内置索引 | 标识+8位唯一索引 |   货品编码   |   货品名称   |   型号规格   | 货品类型(外键) |    生产商    | 计量单位(外键) | 创建时间   | 更新时间   |

## units表

| 字段 | id             |       unid       |     name     | created_at | updated_at |
| :--: | -------------- | :--------------: | :----------: | ---------- | ---------- |
| 类型 | int            |   varchar(20)    | varchar(255) | datetime   | datetime   |
| 说明 | 数据库内置索引 | 表示+8位唯一索引 |   单位名称   | 创建时间   | 更新时间   |

----

## inventories表

| 字段 | id             |          iid           |    number    | goods_list |   inventory_type   | amount |    warehouse     | operator         | comment      | manufacturer | created_at | updated_at |
| :--: | -------------- | :--------------------: | :----------: | ---------- | :----------------: | ------ | :--------------: | ---------------- | ------------ | ------------ | ---------- | ---------- |
| 类型 | int            |      varchar(20)       | varchar(255) | json       |    varchar(255)    | int    |   varchar(255)   | varchar(255)     | varchar(255) | varchar(255) | datetime   | datetime   |
| 说明 | 数据库内置索引 | 标识+8位唯一索引(主键) |    订单号    | 货品列表   | 出入库类型id(外键) | 数量   | 所属仓库id(外键) | 操作员工id(外键) | 备注         | 生产商       |            |            |

## stocks表

| 字段 | id             |  warehouse   |    goods     | quantity | created_at | updated_at |
| :--: | -------------- | :----------: | :----------: | :------: | :--------: | :--------: |
| 类型 | int            | varchar(255) | varchar(255) |  float   |  datetime  |  datetime  |
| 说明 | 数据库内置索引 | 仓库id(外键) | 货品id(外键) |   数量   |  创建时间  |  更新时间  |
