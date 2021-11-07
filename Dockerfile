FROM golang:1.17-alpine AS build
USER root
WORKDIR /go/src/Realip
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add upx
COPY . .
RUN  go build -o app app.go && upx -9 app

# target images
FROM alpine
USER root
WORKDIR /Realip
LABEL name=Realip
LABEL url=https://github.com/zf1976/Realip

ENV TZ=Asia/Shanghai
COPY --from=build ./go/src/Realip/app ./app
EXPOSE 8080
CMD ["./app"]
