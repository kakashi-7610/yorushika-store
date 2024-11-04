
@echo off

rem start product-manager
cd ../product-manager
start server_start.bat

rem start store-controller
cd ../store-controller
go run main.go server start
