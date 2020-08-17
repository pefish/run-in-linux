# Run-in-linux

一款可以切换到linux下执行命令的小工具。

当你正在使用Mac开发Go程序时，想在linux测试下执行效果，可以用到此工具

## Dependence

Docker

## Install

```shell script
GO111MODULE=on go get -u github.com/pefish/run-in-linux/cmd/...
```

## Quick start

```shell
run-in-linux go run ./bin/test
```


## Security Vulnerabilities

If you discover a security vulnerability, please send an e-mail to [pefish@qq.com](mailto:pefish@qq.com). All security vulnerabilities will be promptly addressed.

## License

This project is licensed under the [Apache License](LICENSE).
