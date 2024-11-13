
@echo off

rem start product-manager
cd ../product-manager
start server_start.bat

rem start store-controller
cd ../store-controller
go build
store-controller server start
