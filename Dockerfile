FROM golang:1.17-alpine AS build
USER root
WORKDIR /go/src/Realip
COPY ./go.mod ./go.mod
RUN export GOPROXY=https://goproxy.io
#RUN go list -m all
#RUN go mod download
COPY . .
RUN go build app.go

# target images
FROM alpine
LABEL name=Realip
LABEL url=https://github.com/zf1976/Realip

USER root
WORKDIR /Realip
ENV TZ=Asia/Shanghai
COPY --from=build ./app /go/src/Realip/app
EXPOSE 8080
RUN ./app

