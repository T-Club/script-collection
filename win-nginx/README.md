Nginx Windows脚本
===

## 主要功能
主要解决Nginx在Windows环境下环境变量问题， 运行时只要将 `nginx` 换成 `win-nginx` 就可以了，解决全局运行nginx。

## 用法
1. 修改脚本里 `nginx_dir`， 这个是指的是nginx目录，注意不能有空格。
1. 修改脚本里 `nginxexe`， 若nginx.exe在nginx目录下，就不需要修改了，若不在，在需要修改位nginx.exe所在位置。
1. 将该脚本加入环境变量。
1. 可以全局运行了。