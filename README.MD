# OhMyDuo

~~你是 Duolingo 的忠实用户吗？~~

本 Repo 可以帮助你生成自定义的可爱贴纸！

## demo

(即将上线)

## 用途

制作可爱贴纸！

## 怎么用

0. `clone`本项目, 或下载压缩包并解压

1. 在 config 目录下, 复制`config.example.toml`到`config.toml`, 按需编辑`config.toml`:

- 更改监听端口和服务器 URL 地址前缀
- 更改`oh-my-duo`设置

2. 编译

- 准备好`golang 1.20+`和`gfcli`环境
- `./build.sh`

3. 部署

- 确保`oh-my-duo-bin`和`config.yaml`在同一目录下
- 运行`oh-my-duo-bin`

4. 使用

todo

## 框架

[GoFrame](https://goframe.org)

## 许可证

Apache 2.0

# Makefile 子命令

Project Makefile Commands:

- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.
