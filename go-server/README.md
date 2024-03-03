# 后端接口文档


## 用户

### 注册

**请求路径**：/user/register

**请求方法**：POST

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

### 登录

**请求路径**：/user/login

**请求方法**：POST

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