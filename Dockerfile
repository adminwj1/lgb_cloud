FROM golang:latest
COPY ./ /www/wwwroot/server/cloud/
WORKDIR /www/wwwroot/server/cloud/
EXPOSE 8080
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -ldflags="-w -s" -o main main.go
RUN chmod 755 ./main
CMD ./main