@echo off

:SETVARS
SET CGO_ENABLED=0
SET GOARCH=386

CD /D .\src\

go tool dist install -v runtime
go install -v -a std


:: GOOS=windows GOARCH=386 go build -o appname.exe appname.go

:: GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o appname.linux appname.go

pause