# Start from the latest golang base image
FROM golang:1.21

ENV TZ=Asia/Shanghai

# Add Maintainer Info
LABEL maintainer="Hycer Lance <2765543491@qq.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o server .

# build后删除除server以外的所有文件
RUN rm -rf `ls | grep -v "server"`

# 修改权限
RUN chmod +x server

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the environment variable for the app
ENV MYSQL_ACCOUNT=root
ENV MYSQL_PASSWORD=root
ENV MYSQL_HOST=127.0.0.1
ENV MYSQL_PORT=3306
ENV MYSQL_DBNAME=simple_wms

ENV SMTP_HOST=smtp.qq.com
ENV SMTP_PORT=465
ENV SMTP_USERNAME=your-email@qq.com
ENV SMTP_PASSWORD=your-email-password

VOLUME [ "/static" ]

# Command to run the executable
CMD ["./server"]
