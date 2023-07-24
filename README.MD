# Oh My Duo

[![Go](https://github.com/sb-child/OhMyDuo/actions/workflows/go.yml/badge.svg)](https://github.com/sb-child/OhMyDuo/actions/workflows/go.yml) ![GitHub release](https://img.shields.io/github/v/release/sb-child/OhMyDuo)

## [> View the English version of README <](./translate/en/)

~~你是 [Duolingo](https://duolingo.com/) 的忠实用户吗？~~

本 Repo 可以帮助你生成自定义的可爱贴纸！

## demo

|                               默认内容                               |
| :------------------------------------------------------------------: |
| <img src="https://my.sbchild.top:81/omd/_?x=fj32j340if" width="50%"> |
|                  https://my.sbchild.top:81 `/omd/_`                  |

|                                                   更改角色，原文和翻译                                                   |
| :----------------------------------------------------------------------------------------------------------------------: |
| <img src="https://my.sbchild.top:81/omd/_?c=lily&o=do%20not%20talk%20with%20me&t=不要和我讲话&x=fyufyudtry" width="50%"> |
|                      https://my.sbchild.top:81 `/omd/_?c=lily&o=do not talk with me&t=不要和我讲话`                      |

你在用 Telegram 嘛? 快快试试 <https://t.me/sbchild_bot> 吧~

## 用途

制作可爱贴纸！还可以在自己的主页上展示！

## 怎么用

0. `clone`本项目, 或下载压缩包并解压

> 不想自己编译的可以去 Release ![GitHub release](https://img.shields.io/github/v/release/sb-child/OhMyDuo) 下载预编译版

1. 在 `manifest/config` 目录下, 复制`config.example.toml`到`config.toml`, 按需编辑`config.toml`:

- 更改 HTTP 监听端口
- 更改 `rootDir` 为 URL 地址前缀
- 更改 `httpServer`，如果不需要启用 HTTP 服务，请将值修改为 `false`
- 更改 `telegramBotToken`，如果不需要启用 Telegram Bot，请将值修改为空
- 更改 `telegramBotHttpTimeout`，推荐设为 `3000`
- 更改 `telegramBotHttpProxy`，如果不需要启用 Telegram 代理，请将值修改为空
- 更改 `telegramBotImageServer` 为 Telegram 使用的 OhMyDuo 地址

2. 编译

- 准备好`golang 1.20+`和`gfcli`环境
- `./build.sh`

3. 部署

- 确保`oh-my-duo-*`和`config.yaml`在同一目录下
- 运行`oh-my-duo-*`
- (可选) 按需修改`OhMyDuo.service`, 并拷贝至`/etc/systemd/system/OhMyDuo.service`, 即可创建一个 systemd 服务
- 接下来你可以尝试使用 Web 服务器的反向代理功能，可以方便使用，以及允许 Telegram 从这个服务器请求图片!

4. 测试

- 打开浏览器，输入`你的域名或ip地址/你的地址前缀/_`，例如`https://my.sbchild.top:81/omd/_`
- 如果浏览器显示了一张多儿的照片，恭喜你，它能用啦

5. 到底怎么用

目前有两个方式调用这个生成器:

1. 在`/_`后面添加参数, 第一个参数前加上`?`, 参数之间使用`&`分隔

- `/_` 默认
- `/_?c=lily` 指定角色，不区分大小写, 默认为 Duo
  - Duo
  - Bea
  - Vikram
  - Oscar
  - Junior
  - Eddy
  - Zari
  - Lily
  - Lin
  - Lucy
  - Falstaff
- `/_?o=text...` 原文。某些情况下，空格需要使用`%20`代替
- `/_?t=text...` 译文(原文和译文需要一起指定)。某些情况下，空格需要使用`%20`代替

2. 在`/_/`后面按照如下格式指定参数, 参数之间使用`|`分隔

- `/_/Hello-world|你好-世界` 默认使用 Duo 角色, 原文为`Hello world`, 译文为`你好 世界`。空格使用`-`代替

- `/_/Can-you-speak--English%3f|你会讲英语吗？` 原文为`Can you speak-English?`, 译文为`你会讲英语吗？`。问号`?`使用`%3f`代替, 连字符`-`使用`--`代替

- `/_/lucy|Can-you-speak--English%3f|你会讲英语吗？` 指定`Lucy`角色, 原文为`Can you speak-English?`, 译文为`你会讲英语吗？`

但是在 Telegram Bot 中，需要这样使用:

- `@some_bot Hello world|你好 世界` 列出所有角色的图片, 原文为`Hello world`, 译文为`你好 世界`

- `@some_bot Can u speak english|你能讲英语嘛|zari` 使用 Zari 角色, 原文为`Can u speak english`, 译文为`你能讲英语嘛`

## TODO

- [ ] emoji 支持 😊
- [x] 日文支持
- [ ] 自动截断超长单词
- [ ] 更多语言旗帜 (欢迎 PR)

## 贡献

欢迎提交 PR 来改进这个程序!

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
