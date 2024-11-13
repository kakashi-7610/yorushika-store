# 本来ならバージョン指定の方が良い
FROM golang:latest
WORKDIR /work
COPY ./ ./
RUN apt-get update
RUN apt-get install -y iputils-ping
RUN go mod tidy
RUN go build
CMD ./product-manager server start
