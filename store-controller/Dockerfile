# 本来ならバージョン指定の方が良い
FROM golang:latest
WORKDIR /work
COPY ./ ./
RUN apt-get update
RUN apt-get install -y iputils-ping
RUN go mod tidy
# CMD sh waitingMysql.sh
CMD go run main.go server start
