@echo off

go build -buildmode=c-shared -o go.dll go.go

pause
