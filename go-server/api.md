# 后端接口文档



*测试环境后台地址：*

```
http://47.236.80.244:6007
https://simplewms.hycerlance.fun/api
```



----

## 状态码说明

| code |   2xx    |   4xx    |      5xx       |
| :--: | :------: | :------: | :------------: |
| 说明 | 请求成功 | 请求有误 | 服务端内部错误 |

## 业务码说明

| code |    1xx     |   2xx    |   4xx    |      5xx       |
| :--: | :--------: | :------: | :------: | :------------: |
| 说明 | 鉴权不通过 | 处理成功 | 处理失败 | 服务端内部错误 |

## 请求说明

**所有接口需携带Content-Type头**

示例：

```
headers:{
	"Content-Type": "application/json"
}
```

----

### 版本信息：

**请求路径**：/api/ping

**请求方法**：GET

**返回示例：**

```json
{
    "code": 200,
    "data": {
        "version": "v0.0.1.20240605_Alpha"
    },
    "msg": "Hello Go-SimpleWMS"
}
```



## 用户

### 鉴权：

请求头携带鉴权JWT token

**请求路径**：/api/auth

**请求方法**：GET

示例：

```
headers:{
	"Authorization": "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk1NTUxMjEsInN1YiI6IjAwMDAwMDAyIn0.hvz-Xp9kfhVNsCy6Q9nhS9wM8-c-DgJJ8PLcME17Fto"
}
```

### 参数说明

| 参数名        | 参数类型 | 是否必填 | 参数说明                                                     |
| ------------- | -------- | -------- | ------------------------------------------------------------ |
| Authorization | string   | 是       | 值为以bearer开头并以空格分隔带有成功登陆后返回的token的长字符串 |

### 鉴权相关返回示例

```json
// 无token 401
{
    "code": 101,
	"msg": "No Authorization header provided",
    "data": null
}
// 无效token 401
{
    "msg": "Invalid token",
	"code":    102,
    "data": null
}
// 鉴权通过 200
{
    "code": 201,
    "msg": "Hello",
	"data":  {
    	"uid": uid
	},
}
```

### 鉴权相关返回数据说明

| 参数名 | 参数类型 |       参数说明       |
| :----: | :------: | :------------------: |
|  code  |   int    |        业务码        |
| error  |  string  | 后端服务内部错误消息 |
|  msg   |  string  |     请求返回消息     |

### 鉴权相关返回状态码说明

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  401   | StatusUnauthorized  |    鉴权未通过    |
|  403   |   StatusForbidden   |     权限拒绝     |
|  500   | InternalServerError | 后端服务内部错误 |



### 注册

**请求路径**：/api/user/register

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名        | 参数类型 | 是否必填 | 参数说明     |
|------------| -------- | -------- | ------------ |
| account    | String   | 是       | 用户名（手机号/邮箱） |
| password   | String   | 是       | 密码         |
| permission | int | 是 | 权限大小（1为用户，2为管理员，3为超级管理员） |
| nickname  | string | 是 | 昵称 |
| phone | string | 否 | 电话 |

**返回结果示例**：

```json
// 注册成功 200
{
    "code": 201,
    "msg": "User registered successfully",
    "data": {
        "uid": uid
    }
}
```

**返回数据说明**

| 参数名 | 参数类型 |       参数说明       |
| :----: | :------: | :------------------: |
|  code  |   int    |        业务码        |
|  msg   |  string  |       返回消息       |
|  uid   |  string  | 注册成功返回的用户id |
| error  |  string  |   后端内部错误消息   |
| detail |  string  |       错误详情       |



----

### 批量注册（导入）

**请求路径**：/api/user/upload

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型       | 是否必填 | 参数说明 |
| ------ | -------------- | -------- | -------- |
| list   | Array <Object> | 是       | 用户列表 |

**列表参数**

| 参数名     | 参数类型 | 是否必填 | 参数说明                                      |
| ---------- | -------- | -------- | --------------------------------------------- |
| account    | String   | 是       | 用户名（手机号/邮箱）                         |
| password   | String   | 是       | 密码                                          |
| permission | int      | 是       | 权限大小（1为用户，2为管理员，3为超级管理员） |
| nickname   | string   | 是       | 昵称                                          |
| phone      | string   | 否       | 电话                                          |

**返回结果示例**：

```json
// 部分导入失败 217
{
    "code": 403,
    "msg": "Some users failed to register; 1/2",
    "data": {
    	"detail": [
            {
                "code": 402,
                "msg": "The account '87' has been registered"
            }
    ],
    }
}
// 全部导入失败 400
{
    "code": 402,
    "msg": "All users failed to register",
    "data": {
         "detail": [
             {
                 "code": 402,
                 "msg": "The account '87' has been registered"
             },
             {
                 "code": 402,
                 "msg": "The account '81' has been registered"
             }
    ],
    }
}
// 全部导入成功 200
{
    "code":    201,
	"msg": "All users have completed registration",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |       参数说明       |
| :----: | :------: | :------------------: |
|  code  |   int    |        业务码        |
|  msg   |  string  |       返回消息       |
|  uid   |  string  | 注册成功返回的用户id |
| error  |  string  |   后端内部错误消息   |
| detail |  string  |   每个错误注册详情   |

----



### 登录

**请求路径**：/api/user/login

**请求方法**：POST

**是否需要鉴权：**否

**请求参数**：

| 参数名  | 参数类型 | 是否必填 | 参数说明     |
| :-------: | :------: | :------: | :----------: |
| username  | String   | 是       | 用户名（手机号/邮箱） |
| password  | String   | 是       | 密码         |

**返回结果示例**：

```json
// 登录成功 200
{
    "code": 201,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTgzNTU5ODcsInVpZCI6InUwMDAwMDAwMSIsInBlcm1pc3Npb24iOjMsImNyZWF0ZWQtYXQiOiIyMDI0LTA2LTExIDE1OjUzOjMzICswODAwIENTVCJ9.keeh5VhCwKDQmGph4KgkS0A4o5rZjwzy1V-ILea89eY",
        "user": {
            "created_at": "2024-06-11T15:53:33+08:00",
            "updated_at": "2024-06-11T17:06:27.8157393+08:00",
            "uid": "u00000001",
            "account": "admin",
            "permission": 3,
            "nickname": "admin",
            "phone": "",
            "email": ""
        }
    },
    "msg": "Login successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |        参数说明         |
| :----: | :------: | :---------------------: |
|  code  |   int    |         业务码          |
|  msg   |  string  |        返回消息         |
| token  |  string  | 登录成功返回的用户token |
| error  |  string  |    后端内部错误消息     |
| detail |  string  |        错误详情         |

----



### 删除

**请求路径**：/api/user/delete

**请求方法**：DELETE

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| uid    | String   | 是       | 用户ID   |

**返回结果示例**：

```json
// 删除成功 200
{
    "code": 201,
    "msg": "User deleted successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 更新

**请求路径**：/api/user/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名     | 参数类型 | 是否必填 | 参数说明           |
| ---------- | -------- | -------- | ------------------ |
| uid        | String   | 是       | 用户ID             |
| password   | String   | 否       | 用户更改的密码     |
| nickname   | String   | 否       | 用户更改的昵称     |
| permission | int      | 否       | 用户更改的权限大小 |
| phone      | String   | 否       | 用户更改的电话     |

*注:password,nickname,permission,phone4个可选参数至少需提供一个*

**返回结果示例**：

```json
// 更新成功 200
{
    "code": 201,
    "msg": "User updated successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 列表

**请求路径**：/api/user/list

**请求方法**：GET

**是否需要鉴权：**是

**请求参数**：None

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
|        |          |          |          |

**返回结果示例**：

```json
// 查询成功 200
{
    "code": 201,
    "msg": "Get user list successfully",
    "data": {
        "rows": [
        {
            "account": "admin",
            "nickname": "admin",
            "permission": 3,
            "phone": "1924658487",
            "created_at": "2024-04-07 17:17:44",
            "uid": "00000001"
        },
        {
            "account": "123",
            "nickname": "admin",
            "permission": 3,
            "phone": "",
            "created_at": "2024-04-07 17:17:44",
            "uid": "00000002"
        },
        {
            "account": "1234",
            "nickName": "admin",
            "permission": 3,
            "phone": "",
            "created_at": "2024-04-07 17:17:44",
            "uid": "00000003"
        },
        {
            "account": "1235",
            "nickname": "admin",
            "permission": 3,
            "phone": "",
            "created_at": "2024-04-07 17:17:44",
            "uid": "00000004"
        },
        {
            "account": "1236",
            "nickname": "admin",
            "permission": 3,
            "phone": "123456789",
            "created_at": "2024-04-07 17:17:44",
            "uid": "00000005"
        }
    ]
    }

}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
|  rows  | array[]  |   用户信息数组   |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 获取用户信息

**请求路径**：/api/user/info

**请求方法**：GET

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| :----: | :------: | :------: | :------: |
|        |          |          |          |

**返回结果示例**：

```json
// 查询成功 200
{
    "code": 201,
    "data": {
        "user": {
            "created_at": "2024-06-13T09:35:24+08:00",
            "updated_at": "2024-06-14T09:02:34+08:00",
            "uid": "u00000001",
            "account": "admin",
            "permission": 3,
            "nickname": "admin",
            "phone": "",
            "email": "27653491@qq.com"
        }
    },
    "msg": "success"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
|  user  |  object  |     用户数据     |



### 绑定邮箱（发送验证码）

**请求路径**：/api/user/email/bind

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| :----: | :------: | :------: | :------: |
|  uid   |  String  |    是    |  用户id  |
| email  |  String  |    是    | 绑定邮箱 |

**返回结果示例**：

```json
// 发送成功 200
{
    "msg": "Verification code sent successfully",
	"code":    201,
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 绑定邮箱（校验验证码）

**请求路径**：/api/user/email/verify

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| :----: | :------: | :------: | :------: |
|  uid   |  String  |    是    |  用户id  |
| email  |  String  |    是    | 绑定邮箱 |
|  code  |  String  |    是    |  验证码  |

**返回结果示例**：

```json
// 绑定成功 200
{
    "msg": "Email verification successful",
	"code":    201,
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 重置密码（发送验证码）

**请求路径**：/api/user/psw/reset

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名  | 参数类型 | 是否必填 | 参数说明 |
| :-----: | :------: | :------: | :------: |
| account |  String  |    是    | 用户账号 |
|  email  |  String  |    是    | 绑定邮箱 |

**返回结果示例**：

```json
// 发送成功 200
{
    "msg": "Verification code sent successfully",
	"code":    201,
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 重置密码（校验验证码）

**请求路径**：/api/user/psw/verify

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

|    参数名    | 参数类型 | 是否必填 | 参数说明 |
| :----------: | :------: | :------: | :------: |
|   account    |  String  |    是    | 用户账号 |
|    email     |  String  |    是    | 绑定邮箱 |
|     code     |  String  |    是    |  验证码  |
| new_password |  String  |    是    |  新密码  |

**返回结果示例**：

```json
// 绑定成功 200
{
    "msg": "Password reset successful",
	"code":    201,
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



## 仓库

### 添加

**请求路径**：/api/warehouse/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名  | 参数类型 | 是否必填 | 参数说明                     |
| ------- | -------- | -------- | ---------------------------- |
| name    | String   | 是       | 仓库名                       |
| comment | String   | 否       | 仓库备注                     |
| manager | String   | 是       | 仓库管理员uid                |
| status  | int      | 否       | 仓库启用状态1为启用0为不启用 |

**返回结果示例**：

```json
// 查询成功 200
{
    "code": 201,
    "msg": "Warehouse added successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 删除

**请求路径**：/api/warehouse/delete

**请求方法**：DELETE

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| wid    | String   | 是       | 仓库id   |

**返回结果示例**：

```json
// 删除成功 200
{
    "code": 201,
    "msg": "Warehouse deleted successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 更新

**请求路径**：/api/warehouse/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名  | 参数类型 | 是否必填 | 参数说明                           |
| ------- | -------- | -------- | ---------------------------------- |
| wid     | String   | 是       | 仓库ID                             |
| name    | String   | 否       | 更新的仓库名称                     |
| comment | String   | 否       | 更新的仓库备注                     |
| manager | String   | 否       | 更新的仓库管理员uid                |
| status  | int      | 否       | 更新的仓库启动状态1为启用0为不启用 |

*注:name,comment,manager,status 4个可选参数至少需提供一个*

**返回结果示例**：

```json
// 更新成功 200
{
    "code": 201,
    "msg": "Warehouse updated successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 列表

**请求路径**：/api/warehouse/list

**请求方法**：GET

**是否需要鉴权：**是

**请求参数**：None

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
|        |          |          |          |

**返回结果示例**：

```json
// 查询成功 200
{
    "code": 201,
    "msg": "Get warehouse list successfully",
    "data": {
    	"rows": [
        {
            "created_at": "2024-04-07 17:17:44",
            "comment": "123",
            "manager": "00000001",
            "name": "F",
            "status": 1,
            "wid": "000001"
        },
        {
            "created_at": "2024-04-07 17:17:44",
            "comment": "123",
            "manager": "00000002",
            "name": "A",
            "status": 0,
            "wid": "000002"
        }
    ]
    },
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
|  rows  | array[]  |   仓库信息数组   |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



## 货品类型

### 添加

**请求路径**：/api/gt/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名    | 参数类型 | 是否必填 | 参数说明     |
| --------- | -------- | -------- | ------------ |
| name      | String   | 是       | 货品类型名   |
| type_code | String   | 否       | 货品类型编码 |

**返回结果示例**：

```json
// 添加成功 200
{
    "code": 201,
    "msg": "Goods type added successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 更新

**请求路径**：/api/gt/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名    | 参数类型 | 是否必填 | 参数说明           |
| --------- | -------- | -------- | ------------------ |
| gtid      | String   | 是       | 货品类型ID         |
| name      | String   | 否       | 更新的货品类型名   |
| type_code | String   | 否       | 更新的货品类型编码 |

*注:name，type_code2个可选参数至少需提供一个*

**返回结果示例**：

```json
// 更新成功 200
{
    "code": 201,
    "msg": "Goods type updated successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 删除

**请求路径**：/api/gt/delete

**请求方法**：DELETE

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明   |
| ------ | -------- | -------- | ---------- |
| gtid   | String   | 是       | 货品类型id |

**返回结果示例**：

```json
// 删除成功 200
{
    "code": 201,
    "msg": "Goods type deleted successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 列表

**请求路径**：/api/gt/list

**请求方法**：GET

**是否需要鉴权：**是

**请求参数**：None

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
|        |          |          |          |

**返回结果示例**：

```json
// 查询成功 200
{
    "code": 201,
    "msg": "Get user list successfully",
    "data": {
        "rows": [
        {
            "created_at": "2024-04-07 17:17:44",
            "gtid": "0001",
            "name": "yzl",
            "type_code": "bzx"
        },
        {
            "created_at": "2024-04-07 17:17:44",
            "gtid": "0002",
            "name": "印字轮",
            "type_code": "yzl"
        }
    ]
    },
    
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
|  rows  | array[]  |   仓库信息数组   |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



## 部门

### 添加

**请求路径**：/api/dept/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| name   | String   | 是       | 部门名称 |

**返回结果示例**：

```json
// 添加成功 200
{
    "code": 201,
    "msg": "Department added successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 更新

**请求路径**：/api/dept/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| did    | String   | 是       | 部门id   |
| name   | String   | 是       | 部门名称 |

**返回结果示例**：

```json
// 更新成功 200
{
    "code": 201,
    "msg": "Department updated successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 删除

**请求路径**：/api/dept/delete

**请求方法**：DELETE

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| did    | String   | 是       | 部门id   |

**返回结果示例**：

```json
// 删除成功 200
{
    "code": 201,
    "msg": "Department deleted successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 列表

**请求路径**：/api/dept/list

**请求方法**：GET

**是否需要鉴权：**是

**请求参数**：None

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
|        |          |          |          |

**返回结果示例**：

```json
// 查询成功 200
{
    "code": 201,
    "msg": "Get departments list successfully",
    “data": {
        "rows": [
            {
                "created_at": "2024-04-07 17:17:44",
                "did": "d0002",
                "name": "生产"
            }
        ]
	}
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
|  rows  | array[]  |   仓库信息数组   |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



## 员工

### 添加

**请求路径**：/api/staff/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名     | 参数类型 | 是否必填 | 参数说明   |
| ---------- | -------- | -------- | ---------- |
| name       | String   | 是       | 员工名称   |
| phone      | string   | 否       | 电话       |
| department | string   | 是       | 所属部门id |

**返回结果示例**：

```json
// 添加成功 200
{
    "code": 201,
    "msg": "Staff added successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 批量添加（导入）

**请求路径**：/api/staff/upload

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型       | 是否必填 | 参数说明 |
| ------ | -------------- | -------- | -------- |
| list   | Array <Object> | 是       | 员工列表 |

**列表参数**

| 参数名     | 参数类型 | 是否必填 | 参数说明   |
| ---------- | -------- | -------- | ---------- |
| name       | String   | 是       | 员工名称   |
| phone      | string   | 否       | 电话       |
| department | string   | 是       | 所属部门id |

**返回结果示例**：

```json
// 部分导入失败 217
{
    "code": 202,
    "data": {
        "detail": [
        {
            "code": 403,
            "msg": "The staff's zzt2 department does not exist"
        }
    ]
    },
    "msg": "Some staffs failed to register; 1/2"
}
// 全部导入失败 400
{
    "code": 203,
    "data": {
        "detail": [
        {
            "code": 403,
            "msg": "The staff's zzt3 department does not exist"
        },
        {
            "code": 403,
            "msg": "The staff's zzt2 department does not exist"
        }
    ]
    },
    "msg": "All staffs failed to register"
}
```

**返回数据说明**

| 参数名 | 参数类型 |       参数说明       |
| :----: | :------: | :------------------: |
|  code  |   int    |        业务码        |
|  msg   |  string  |       返回消息       |
|  uid   |  string  | 注册成功返回的用户id |
| error  |  string  |   后端内部错误消息   |
| detail |  string  |   每个错误注册详情   |

----



### 更新

**请求路径**：/api/staff/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名  | 参数类型 | 是否必填 | 参数说明       |
| ------- | -------- | -------- | -------------- |
| sid     | String   | 是       | 员工id         |
| name    | String   | 否       | 员工名称       |
| phone   | string   | 否       | 电话           |
| dept_id | string   | 否       | 员工所属部门id |

**返回结果示例**：

```json
// 更新成功 200
{
    "code": 201,
    "msg": "Staff updated successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 删除

**请求路径**：/api/staff/delete

**请求方法**：DELETE

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| sid    | String   | 是       | 员工id   |

**返回结果示例**：

```json
// 删除成功 200
{
    "code": 201,
    "msg": "Staff deleted successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 列表

**请求路径**：/api/staff/list

**请求方法**：GET

**是否需要鉴权：**是

**请求参数**：None

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
|        |          |          |          |

**返回结果示例**：

```json
// 查询成功 200
{
    "code": 201,
    "msg": "Get staffs list successfully",
    "data": {
        "rows": [
        {
            "created_at": "2024-04-07 17:17:44",
            "department": "d0002",
            "name": "李四",
            "phone": "",
            "sid": "s00000002"
        },
        {
            "created_at": "2024-04-07 17:17:44",
            "department": "d0002",
            "name": "王五",
            "phone": "",
            "sid": "s00000003"
        },
        {
            "created_at": "2024-04-07 17:17:44",
            "department": "d0002",
            "name": "王liu",
            "phone": "",
            "sid": "s00000004"
        }
    ]
    }
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
|  rows  | array[]  |   仓库信息数组   |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



## 出入库类型

### 添加

**请求路径**：/api/invt/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名    | 参数类型 | 是否必填 | 参数说明        |
| --------- | -------- | -------- | --------------- |
| name      | String   | 是       | 出入库类型名称  |
| type_code | string   | 否       | 类型编码        |
| type      | int      | 是       | 1为入库 2为出库 |

**返回结果示例**：

```json
// 添加成功 200
{
    "code": 201,
    "msg": "Inventory type added successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 更新

**请求路径**：/api/invt/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名    | 参数类型 | 是否必填 | 参数说明   |
| --------- | -------- | -------- | ---------- |
| itid      | String   | 是       | 类型id     |
| name      | String   | 否       | 类型名称   |
| type_code | string   | 否       | 类型编码   |
| type      | int      | 否       | 出入库标识 |

*注:name，type_code2个可选参数至少需提供一个*

**返回结果示例**：

```json
// 更新成功 200
{
    "code": 201,
    "msg": "Inventory type updated successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 删除

**请求路径**：/api/invt/delete

**请求方法**：DELETE

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| itid   | String   | 是       | 类型id   |

**返回结果示例**：

```json
// 删除成功 200
{
    "code": 201,
    "msg": "Inventory type deleted successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 列表

**请求路径**：/api/invt/list

**请求方法**：GET

**是否需要鉴权：**是

**请求参数**：None

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
|        |          |          |          |

**返回结果示例**：

```json
// 查询成功 200
{
    "code": 201,
    "msg": "Get inventory type list successfully",
    "data": {
        "rows": [
        {
            "created_at": "2024-06-03T10:09:47+08:00",
            "itid": "_default1_",
            "name": "默认入库",
            "type": 1,
            "type_code": "default_in"
        },
        {
            "created_at": "2024-06-03T10:09:47+08:00",
            "itid": "_default2_",
            "name": "默认出库",
            "type": 2,
            "type_code": "default_out"
        }
    ]
    }
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
|  rows  | array[]  |   仓库信息数组   |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

-----



## 计量单位

### 添加

**请求路径**：/api/unit/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| name   | String   | 是       | 单位     |

**返回结果示例**：

```json
// 添加成功 200
{
    "code": 201,
    "msg": "Unit added successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 更新

**请求路径**：/api/unit/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| unid   | String   | 是       | 单位id   |
| name   | String   | 是       | 单位名称 |

**返回结果示例**：

```json
// 更新成功 200
{
    "code": 201,
    "msg": "unit type updated successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 删除

**请求路径**：/api/unit/delete

**请求方法**：DELETE

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| unid   | String   | 是       | 单位id   |

**返回结果示例**：

```json
{
    "code": 201,
    "msg": "Unit deleted successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 列表

**请求路径**：/api/unit/list

**请求方法**：GET

**是否需要鉴权：**是

**请求参数**：None

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
|        |          |          |          |

**返回结果示例**：

```json
// 查询成功 200
{
    "code": 201,
    "msg": "Get inventory type list successfully",
    "data": {
         "rows": [
        {
            "name": "km",
            "unid": "un0002"
        },
        {
            "name": "kg",
            "unid": "un0003"
        },
        {
            "name": "轴",
            "unid": "un0004"
        }
    ]
    }  
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
|  rows  | array[]  |   仓库信息数组   |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



## 货品

### 添加

**请求路径**：/api/goods/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名       | 参数类型 | 是否必填 | 参数说明           |
| ------------ | -------- | -------- | ------------------ |
| name         | String   | 是       | 货品名称           |
| model        | String   | 否       | 类型规格           |
| goods_code   | String   | 是       | 货品编码           |
| goods_type   | String   | 是       | 货品类型id（gtid） |
| manufacturer | String   | 否       | 制造商             |
| unit         | String   | 是       | 单位id（unid）     |
| unit_price   | float    | 否       | 单价               |

**返回结果示例**：

```json
// 添加成功 200
{
    "code": 201,
    "msg": "Goods added successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 更新

**请求路径**：/api/goods/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名       | 参数类型 | 是否必填 | 参数说明           |
| ------------ | -------- | -------- | ------------------ |
| gid          | string   | 是       | 货品id             |
| name         | String   | 否       | 货品名称           |
| model        | String   | 否       | 规格类型           |
| goods_code   | String   | 否       | 货品编码           |
| goods_type   | String   | 否       | 货品类型id（gtid） |
| manufacturer | String   | 否       | 制造商             |
| unit         | String   | 否       | 单位id（unid）     |
| unit_price   | float    | 否       | 单价               |

*注:8个可选参数至少需提供一个*

**返回结果示例**：

```json
// 更新成功 200
{
    "code": 201,
    "msg": "Goods updated successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 删除

**请求路径**：/api/goods/delete

**请求方法**：DELETE

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| gid    | String   | 是       | 货品id   |

**返回结果示例**：

```json
// 删除成功 200
{
    "code": 201,
    "msg": "Goods deleted successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 查询

**请求路径**：/api/goods/search

**请求方法**：GET

**是否需要鉴权：**是

**请求参数**：None

| 参数名       | 参数类型 | 是否必填 | 参数说明                 |
| ------------ | -------- | -------- | ------------------------ |
| page         | int      | 否       | 页数，默认为1            |
| page_size    | int      | 否       | 单页大小，默认为10       |
| gid          | string   | 否       | 依据gid查询              |
| name         | string   | 否       | 依货品名称查询           |
| model        | string   | 否       | 依货品规格类型查询       |
| goods_type   | string   | 否       | 依货品类型id（gtid）查询 |
| manufacturer | string   | 否       | 依货品所属生产厂商查询   |
| unit_price   | float    | 否       | 依单价查询               |
| keyword      | string   | 否       | 关键字模糊查询           |

**返回结果示例**：

```json
// 查询成功（有数据） 200
{
    "code": 201,
    "msg": "Query successfully",
    "data": {
    "keyword": "1",
    "page": 1,
    "page_size": 10,
    "rows": [
        {
            "created_at": "2024-06-11T16:21:00+08:00",
            "updated_at": "2024-06-11T16:35:28+08:00",
            "gid": "ga42b87cf",
            "goods_code": "HP001",
            "name": "货品1",
            "model": "YYPACW",
            "goods_type": "_default_",
            "manufacturer": "生产商1",
            "unit": "_default_",
            "images": [
                {
                    "path": "static/res/goodsImage/goods_17180948221.png"
                }
            ],
            "files": [
                {
                    "path": "static/res/goodsFile/goods_17180949281.jpg"
                },
                {
                    "path": "static/res/goodsFile/goods_17180949282.jpg"
                }
            ],
            "unit_price": 139
        },
        {
            "created_at": "2024-06-11T16:21:00+08:00",
            "updated_at": "2024-06-11T16:48:04+08:00",
            "gid": "g242bcd49",
            "goods_code": "HP002",
            "name": "货品2",
            "model": "GPARDR",
            "goods_type": "gt09a5d1ac",
            "manufacturer": "生产商2",
            "unit": "un0b7be337",
            "images": null,
            "files": [
                {
                    "path": "static/res/goodsFile/goods_g242bcd49_1.xlsx"
                }
            ],
            "unit_price": 260
        },
        {
            "created_at": "2024-06-11T16:21:00+08:00",
            "updated_at": "2024-06-11T16:21:00+08:00",
            "gid": "g05c1a76d",
            "goods_code": "HP004",
            "name": "货品4",
            "model": "ZESZYA",
            "goods_type": "gt09a5d1ac",
            "manufacturer": "生产商4",
            "unit": "un0b7be337",
            "images": null,
            "files": null,
            "unit_price": 548
        },
        {
            "created_at": "2024-06-11T16:21:00+08:00",
            "updated_at": "2024-06-11T16:21:00+08:00",
            "gid": "g0ee64da5",
            "goods_code": "HP0010",
            "name": "货品10",
            "model": "GFDTJL",
            "goods_type": "gt8f736a99",
            "manufacturer": "生产商10",
            "unit": "unb91a76f8",
            "images": null,
            "files": null,
            "unit_price": 348
        }
    ],
    "total": 4,
    "total_pages": 1
    }
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
|  rows  | array[]  |   仓库信息数组   |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 图片上传

**请求路径**：/api/upload/goods_img

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型  | 是否必填 | 参数说明 |
| ------ | --------- | -------- | -------- |
| goods  | String    | 是       | 货品id   |
| image  | Multipart | 是       | 图片文件 |

*支持图片格式：*`.png, .jpg, .jpeg, .gif, .svg`

**返回结果示例**：

```json
// 上传成功 200
{
  "code": 201,
  "data": [
    {
      "gid": "ga42b87cf",
      "name": "goods_17180949281.jpg",
      "path": "static/res/goodsFile/goods_17180949281.jpg"
    },
    {
      "gid": "ga42b87cf",
      "name": "goods_17180949282.jpg",
      "path": "static/res/goodsFile/goods_17180949282.jpg"
    }
  ],
  "msg": "Upload successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |
|  data  |  string  |    返回数据体    |
|  gid   |  string  |    目标货品id    |
|  name  |  string  |      图片名      |
|  path  |  string  |     访问路径     |

----

### 附件上传

**请求路径**：/api/upload/goods_file

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型  | 是否必填 | 参数说明 |
| ------ | --------- | -------- | -------- |
| goods  | String    | 是       | 货品id   |
| file   | Multipart | 是       | 图片文件 |

*支持文件格式：*`.doc .docx .pdf .xls .xlsx .ppt .pptx .txt .zip .rar .png .jpg .jpeg .gif .svg`

**返回结果示例**：

```json
// 上传成功 200
{
  "code": 201,
  "data": [
    {
      "gid": "ga42b87cf",
      "name": "goods_17180949281.jpg",
      "path": "static/res/goodsFile/goods_17180949281.jpg"
    },
    {
      "gid": "ga42b87cf",
      "name": "goods_17180949282.jpg",
      "path": "static/res/goodsFile/goods_17180949282.jpg"
    }
  ],
  "msg": "Upload successfully"
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |
|  data  |  string  |    返回数据体    |
|  gid   |  string  |    目标货品id    |
|  name  |  string  |      附件名      |
|  path  |  string  |     访问路径     |

----



## 出入库

### 添加

**请求路径**：/api/inv/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名         | 参数类型 | 是否必填 | 参数说明                          |
| -------------- | -------- | -------- | --------------------------------- |
| date           | String   | 否       | 单据日期 格式:2006-01-02 15:04:05 |
| number         | String   | 否       | 单号                              |
| department     | String   | 否       | 所属部门id(did)                   |
| goods_list     | String   | 是       | 货品数组 格式见下                 |
| warehouse      | String   | 是       | 所属仓库(wid)                     |
| inventory_type | String   | 是       | 出入库类型(itid)                  |
| manufacturer   | String   | 否       | 制造商                            |
| operator       | String   | 是       | 操作员(sid)                       |
| comment        | String   | 否       | 备注                              |

** goods_list数据结构*：

```json
[
    {"goods": "g4c182157", "amount":26,"comment":"test"},
    {"goods": "g08943f59", "amount":23,"comment":"test"}
]
```

**参数说明**

| 字段 | goods       | amount | comment |
| ---- | ----------- | ------ | ------- |
| 类型 | string      | float  | string  |
| 说明 | 货品id(gid) | 数量   | 备注    |



**返回结果示例**：

```json
// 添加成功 200
{
    "code": 201,
    "msg": "Inventory added successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----

### 更新

**请求路径**：/api/inv/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名         | 参数类型 | 是否必填 | 参数说明                          |
| -------------- | -------- | -------- | --------------------------------- |
| iid            | String   | 是       | 出入库单id                        |
| date           | String   | 否       | 单据日期 格式:2006-01-02 15:04:05 |
| number         | String   | 否       | 单号                              |
| department     | String   | 否       | 所属部门id(did)                   |
| goods_list     | String   | 是       | 货品数组 格式见下                 |
| warehouse      | String   | 是       | 所属仓库(wid)                     |
| inventory_type | String   | 是       | 出入库类型(itid)                  |
| manufacturer   | String   | 否       | 制造商                            |
| operator       | String   | 是       | 操作员(sid)                       |
| comment        | String   | 否       | 备注                              |

** goods_list数据结构*：

```json
[
    {"goods": "g4c182157", "amount":26,"comment":"test"},
    {"goods": "g08943f59", "amount":23,"comment":"test"}
]
```

**参数说明**

| 字段 | goods       | amount | comment |
| ---- | ----------- | ------ | ------- |
| 类型 | string      | float  | string  |
| 说明 | 货品id(gid) | 数量   | 备注    |



**返回结果示例**：

```json
// 更新成功 200
{
    "code": 201,
    "msg": "Inventory added successfully",
    "data": null
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



### 查询

**请求路径**：/api/inv/search

**请求方法**：GET

**是否需要鉴权：**是

| 参数名         | 参数类型 | 是否必填 | 参数说明                        |
| -------------- | -------- | -------- | ------------------------------- |
| page           | int      | 否       | 页数，默认为1 ；为-1时不分页    |
| page_size      | int      | 否       | 单页大小，默认为10              |
| goods          | string   | 否       | 依货品id(gid)查询               |
| number         | string   | 否       | 依订单编号查询                  |
| amount         | int      | 否       | 依订单数量查询                  |
| warehouse      | string   | 否       | 依货品所属仓库id（wid）查询     |
| manufacturer   | string   | 否       | 依货品所属生产厂商查询          |
| inventory_type | string   | 否       | 依出入库类型查询                |
| type           | int      | 否       | 依出入库类型标识查询            |
| operator       | string   | 否       | 依操作员查询(sid)               |
| comment        | string   | 否       | 依备注查询                      |
| date           | string   | 否       | 依创建日期查询 格式：2024-06-03 |
| keyword        | string   | 否       | 关键字模糊查询                  |

**返回结果示例**：

```json
// 查询成功（有数据） 200
{
    "code": 201,
     "msg": "Query successfully",
    "data": {
    "keyword": "",
    "page": 1,
    "page_size": 2,
    "rows": [
        {
            "amount": 120,
            "comment": "test",
            "created_at": "2024-06-03T11:26:07+08:00",
            "goods": {
                "created_at": "2024-06-03T11:26:07+08:00",
                "gid": "gb003821b",
                "goods_code": "",
                "goods_type": "_default_",
                "image": null,
                "file":null
                "manufacturer": "",
                "model": "",
                "name": "嘛",
                "quantity": 120,
                "unit": "_default_",
                "unit_price": 0,
                "updated_at": "2024-06-03T11:26:07+08:00"
            },
            "iid": "i211bb1ee",
            "inventory_type": "_default1_",
            "manufacturer": "",
            "number": "I202406031126G323a",
            "operator": "_default_",
            "update_at": "2024-06-03T11:26:07+08:00",
            "warehouse": "_default_"
        }
    ],
    "total": 1,
    "total_pages": 1
    }
}
```

**返回数据说明**

| 参数名 | 参数类型 |     参数说明     |
| :----: | :------: | :--------------: |
|  code  |   int    |      业务码      |
|  msg   |  string  |     返回消息     |
|  rows  | array[]  |   仓库信息数组   |
| error  |  string  | 后端内部错误消息 |
| detail |  string  |     错误详情     |

----



## 库存

### 查询

**请求路径**：/api/stock/get

**请求方法**：GET

**是否需要鉴权：**是

| 参数名    | 参数类型 | 是否必填 | 参数说明    |
| --------- | -------- | -------- | ----------- |
| warehouse | string   | 否       | 仓库id(wid) |
| goods     | string   | 否       | 货品id(gid) |

*注：只提供warehouse返回该仓库所有货品信息，只提供goods返回该货品所有库存信息*

**返回结果示例**：

```json
// 查询成功（1对1库存） 201
{
    "code": 200,
    "data": {
        "goods": "g097e2ee2",
        "quantity": 150,
        "warehouse": "_default_"
    },
    "msg": "Get stock success"
}
// 查询成功（1对多） 202
{
    "code": 202,
    "msg": "Get stock success",
    "rows": [
        {
            "created_at": "2024-06-07T16:59:34+08:00",
            "updated_at": "2024-06-07T17:00:58+08:00",
            "goods": "g097e2ee2",
            "warehouse": "_default_",
            "quantity": 150
        }
    ]
}
// 查询成功（总库存） 203
{
    "code": 203,
    "data": {
        "rows": [
            {
                "created_at": "2024-06-12T10:36:41+08:00",
                "updated_at": "2024-06-12T11:00:49+08:00",
                "goods": "gf02c82db",
                "warehouse": "_default_",
                "quantity": 200
            },
            {
                "created_at": "2024-06-12T11:06:25+08:00",
                "updated_at": "2024-06-12T11:06:24+08:00",
                "goods": "gf02c82db",
                "warehouse": "wcd74b1e5",
                "quantity": 100
            },
            {
                "created_at": "2024-06-12T11:06:26+08:00",
                "updated_at": "2024-06-12T11:06:27+08:00",
                "goods": "gf02c82db",
                "warehouse": "wf716f6dc",
                "quantity": 500
            }
        ],
        "total": 800
    },
    "msg": "Get stock success"
}
```

**返回数据说明**

| 参数名 | 参数类型 |   参数说明   |
| :----: | :------: | :----------: |
|  code  |   int    |    业务码    |
|  msg   |  string  |   返回消息   |
|  data  |  string  |   返回数据   |
|  rows  |  string  | 返回数据列表 |
