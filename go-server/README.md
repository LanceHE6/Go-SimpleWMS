# 后端接口文档

## 请求说明

**所有接口需携带Content-Type头**

示例：

```
headers:{
	"Content-Type": "application/json"
}
```



## 鉴权说明

### 鉴权方式：

请求头携带鉴权JWT token

**请求路径**：/auth

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

**请求路径**：/user/register

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

**请求路径**：/user/upload

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名    | 参数类型       | 是否必填 | 参数说明 |
| --------- | -------------- | -------- | -------- |
| user_list | Array <Object> | 是       | 用户列表 |

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
    "code": 201,
    "message": "User registered successfully"
}

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

**请求路径**：/user/login

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

**请求路径**：/user/delete

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

**请求路径**：/user/update

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

**请求路径**：/user/list

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

**请求路径**：/warehouse/add

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

**请求路径**：/warehouse/delete

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

**请求路径**：/warehouse/update

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

**请求路径**：/warehouse/list

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

**请求路径**：/gt/add

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

**请求路径**：/gt/update

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

**请求路径**：/gt/delete

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

**请求路径**：/gt/list

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

**请求路径**：/dept/add

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

**请求路径**：/dept/update

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

**请求路径**：/dept/delete

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

**请求路径**：/dept/list

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

**请求路径**：/staff/add

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

**请求路径**：/staff/update

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

**请求路径**：/staff/delete

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

**请求路径**：/staff/list

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

**请求路径**：/invt/add

**请求方法**：POST

**是否需要鉴权：**是

**请求参数**：

| 参数名    | 参数类型 | 是否必填 | 参数说明       |
| --------- | -------- | -------- | -------------- |
| name      | String   | 是       | 出入库类型名称 |
| type_code | string   | 否       | 类型编码       |

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

**请求路径**：/invt/update

**请求方法**：PUT

**是否需要鉴权：**是

**请求参数**：

| 参数名    | 参数类型 | 是否必填 | 参数说明 |
| --------- | -------- | -------- | -------- |
| itid      | String   | 是       | 类型id   |
| name      | String   | 否       | 类型名称 |
| type_code | string   | 否       | 类型编码 |

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

**请求路径**：/invt/delete

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

**请求路径**：/invt/list

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

**请求路径**：/unit/add

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

### 

----

### 删除

**请求路径**：/unit/delete

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

**请求路径**：/unit/list

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







# 数据库建表示例

## user表

| 字段 |          uid           |   account    |   password   |  nick_name   |   permission    |  created_at  |     token      |    phone     |
| :--: | :--------------------: | :----------: | :----------: | :----------: | :-------------: | :----------: | :------------: | :----------: |
| 类型 |      varchar(20)       | varchar(255) | varchar(255) | varchar(255) |       int       | varchar(255) |  varchar(255)  | varchar(255) |
| 说明 | 标识+8位唯一索引(主键) |     账号     |     密码     |     昵称     | 权限（1，2，3） |  注册时间戳  | 登录生成的凭证 |     电话     |

----

## warehouse表

| 字段 |          wid           |     name     |  created_at  |   comment    |     manager     |  status  |
| :--: | :--------------------: | :----------: | :----------: | :----------: | :-------------: | :------: |
| 类型 |      varchar(20)       | varchar(255) | varchar(255) | varchar(255) |  varchar(255)   |   int    |
| 说明 | 标识+4位唯一索引(主键) |    仓库名    |  添加时间戳  |     备注     | 负责人sid(外键) | 仓库状态 |

----

## goods_type表

| 字段 |          gtid          |     name     |  type_code   |  created_at  |
| :--: | :--------------------: | :----------: | :----------: | :----------: |
| 类型 |      varchar(20)       | varchar(255) | varchar(255) | varchar(255) |
| 说明 | 标识+4位唯一索引(主键) |  货品类型名  | 货品类型编码 |  添加时间戳  |

----

## department表

| 字段 |          did           |     name     |  created_at  |
| :--: | :--------------------: | :----------: | :----------: |
| 类型 |      varchar(20)       | varchar(255) | varchar(255) |
| 说明 | 标识+4位唯一索引(主键) |   部门名称   |  添加时间戳  |

----

## staff表

| 字段 |          sid           |     name     |    phone     |      department      |  created_at  |
| :--: | :--------------------: | :----------: | :----------: | :------------------: | :----------: |
| 类型 |      varchar(20)       | varchar(255) | varchar(255) |     varchar(255)     | varchar(255) |
| 说明 | 标识+8位唯一索引(主键) |   员工名称   |     电话     | 员工所属部门id(外键) |  添加时间戳  |

----

## inventory_type表

| 字段 |          itid          |      name      |   type_code    |  created_at  |
| :--: | :--------------------: | :------------: | :------------: | :----------: |
| 类型 |      varchar(20)       |  varchar(255)  |  varchar(255)  | varchar(255) |
| 说明 | 标识+4位唯一索引(主键) | 出入库类型名称 | 出入库类型编码 |  添加时间戳  |

## goods表

| 字段 |       gid        |  goods_code  |     name     |    model     |   type_code    |   warehouse    | manufacturer |   quantity   |     unit     |
| :--: | :--------------: | :----------: | :----------: | :----------: | :------------: | :------------: | :----------: | :----------: | :----------: |
| 类型 |   varchar(20)    | varchar(255) | varchar(255) | varchar(255) |  varchar(255)  |  varchar(255)  | varchar(255) | varchar(255) | varchar(255) |
| 说明 | 标识+8位唯一索引 |   货品编码   |   货品名称   |   型号规格   | 货品类型(外键) | 存储仓库(外键) |    生产商    |     数量     |   计量单位   |

## unit表

| 字段 |       unid       |     name     |
| :--: | :--------------: | :----------: |
| 类型 |   varchar(20)    | varchar(255) |
| 说明 | 表示+3位唯一索引 |   单位名称   |

