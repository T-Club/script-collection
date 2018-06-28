@echo off
set nginx_dir=C:\nginx-1.13.8\
set nginxexe=%nginx_dir%nginx.exe

set run_nginx = %nginx_dir %, -p , %nginx_dir%
if "%1" == "" (
  start %nginxexe%  -p %nginx_dir%
) else %nginxexe% -p %nginx_dir% %*

set nginx_dir=
set nginxexe=