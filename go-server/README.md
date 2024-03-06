# 后端接口文档


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

