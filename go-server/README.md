# 后端接口文档

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
	"message": "No Authorization header provided"
}
```

### 鉴权相关返回数据说明

| 参数名  | 参数类型 | 参数说明             |
| ------- | -------- | -------------------- |
| error   | string   | 后端服务内部错误消息 |
| message | string   | 请求返回消息         |

### 鉴权相关返回状态码说明

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 401    | StatusUnauthorized  | 鉴权未通过       |
| 403    | StatusForbidden     | 权限拒绝         |
| 500    | InternalServerError | 后端服务内部错误 |

----



## 用户

### 注册

**请求路径**：/user/register

**请求方法**：POST

**是否需要鉴权：**否

**请求参数**：

| 参数名        | 参数类型 | 是否必填 | 参数说明     |
|------------| -------- | -------- | ------------ |
| account    | String   | 是       | 用户名（手机号/邮箱） |
| password   | String   | 是       | 密码         |
| permission | int | 是 | 权限大小（1为管理员，2为超级管理员） |
| nick_name  | string | 是 | 昵称 |
| phone | string | 否 | 电话 |

**返回结果示例**：

```json
{
    "message": "User registered successfully",
    "uid": "00000002"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明             |
| ------- | -------- | -------------------- |
| message | string   | 返回消息             |
| uid     | string   | 注册成功返回的用户id |
| error   | string   | 后端内部错误消息     |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 注册成功         |
| 400    | BadRequest          | 请求参数不全     |
| 403    | Forbidden           | 账号已被注册     |
| 500    | InternalServerError | 后端服务内部错误 |

----



### 登录

**请求路径**：/user/login

**请求方法**：POST

**是否需要鉴权：**否

**请求参数**：

| 参数名  | 参数类型 | 是否必填 | 参数说明     |
| --------- | -------- | -------- | ------------ |
| username  | String   | 是       | 用户名（手机号/邮箱） |
| password  | String   | 是       | 密码         |

**返回结果示例**：

```json
{
    "message": "Login successfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk1NTUxMjEsInN1YiI6IjAwMDAwMDAyIn0.hvz-Xp9kfhVNsCy6Q9nhS9wM8-c-DgJJ8PLcME17Fto"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明                |
| ------- | -------- | ----------------------- |
| message | string   | 返回消息                |
| token   | string   | 登录成功返回的用户token |
| error   | string   | 后端内部错误消息        |

**返回状态码说明**

| 状态码 | 含义                 | 说明             |
| ------ | -------------------- | ---------------- |
| 200    | OK                   | 登录成功         |
| 203    | NonAuthoritativeInfo | 账号或密码不正确 |
| 400    | BadRequest           | 请求参数不全     |
| 500    | InternalServerError  | 后端服务内部错误 |

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
    "message": "User deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 登录成功         |
| 400    | BadRequest          | 请求参数不全     |
| 401    | Unauthorized        | 鉴权未通过       |
| 500    | InternalServerError | 后端服务内部错误 |

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
| nickName   | String   | 否       | 用户更改的昵称     |
| permission | int      | 否       | 用户更改的权限大小 |
| phone      | String   | 否       | 用户更改的电话     |

*注:password,nickName,permission,phone4个可选参数至少需提供一个*

**返回结果示例**：

```json
{
    "message": "User updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 修改成功         |
| 400    | BadRequest          | 请求参数不全     |
| 401    | Unauthorized        | 鉴权未通过       |
| 500    | InternalServerError | 后端服务内部错误 |

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
    "message": "Get user list successfully",
    "rows": [
        {
            "account": "admin",
            "nickName": "admin",
            "permission": 3,
            "phone": "1924658487",
            "register_time": "123456",
            "uid": "00000001"
        },
        {
            "account": "123",
            "nickName": "admin",
            "permission": 3,
            "phone": "",
            "register_time": "1711266620",
            "uid": "00000002"
        },
        {
            "account": "1234",
            "nickName": "admin",
            "permission": 3,
            "phone": "",
            "register_time": "1711267355",
            "uid": "00000003"
        },
        {
            "account": "1235",
            "nickName": "admin",
            "permission": 3,
            "phone": "",
            "register_time": "1711267524",
            "uid": "00000004"
        },
        {
            "account": "1236",
            "nickName": "admin",
            "permission": 3,
            "phone": "123456789",
            "register_time": "1711267576",
            "uid": "00000005"
        }
    ]
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| rows    | array[]  | 用户信息数组     |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 修改成功         |
| 401    | Unauthorized        | 鉴权未通过       |
| 500    | InternalServerError | 后端服务内部错误 |

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
    "message": "Warehouse added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 修改成功         |
| 400    | BadRequest          | 请求参数不全     |
| 401    | Unauthorized        | 鉴权未通过       |
| 403    | Forbidden           | 仓库已存在       |
| 500    | InternalServerError | 后端服务内部错误 |

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
    "message": "Warehouse deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 修改成功         |
| 400    | BadRequest          | 请求参数不全     |
| 401    | Unauthorized        | 鉴权未通过       |
| 500    | InternalServerError | 后端服务内部错误 |

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
    "message": "Warehouse updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明                        |
| ------ | ------------------- | --------------------------- |
| 200    | OK                  | 修改成功                    |
| 400    | BadRequest          | 请求参数不全                |
| 401    | Unauthorized        | 鉴权未通过                  |
| 403    | Forbidden           | 仓库名已经存在/该仓库不存在 |
| 500    | InternalServerError | 后端服务内部错误            |

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
    "message": "Get warehouse list successfully",
    "rows": [
        {
            "add_time": "1711267718",
            "comment": "123",
            "manager": "00000001",
            "name": "F",
            "status": 1,
            "wid": "000001"
        },
        {
            "add_time": "1711267858",
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

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| rows    | array[]  | 仓库信息数组     |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 修改成功         |
| 401    | Unauthorized        | 鉴权未通过       |
| 500    | InternalServerError | 后端服务内部错误 |

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
    "message": "Goods type added successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 修改成功         |
| 400    | BadRequest          | 请求参数不全     |
| 401    | Unauthorized        | 鉴权未通过       |
| 403    | Forbidden           | 仓库已存在       |
| 500    | InternalServerError | 后端服务内部错误 |

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
    "message": "Goods type updated successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明                        |
| ------ | ------------------- | --------------------------- |
| 200    | OK                  | 修改成功                    |
| 400    | BadRequest          | 请求参数不全                |
| 401    | Unauthorized        | 鉴权未通过                  |
| 403    | Forbidden           | 仓库名已经存在/该仓库不存在 |
| 500    | InternalServerError | 后端服务内部错误            |

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
    "message": "Goods type deleted successfully"
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 修改成功         |
| 400    | BadRequest          | 请求参数不全     |
| 401    | Unauthorized        | 鉴权未通过       |
| 500    | InternalServerError | 后端服务内部错误 |

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
    "message": "Get user list successfully",
    "rows": [
        {
            "addTime": "1711273694",
            "gtid": "0001",
            "name": "yzl",
            "type_code": "bzx"
        },
        {
            "addTime": "1711275118",
            "gtid": "0002",
            "name": "印字轮",
            "type_code": "yzl"
        }
    ]
}
```

**返回数据说明**

| 参数名  | 参数类型 | 参数说明         |
| ------- | -------- | ---------------- |
| message | string   | 返回消息         |
| rows    | array[]  | 仓库信息数组     |
| error   | string   | 后端内部错误消息 |

**返回状态码说明**

| 状态码 | 含义                | 说明             |
| ------ | ------------------- | ---------------- |
| 200    | OK                  | 修改成功         |
| 401    | Unauthorized        | 鉴权未通过       |
| 500    | InternalServerError | 后端服务内部错误 |

----

# 数据库建表示例

## user表

| 字段 |     uid      |   account    |   password   |  nick_name   |   permission    | register_time |     token      |    phone     |
| :--: | :----------: | :----------: | :----------: | :----------: | :-------------: | :-----------: | :------------: | :----------: |
| 类型 | varchar(255) | varchar(255) | varchar(255) | varchar(255) |       int       | varchar(255)  |  varchar(255)  | varchar(255) |
| 说明 | 8位唯一索引  |     账号     |     密码     |     昵称     | 权限（1，2，3） |  注册时间戳   | 登录生成的凭证 |     电话     |

----

## warehouse表

| 字段 |     wid      |     name     |   add_time   |   comment    |   manager    |  status  |
| :--: | :----------: | :----------: | :----------: | :----------: | :----------: | :------: |
| 类型 | varchar(255) | varchar(255) | varchar(255) | varchar(255) | varchar(255) |   int    |
| 说明 | 6位唯一索引  |    仓库名    |  添加时间戳  |     备注     |  负责人uid   | 仓库状态 |
