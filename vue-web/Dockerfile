# 使用 Node.js 官方提供的 Docker 镜像作为基础镜像
FROM node:lts-alpine as build-stage

# Add Maintainer Info
LABEL maintainer="Hycer Lance <2765543491@qq.com>"

# 设置工作目录
WORKDIR /app

# 复制项目文件到 Docker 容器的工作目录中
COPY package*.json ./

# 安装项目依赖
RUN npm install

# 复制项目的源代码到 Docker 容器的工作目录中
COPY . .

# 构建生产版本的项目
RUN npm run build

# 创建一个用于运行应用的镜像阶段
FROM nginx:stable-alpine as production-stage
# 复制构建的生产版本的项目到 Nginx 的默认网站目录中
COPY --from=build-stage /app/dist /usr/share/nginx/html

# 需要自定义 Nginx 配置，可以将配置文件复制到容器中
# COPY nginx.conf /etc/nginx/nginx.conf

# 暴露端口
EXPOSE 80

# 启动 Nginx 服务
CMD ["nginx", "-g", "daemon off;"]