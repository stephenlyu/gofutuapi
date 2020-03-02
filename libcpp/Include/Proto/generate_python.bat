@echo off
chcp 65001
for /r %~dp0 %%h in (*.proto) do (
echo 转换：%%h
protoc.exe -I=%~dp0 --cpp_out=%~dp0 %%h
)
pause