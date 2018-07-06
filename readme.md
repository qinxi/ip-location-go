IP地址查询

- go build
 - [windows] go build -o build/ip-location.exe
 - [linux] go build -o build/ip-location
- gox build
 - gox -output="build/ip-location_{{.OS}}_{{.Arch}}"


运行参数

-port 监听端口 默认8080

-datafile   17monipdb.datx数据文件位置 默认程序当前目录

请求 http://host:port/?ip=8.8.8.8


