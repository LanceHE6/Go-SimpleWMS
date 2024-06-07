# 后端接口文档



*测试环境后台地址：*

```
http://47.236.80.244:6007
https://api.simplewms.hycerlance.fun/api
```



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
    "message": "Hello Go-SimpleWMS"
}
```

----




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
{
    "code": 401,
	"message": "No Authorization header provided"
}
```

### 鉴权相关返回数据说明

| 参数名  | 参数类型 |       参数说明       |
| :-----: | :------: | :------------------: |
|  code   |   int    |        业务码        |
|  error  |  string  | 后端服务内部错误消息 |
| message |  string  |     请求返回消息     |

### 鉴权相关返回状态码说明

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  401   | StatusUnauthorized  |    鉴权未通过    |
|  403   |   StatusForbidden   |     权限拒绝     |
|  500   | InternalServerError | 后端服务内部错误 |

----

## 业务码 code 说明

| code | 1xx  |   2xx    |   4xx    |   5xx    |
| :--: | :--: | :------: | :------: | :------: |
| 说明 | 鉴权 | 处理成功 | 处理失败 | 内部错误 |



----



## 用户

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
{
    "code": 201,
    "message": "User registered successfully",
    "uid": "00000002"
}
```

**返回数据说明**

| 参数名  | 参数类型 |       参数说明       |
| :-----: | :------: | :------------------: |
|  code   |   int    |        业务码        |
| message |  string  |       返回消息       |
|   uid   |  string  | 注册成功返回的用户id |
|  error  |  string  |   后端内部错误消息   |
| detail  |  string  |       错误详情       |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     注册成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  403   |      Forbidden      |   账号已被注册   |
|  500   | InternalServerError | 后端服务内部错误 |

----

### 批量注册（上传用户）

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
{
    "code": 403,
    "detail": [
        {
            "code": 402,
            "message": "The account '87' has been registered"
        }
    ],
    "message": "Some users failed to register; 1/2"
}

{
    "code": 402,
    "detail": [
        {
            "code": 402,
            "message": "The account '87' has been registered"
        },
        {
            "code": 402,
            "message": "The account '81' has been registered"
        }
    ],
    "message": "All users failed to register"
}
```

**返回数据说明**

| 参数名  | 参数类型 |       参数说明       |
| :-----: | :------: | :------------------: |
|  code   |   int    |        业务码        |
| message |  string  |       返回消息       |
|   uid   |  string  | 注册成功返回的用户id |
|  error  |  string  |   后端内部错误消息   |
| detail  |  string  |   每个错误注册详情   |

**返回状态码说明**

| 状态码 |        含义         |             说明             |
| :----: | :-----------------: | :--------------------------: |
|  200   |         OK          |           注册成功           |
|  217   |   PartialContent    |        注册不完全成功        |
|  400   |     BadRequest      | 请求参数不全获取注册全部失败 |
|  500   | InternalServerError |       后端服务内部错误       |

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
{
    "code": 201,
    "message": "Login successfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk1NTUxMjEsInN1YiI6IjAwMDAwMDAyIn0.hvz-Xp9kfhVNsCy6Q9nhS9wM8-c-DgJJ8PLcME17Fto"
}
```

**返回数据说明**

| 参数名  | 参数类型 |        参数说明         |
| :-----: | :------: | :---------------------: |
|  code   |   int    |         业务码          |
| message |  string  |        返回消息         |
|  token  |  string  | 登录成功返回的用户token |
|  error  |  string  |    后端内部错误消息     |
| detail  |  string  |        错误详情         |

**返回状态码说明**

| 状态码 |         含义         |       说明       |
| :----: | :------------------: | :--------------: |
|  200   |          OK          |     登录成功     |
|  203   | NonAuthoritativeInfo | 账号或密码不正确 |
|  400   |      BadRequest      |   请求参数不全   |
|  500   | InternalServerError  | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "User deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     登录成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "User updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Get user list successfully",
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
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  rows   | array[]  |   用户信息数组   |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Warehouse added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  403   |      Forbidden      |    仓库已存在    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Warehouse deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Warehouse updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |            说明             |
| :----: | :-----------------: | :-------------------------: |
|  200   |         OK          |          修改成功           |
|  400   |     BadRequest      |        请求参数不全         |
|  401   |    Unauthorized     |         鉴权未通过          |
|  403   |      Forbidden      | 仓库名已经存在/该仓库不存在 |
|  500   | InternalServerError |      后端服务内部错误       |

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
{
    "code": 201,
    "message": "Get warehouse list successfully",
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
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  rows   | array[]  |   仓库信息数组   |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Goods type added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  403   |      Forbidden      |    仓库已存在    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Goods type updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |            说明             |
| :----: | :-----------------: | :-------------------------: |
|  200   |         OK          |          修改成功           |
|  400   |     BadRequest      |        请求参数不全         |
|  401   |    Unauthorized     |         鉴权未通过          |
|  403   |      Forbidden      | 仓库名已经存在/该仓库不存在 |
|  500   | InternalServerError |      后端服务内部错误       |

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
{
    "code": 201,
    "message": "Goods type deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Get user list successfully",
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
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  rows   | array[]  |   仓库信息数组   |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Department added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  403   |      Forbidden      |    仓库已存在    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Department updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |            说明             |
| :----: | :-----------------: | :-------------------------: |
|  200   |         OK          |          修改成功           |
|  400   |     BadRequest      |        请求参数不全         |
|  401   |    Unauthorized     |         鉴权未通过          |
|  403   |      Forbidden      | 仓库名已经存在/该仓库不存在 |
|  500   | InternalServerError |      后端服务内部错误       |

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
{
    "code": 201,
    "message": "Department deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Get departments list successfully",
    "rows": [
        {
            "created_at": "2024-04-07 17:17:44",
            "did": "d0002",
            "name": "生产"
        }
    ]
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  rows   | array[]  |   仓库信息数组   |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

----



## 员工

### 添加

**请求路径**：/api/staff/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名  | 参数类型 | 是否必填 | 参数说明   |
| ------- | -------- | -------- | ---------- |
| name    | String   | 是       | 员工名称   |
| phone   | string   | 否       | 电话       |
| dept_id | string   | 是       | 所属部门id |

**返回结果示例**：

```json
{
    "code": 201,
    "message": "Staff added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  403   |      Forbidden      |    仓库已存在    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Staff updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |            说明             |
| :----: | :-----------------: | :-------------------------: |
|  200   |         OK          |          修改成功           |
|  400   |     BadRequest      |        请求参数不全         |
|  401   |    Unauthorized     |         鉴权未通过          |
|  403   |      Forbidden      | 仓库名已经存在/该仓库不存在 |
|  500   | InternalServerError |      后端服务内部错误       |

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
{
    "code": 201,
    "message": "Staff deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Get staffs list successfully",
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
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  rows   | array[]  |   仓库信息数组   |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Inventory type added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  403   |      Forbidden      |    类型已存在    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Inventory type updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |            说明             |
| :----: | :-----------------: | :-------------------------: |
|  200   |         OK          |          修改成功           |
|  400   |     BadRequest      |        请求参数不全         |
|  401   |    Unauthorized     |         鉴权未通过          |
|  403   |      Forbidden      | 类型名已经存在/该类型不存在 |
|  500   | InternalServerError |      后端服务内部错误       |

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
{
    "code": 201,
    "message": "Inventory type deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Get inventory type list successfully",
    "rows": [
        {
            "created_at": "2024-04-07 17:17:44",
            "gtid": "it0001",
            "name": "采购入库",
            "type_code": "cg"
        },
        {
            "created_at": "2024-04-07 17:17:44",
            "gtid": "it0002",
            "name": "生产入库",
            "type_code": "sc"
        }
    ]
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  rows   | array[]  |   仓库信息数组   |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Unit added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  403   |      Forbidden      |    类型已存在    |
|  500   | InternalServerError | 后端服务内部错误 |



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
{
    "code": 201,
    "message": "unit type updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |            说明             |
| :----: | :-----------------: | :-------------------------: |
|  200   |         OK          |          修改成功           |
|  400   |     BadRequest      |        请求参数不全         |
|  401   |    Unauthorized     |         鉴权未通过          |
|  403   |      Forbidden      | 类型名已经存在/该类型不存在 |
|  500   | InternalServerError |      后端服务内部错误       |

----

### 

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
    "message": "Unit deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Get inventory type list successfully",
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
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  rows   | array[]  |   仓库信息数组   |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
| goods_code   | String   | 否       | 货品编码           |
| goods_type   | String   | 是       | 货品类型id（gtid） |
| manufacturer | String   | 否       | 制造商             |
| unit         | String   | 是       | 单位id（unid）     |
| unit_price   | float    | 否       | 单价               |

**返回结果示例**：

```json
{
    "code": 201,
    "message": "Goods added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  403   |      Forbidden      |    类型已存在    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Goods updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |            说明             |
| :----: | :-----------------: | :-------------------------: |
|  200   |         OK          |          修改成功           |
|  400   |     BadRequest      |        请求参数不全         |
|  401   |    Unauthorized     |         鉴权未通过          |
|  403   |      Forbidden      | 类型名已经存在/该类型不存在 |
|  500   | InternalServerError |      后端服务内部错误       |

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
{
    "code": 201,
    "message": "Goods deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "keyword": "重",
    "message": "Query successfully",
    "page": 1,
    "page_size": 10,
    "rows": [
        {
            "ID": 1,
            "CreatedAt": "2024-04-15T20:24:46+08:00",
            "UpdatedAt": "2024-04-15T20:24:46+08:00",
            "Gid": "g07a99cb7",
            "GoodsCode": "XJ001",
            "Name": "橡胶",
            "Model": "XJ-JT",
            "GoodsType": "gtd180d212",
            "Warehouse": "w62d263ec",
            "Manufacturer": "重游",
            "Unit": "un546af271",
            "Image": "",
            "Quantity": 200
        },
        {
            "ID": 2,
            "CreatedAt": "2024-04-15T20:27:39+08:00",
            "UpdatedAt": "2024-04-15T20:27:39+08:00",
            "Gid": "gc034b22e",
            "GoodsCode": "XJ002",
            "Name": "橡胶",
            "Model": "XJ-JT",
            "GoodsType": "gtd180d212",
            "Warehouse": "w62d263ec",
            "Manufacturer": "重游",
            "Unit": "un546af271",
            "Image": "",
            "Quantity": 300
        },
        {
            "ID": 3,
            "CreatedAt": "2024-04-15T20:27:47+08:00",
            "UpdatedAt": "2024-04-15T20:27:47+08:00",
            "Gid": "g643cb699",
            "GoodsCode": "XJ003",
            "Name": "橡胶",
            "Model": "XJ-JT",
            "GoodsType": "gtd180d212",
            "Warehouse": "w62d263ec",
            "Manufacturer": "重游",
            "Unit": "un546af271",
            "Image": "",
            "Quantity": 400
        }
    ],
    "total": 3,
    "total_pages": 1
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  rows   | array[]  |   仓库信息数组   |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

----



### 图片上传

**请求路径**：/api/upload/goods_img

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名 | 参数类型 | 是否必填 | 参数说明 |
| ------ | -------- | -------- | -------- |
| gid    | String   | 是       | 货品id   |
| image  | File     | 是       | 图片文件 |

**返回结果示例**：

```json
{
    "code": 201,
    "data": {
        "gid": "g07a99cb7",
        "image_name": "goods_1716190218.png",
        "image_path": "static/res/goodsImage/goods_1716190218.png"
    },
    "message": "Upload successfully"
}
```

**返回数据说明**

|   参数名   | 参数类型 |     参数说明     |
| :--------: | :------: | :--------------: |
|    code    |   int    |      业务码      |
|  message   |  string  |     返回消息     |
|   error    |  string  | 后端内部错误消息 |
|   detail   |  string  |     错误详情     |
|    data    |  string  |    返回数据体    |
|    gid     |  string  |    目标货品id    |
| image_name |  string  |      图片名      |
| image_path |  string  |     访问路径     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Inventory added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  403   |      Forbidden      |    类型已存在    |
|  500   | InternalServerError | 后端服务内部错误 |

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
{
    "code": 201,
    "message": "Inventory added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  400   |     BadRequest      |   请求参数不全   |
|  401   |    Unauthorized     |    鉴权未通过    |
|  403   |      Forbidden      |    类型已存在    |
|  500   | InternalServerError | 后端服务内部错误 |

----



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
{
    "code": 201,
    "keyword": "",
    "message": "Query successfully",
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
                "image": "",
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
```

**返回数据说明**

| 参数名  | 参数类型 |     参数说明     |
| :-----: | :------: | :--------------: |
|  code   |   int    |      业务码      |
| message |  string  |     返回消息     |
|  rows   | array[]  |   仓库信息数组   |
|  error  |  string  | 后端内部错误消息 |
| detail  |  string  |     错误详情     |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |

----

