# OhMyDuo

~~你是 Duolingo 的忠实用户吗？~~

本 Repo 可以帮助你生成自定义的可爱贴纸！

## demo

| 默认内容 |
|:-----: |
| <img src="https://my.sbchild.top:81/omd/_" width="50%"> |
| https://my.sbchild.top:81 `/omd/_` |

| 更改角色，原文和翻译 |
| :------: |
| https://my.sbchild.top:81 `/omd/_?c=lily&o=do not talk with me&t=不要和我讲话` |
| [不知道为什么它加载失败...点击这个链接试一下吧~](<https://my.sbchild.top:81/omd/_?c=lily&o=do not talk with me&t=不要和我讲话>) |

## 用途

制作可爱贴纸！还可以在自己的主页上展示！

## 怎么用

0. `clone`本项目, 或下载压缩包并解压

1. 在 config 目录下, 复制`config.example.toml`到`config.toml`, 按需编辑`config.toml`:

- 更改 HTTP 监听端口
- 更改`oh-my-duo`的 URL 地址前缀

2. 编译

- 准备好`golang 1.20+`和`gfcli`环境
- `./build.sh`

3. 部署

- 确保`oh-my-duo-*`和`config.yaml`在同一目录下
- 运行`oh-my-duo-*`
- (可选) 按需修改`OhMyDuo.service`, 并拷贝至`/etc/systemd/system/OhMyDuo.service`, 即可创建一个 systemd 服务

4. 测试和部署

- 打开浏览器，输入`你的域名或ip地址/你的地址前缀/_`，例如`https://my.sbchild.top:81/omd/_`
- 如果浏览器显示了一张多儿的照片，恭喜你，它能用啦
- 接下来你可以尝试把它部署到服务器上，使用Web服务器的反向代理功能，让你拥有属于自己的**OhMyDuo**!

## 框架

[GoFrame](https://goframe.org)

## 许可证

Apache 2.0

## Makefile 子命令

Project Makefile Commands:

- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.
