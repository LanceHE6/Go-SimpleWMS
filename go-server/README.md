# 后端接口文档

## 鉴权说明

### 鉴权方式：

请求头携带鉴权JWT token

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

| 参数名  | 参数类型 | 是否必填 | 参数说明     |
| --------- | -------- | -------- | ------------ |
| username  | String   | 是       | 用户名（手机号/邮箱） |
| password  | String   | 是       | 密码         |
| permission | int | 是 | 权限大小（1为管理员，2为超级管理员） |
| nick_name | string | 是 | 昵称 |

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

*注:password,nickName,permission3个参数至少需提供一个*

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

# 数据库建表示例

## user表

| 字段 | uid         | account | password | nick_name | permission      | register_time | token          |
| :--: | ----------- | ------- | -------- | --------- | --------------- | ------------- | -------------- |
| 类型 | text        | text    | text     | text      | int             | text          | text           |
| 说明 | 8位唯一索引 | 账号    | 密码     | 昵称      | 权限（1，2，3） | 注册时间戳    | 登录生成的凭证 |
