@echo off

set GOOS=windows
set CGO_ENABLED=1
go build -buildmode=c-shared -o go.dll go.go

pause
