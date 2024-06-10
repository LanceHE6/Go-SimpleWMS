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
// 无token 400
{
    "code": 401,
	"message": "No Authorization header provided"
}
// 无效token 400
{
    "message": "Invalid token",
	"code":    401,
}
// 鉴权通过 200
{
    "code": 201,
	"uid":  uid,
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
// 注册成功 200
{
    "code": 201,
    "message": "User registered successfully",
    "uid": "00000002"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 账号已被注册 403
{
    "message": "The account '%s' has been registered",
	"code":    402,
}
// 无法注册新用户 500
{
    "error":  "Cannot insert new user",
	"detail": err.Error(),
	"code":   505,
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
// 部分导入失败 217
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
// 全部导入失败 400
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
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 全部导入成功 200
{
    "code":    201,
	"message": "All users have completed registration",
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
// 登录成功 200
{
    "code": 201,
    "message": "Login successfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk1NTUxMjEsInN1YiI6IjAwMDAwMDAyIn0.hvz-Xp9kfhVNsCy6Q9nhS9wM8-c-DgJJ8PLcME17Fto"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 无法生成token 500
{
    "error":  "Cannot generate token",
	"detail": err.Error(),
	"code":   501,
}
// 无法更新token 500
{
    "error":  "Cannot update token",
	"detail": err.Error(),
	"code":   502,
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
// 删除成功 200
{
    "code": 201,
    "message": "User deleted successfully"
}
// 参数有误 400 
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 非法删除对象 403
{
    "message": "Invalid target uid",
	"code":    402,
}
// 无法删除用户 500
{
    "error":  "Cannot delete user",
	"detail": err.Error(),
	"code":   501,
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
// 更新成功 200
{
    "code": 201,
    "message": "User updated successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 用户不存在 403
{
    "message": "The user does not exist",
	"code":    403,
}
// 无法更新用户 500
{
    "error":  "Cannot update user",
	"detail": err.Error(),
	"code":   501,
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
// 查询成功 200
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
// 无法获取列表 500
{
    "error":  "Cannot get the list of users",
	"detail": err.Error(),
	"code":   "501",
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
// 查询成功 200
{
    "code": 201,
    "message": "Warehouse added successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 该仓库已存在 409
{
    "message": "The warehouse already exists",
	"code":    402,
}
// 无法插入新仓库 500
{
    "error":  "Cannot insert the warehouse",
	"detail": err.Error(),
	"code":   501,
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
// 删除成功 200
{
    "code": 201,
    "message": "Warehouse deleted successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 无法删除仓库 500
{
    "error":  "Cannot delete the warehouse",
	"detail": err.Error(),
	"code":   501,
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
// 更新成功 200
{
    "code": 201,
    "message": "Warehouse updated successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 该仓库不存在 403
{
    "message": "The department does not exist",
	"code":    403,
}
// 无法更新仓库 500
{
    "error":  "Cannot update the warehouse",
	"detail": err.Error(),
	"code":   501,
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
// 查询成功 200
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
// 无法获取列表 500
{
    "error":  "Cannot get the list of warehouses",
	"detail": err.Error(),
	"code":   501,
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
// 添加成功 200
{
    "code": 201,
    "message": "Goods type added successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 该类型已存在 403
{
    "message": "The type name already exists",
	"code":    402,
}
// 无法插入新类型 500
{
    "error":  "Cannot insert the goods type",
	"detail": err.Error(),
	"code":   505,
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
// 更新成功 200
{
    "code": 201,
    "message": "Goods type updated successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 目标类型不存在 403 
{
    "message": "The goods type does not exist",
	"code":    403,
}
// 该类型名已存在 400
{
    "error":  "The name is already exists",
	"detail": err.Error(),
	"code":   404,
}
// 无法更新类型 500
{
    "error":  "Cannot update the goods type",
	"detail": err.Error(),
	"code":   504,
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
// 删除成功 200
{
    "code": 201,
    "message": "Goods type deleted successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 无法删除 500
{
    "error":  "Cannot delete the Goods type",
	"detail": err.Error(),
	"code":   502,
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
// 查询成功 200
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
// 无法获取列表 500
{
	"error":  "Cannot get the list of goods type",
	"detail": err.Error(),
	"code":   501,
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
// 添加成功 200
{
    "code": 201,
    "message": "Department added successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 该部门已经存在
{
    "message": "The department name already exists",
	"code":    402,
}
// 无法插入新部门 500
{
    "error":  "Cannot insert new department",
	"detail": err.Error(),
	"code":   501,
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
// 更新成功 200
{
    "code": 201,
    "message": "Department updated successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 目标部门不存在 403
{
    "message": "The department does not exist",
	"code":    403,
}
// 部门已存在 400
{
    "error":  "The name is already exists",
	"detail": err.Error(),
	"code":   402,
}
// 无法更新 500
{
    "error":  "Cannot update department",
	"detail": err.Error(),
	"code":   501,
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
// 删除成功 200
{
    "code": 201,
    "message": "Department deleted successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 无法删除 500
{
    "error":  "Cannot delete department",
	"detail": err.Error(),
	"code":   501,
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
// 查询成功 200
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
// 无法获取列表 500
{
    "error":  "Can not get the list of departments",
	"detail": err.Error(),
	"code":   201,
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
// 添加成功 200
{
    "code": 201,
    "message": "Staff added successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 员工已存在 403
{
    "message": "The staff already exists",
	"code":    402,
}
// 员工所属部门不存在 403
{
    "message": "The staff's department does not exist",
	"code":    403,
}
// 无法插入新员工 500
{
    "error":  "Cannot insert the staff",
	"detail": err.Error(),
	"code":   501,
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
// 更新成功 200
{
    "code": 201,
    "message": "Staff updated successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 目标员工不存在 403
{
    "message": "The staff does not exist",
	"code":    403,
}
// 员工所属部门不存在 403
{
    "message": "The staff's department does not exist",
	"code":    404,
}
// 无法更新员工 500
{
    "error":  "Cannot update staff",
	"detail": err.Error(),
	"code":   501,
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
// 删除成功 200
{
    "code": 201,
    "message": "Staff deleted successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 无法删除 500
{
    "error":  "Cannot delete staff",
	"detail": err.Error(),
	"code":   501,
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
// 查询成功 200
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
// 无法获取列表 500
{
    "error":  "Cannot get the list of staffs",
	"detail": err.Error(),
	"code":   502,
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
// 添加成功 200
{
    "code": 201,
    "message": "Inventory type added successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 该类型已存在 403
{
    "message": "The type name already exists",
	"code":    401,
}
// 无法插入新类型 500
{
    "error":  "Cannot insert the inventory type",
	"detail": err.Error(),
	"code":   505,
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
// 更新成功 200
{
    "code": 201,
    "message": "Inventory type updated successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 目标类型不存在 403
{
    "message": "The inventory type does not exist",
	"code":    403,
}
// 更新失败 500
{
    "error":  "Cannot get the number of inventory type for this itid",
	"detail": err.Error(),
	"code":   502,
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
// 删除成功 200
{
    "code": 201,
    "message": "Inventory type deleted successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 无法删除 500
{
    "error":  "Cannot delete the inventory type",
	"detail": err.Error(),
	"code":   501,
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
// 查询成功 200
{
    "code": 201,
    "message": "Get inventory type list successfully",
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
// 无法获取列表 500
{
    "error":  "Cannot get the list of inventory type",
	"detail": err.Error(),
	"code":   501,
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
// 添加成功 200
{
    "code": 201,
    "message": "Unit added successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 单位已存在 403
{
    "message": "The unit already exists",
	"code":    402,
}
// 无法插入新单位 500
{
    "error":  "Cannot insert new unit",
	"detail": err.Error(),
	"code":   501,
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
// 更新成功 200
{
    "code": 201,
    "message": "unit type updated successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 目标单位不存在 403
{
    "message": "The unit does not exist",
	"code":    403,
}
// 无法更新 500
{
    "error":  "Cannot update unit",
	"detail": err.Error(),
	"code":   501,
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
// 查询成功 200
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
// 无法获取列表 500
{
    "error":  "Can not get the list of units",
	"detail": err.Error(),
	"code":   201,
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
    "message": "Goods added successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 已存在该编码的货品 403
{
    "message": "The goods with this code already exists",
	"code":    402,
}
// 无法插入新货品 500
{
    "error":  "Cannot insert the goods",
	"detail": err.Error(),
	"code":   501,
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
// 更新成功 200
{
    "code": 201,
    "message": "Goods updated successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 已存在该编码的货品 400
{
    "error":  "The code of the goods is already exists",
	"detail": err.Error(),
	"code":   404,
}
// 无法更新货品 500
{
    "message": "Update goods failed",
	"code":    403,
	"detail":  err.Error(),
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
// 删除成功 200
{
    "code": 201,
    "message": "Goods deleted successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 无法删除 500
{
    "error":  "Cannot delete the Goods",
	"detail": err.Error(),
	"code":   502,
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
// 查询成功（有数据） 200
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
// 查询成功（无数据） 200
{
   "code":        202,
	"message":     "No data",
	"page":        1,
	"page_size":   10,
	"total":       0,
	"total_pages": 0,
	"rows":        [], 
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
// 上传成功 200
{
    "code": 201,
    "data": {
        "gid": "g07a99cb7",
        "image_name": "goods_1716190218.png",
        "image_path": "static/res/goodsImage/goods_1716190218.png"
    },
    "message": "Upload successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 无法创建文件夹 500
{
    "error":  "Upload failed",
	"detail": err.Error(),
	"code":   501,
}
// 图片格式不支持 400
{
    "message": "Unsupported image format",
	"code":    402,
}
// 货品不存在 400
{
    "error":  "The goods does not exist",
	"detail": err.Error(),
	"code":   403,
}
// 无法更新货品表 500
{
    "error":  "Updating database failed",
	"detail": err.Error(),
	"code":   504,
}
// 保存图片失败 500
{
    "error":  "Upload failed",
	"detail": ioErr.Error(),
	"code":   501,
}
// 无法创建图片文件 500
{
    "error":  "Upload failed",
	"detail": err.Error(),
	"code":   502,
}
// 无法从请求中读取图片文件 500
{
    "error":  "Upload failed",
	"detail": err.Error(),
	"code":   503,
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
// 添加成功 200
{
    "code": 201,
    "message": "Inventory added successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 货品列表参数格式错误 400
{
    "message": "The format of the goods_list is incorrect",
	"code":    402,
	"detail":  err.Error(),
}
// 出入库类型不存在 400
{
    "message": "The inventory type does not exist",
	"code":    402,
}
// 日期格式错误 400
{
    "message": "The format of the date is incorrect",
	"code":    403,
	"detail":  err.Error(),
}
// 无法插入新单据 500
{
    "error":  "Cannot insert new inventory",
	"code":   501,
	"detail": err.Error(),
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
// 更新成功 200
{
    "code": 201,
    "message": "Inventory added successfully"
}
// 参数有误 400
{
    "message": "Missing parameters or incorrect format",
	"code":    401,
	"detail":  err.Error(),
}
// 货品列表参数格式错误 400
{
    "message": "The format of the goods_list is incorrect",
	"code":    402,
	"detail":  err.Error(),
}
// 启动数据库事务失败 500
{
    "message": "Failed to start transaction",
	"code":    501,
	"detail":  tx.Error.Error(),
}
// 无法获取旧单据记录 500
{
    "message": "Failed to get original inventory record",
	"code":    502,
	"detail":  err.Error(),
}
// 无法获取单据对应的出入库类型 500
{
    "message": "Failed to get inventory type",
	"code":    503,
	"detail":  err.Error(),
}
// 原库存数据不存在 500
{
    "message": "Contains invalid Stock records",
	"code":    403,
}
// 无法更新库存 500
{
    "message": "Failed to update stock",
	"code":    504,
	"detail":  err.Error(),
}
// 无法获取新单据出入库类型 500
{
    "message": "Failed to get new inventory type",
	"code":    505,
	"detail":  err.Error(),
}
// 新库存数据造成库存不足 400
{
    "message": "Contains invalid Stock records",
	"code":    405,
}
// 单号已存在 400
{
    "error":  "The Number already exists",
	"detail": err.Error(),
	"code":   404,
}
// 更新失败 500
{
    "message": "Failed to update inventory record",
	"code":    505,
	"detail":  err.Error(),
}
// 无法获取新单据记录 500
{
    "message": "Failed to get updated inventory record",
	"code":    506,
	"detail":  err.Error(),
}
// 提交数据库事务失败 500
{
    "message": "Failed to commit transaction",
	"code":    508,
	"detail":  err.Error(),
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
// 查询成功（有数据） 200
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
// 查询成功（无数据） 200
{
    "code":        202,
	"message":     "No data",
	"page":        1,
	"page_size":   10,
	"total":       0,
	"total_pages": 0,
	"rows":        [],
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



## 库存

### 查询

**请求路径**：/api/stock/get

**请求方法**：GET

**是否需要鉴权：**是

| 参数名    | 参数类型 | 是否必填 | 参数说明    |
| --------- | -------- | -------- | ----------- |
| warehouse | string   | 是       | 仓库id(wid) |
| goods     | string   | 否       | 货品id(gid) |

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
    "message": "Get stock success"
}
// 查询成功（1对多） 202
{
    "code": 202,
    "message": "Get stock success",
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
// 参数缺失 400
{
    "message": "Missing parameters",
	"code":    400,
}
// 获取记录失败 500
{
    "message": "Get stock failed",
	"code":    500,
}
```

**返回数据说明**

| 参数名  | 参数类型 |   参数说明   |
| :-----: | :------: | :----------: |
|  code   |   int    |    业务码    |
| message |  string  |   返回消息   |
|  data   |  string  |   返回数据   |
|  rows   |  string  | 返回数据列表 |

**返回状态码说明**

| 状态码 |        含义         |       说明       |
| :----: | :-----------------: | :--------------: |
|  200   |         OK          |     修改成功     |
|  401   |    Unauthorized     |    鉴权未通过    |
|  500   | InternalServerError | 后端服务内部错误 |
