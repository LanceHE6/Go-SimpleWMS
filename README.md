<p align="center">
<img alt="logo" src="./imgs/WMS-Logo.png">
</p>
<h1 align="center" style="margin: 30px 0 30px; font-weight: bold;">Go-SimpleWMS</h1>
<h1 align="center" style="margin: 30px 0 30px; font-weight: bold;">简行云仓库</h1>
<div align="center"> 


简体中文/[English](./README-EN.md) 
</div>
<h4 align="center">基于gin+Vue3前后端分离式仓库管理系统</h4>
<div align="center">

![Static Badge](https://img.shields.io/badge/Licence-MIT-blue)
![Static Badge](https://img.shields.io/badge/前端-vue-orange)
![Static Badge](https://img.shields.io/badge/后端-gin-green)

</div>


## 平台简介

* 前端技术栈 [Vue3](https://v3.cn.vuejs.org) + [Element Plus](https://element-plus.org/zh-CN) + [Vite](https://cn.vitejs.dev) 。
* 后端[gin](https://gin-gonic.com/zh-cn/)+[gorm](https://gorm.io/zh_CN/docs/index.html)。
* 数据库[mysql]([MySQL](https://www.mysql.com/cn/))。

## 内置功能

### 权限组

#### 普通用户

* 浏览信息

#### 管理员

* 货品添加，更新
* 出入库单添加，更新

#### 超级管理员

* 所有增删改查权限

### 管理

* 用户管理

* 部门管理

* 计量单位管理

* 仓库管理

* 员工管理

* 出入库类型管理

* 货品管理

* 货品类型管理

* 出入库订单管理

  

## 部署

#### 直接部署

##### 前端

```bash
# 克隆项目
git clone https://github.com/LanceHE6/Go-SimpleWMS.git

# 进入项目目录
cd vue-web

# 安装依赖
npm install

# 启动服务
npm run dev

# 前端访问地址 http://localhost:5173
```

##### 后端

```bash
# 进入项目目录
cd go-server

# 下载依赖
go mod download

# 构建项目
go build -o main .
```

----

#### Docker部署

##### 前端

```bash
# 克隆项目
git clone https://github.com/LanceHE6/Go-SimpleWMS.git

# 进入项目目录
cd vue-web

# 构建镜像
docker build -t simple-wms-web:latest .

# 运行容器
docker run --name simple-wms-web -p 5173:80 -d simple-wms-web:latest
```

##### 后端端

```bash
# 进入项目目录
cd go-server

# 构建镜像
docker build -t simple-wms-server:latest .

# 运行容器
docker run --name simple-wms-server -p 8080:8080 -d simple-wms-server:latest
```



## 文档

[前端文档](./vue-web/README.md)

[后端文档](./go-server/README.md)

[接口文档](./go-server/api.md)


## 在线预览

[Go-SimpleWMS](https://lancehe6.github.io/Go-SimpleWMS/)



## 对标预期效果

![](./imgs/homepage.png)

![ruku](./imgs/ruku.png)

![setting](./imgs/setting.png)

![](./imgs/goods.png)



