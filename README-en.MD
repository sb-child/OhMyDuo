# Oh My Duo

## [> 查看中文版 README <](./README.MD)

~~Are you a loyal Duolingo user? ~~

This repository can help you generate custom Duolingo stickers!

## demo

|                           default content                            |
| :------------------------------------------------------------------: |
| <img src="https://my.sbchild.top:81/omd/_?x=fj32j340if" width="50%"> |
|                  https://my.sbchild.top:81 `/omd/_`                  |

|                                  change the character, origin text and translated text                                   |
| :----------------------------------------------------------------------------------------------------------------------: |
| <img src="https://my.sbchild.top:81/omd/_?c=lily&o=do%20not%20talk%20with%20me&t=不要和我讲话&x=fyufyudtry" width="50%"> |
|                      https://my.sbchild.top:81 `/omd/_?c=lily&o=do not talk with me&t=不要和我讲话`                      |

## What is it

Make some cute stickers, and show it on your site easily!

## How to use

0. `clone` or download this repository

1. In `manifest/config`, copy `config.example.toml` to `config.toml`, then edit it:

- change the HTTP server address
- change the `rootDir` of `ohmyduo`

2. Compile

- Install `golang 1.20+` and `gfcli` (execute `make cli` command to install)
- Run `./build.sh`

3. Deploy

- On your server, make sure `oh-my-duo-*` and `config.yaml` are at the same directory
- Run `oh-my-duo-*`
- (optional) modify `OhMyDuo.service`, then copy to `/etc/systemd/system/OhMyDuo.service` to create a systemd service
- Now you can configure the reverse proxy settings of your web server

4. Test

- Open your web browser and open `your-domain-or-ip/your-root-dir/_`, for example `https://my.sbchild.top:81/omd/_`
- If it shows a picture of Duo, then it finally works!

5. Generate your own sticker

There are two ways to use this service for now:

1. Add some parameters after `/_`, add `?` before the first parameter, use `&` to separate every parameters

- `/_` default parameters
- `/_?c=lily` set the character, the default is Duo
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
- `/_?o=text...` the origin text, replace the space to `%20`
- `/_?t=text...` the translated text

2. Add some parameters after `/_/`, use `|` to separate every parameters

- `/_/Hello-world|你好-世界` use the default character, origin text is `Hello world`, translated text is `你好 世界`. replace the space to `-`

- `/_/Can-you-speak--English%3f|你会讲英语吗？` origin text is `Can you speak-English?`, translated text is `你好 世界`. replace `?` to `%3f`, and replace `-` to `--`

- `/_/lucy|Can-you-speak--English%3f|你会讲英语吗？` set the character to `Lucy`, origin text is `Can you speak-English?`, translated text is `你好 世界`

## TODO

- [ ] split very long words automatically
- [ ] more language flags (PR welcome)

## Contributing

Pull requests are welcome to improve this program

## Framework

[GoFrame](https://goframe.org)

## License

Apache 2.0

## Makefile sub commands

Project Makefile Commands:

- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.
