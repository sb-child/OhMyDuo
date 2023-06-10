# OhMyDuo

~~你是 Duolingo 的忠实用户吗？~~

本 Repo 可以帮助你生成自定义的可爱贴纸！

## demo

|                                    数据                                    |
| :------------------------------------------------------------------------: |
| <img src="https://my.sbchild.top:81/_sbcounter/test/rw/card" width="100%"> |

|                                                           图解(硬核)                                                           |
| :----------------------------------------------------------------------------------------------------------------------------: |
| <img src="https://user-images.githubusercontent.com/55868015/169673597-322babec-b2fb-41ce-bf7b-c773970d7479.png" width="100%"> |

这些链接可供测试

- <https://my.sbchild.top:81/_sbcounter/t1/rw/card>
- <https://my.sbchild.top:81/_sbcounter/t2/rw/card>
- <https://my.sbchild.top:81/_sbcounter/t3/rw/card>
- <https://my.sbchild.top:81/_sbcounter/t4/rw/card>
- <https://my.sbchild.top:81/_sbcounter/t5/rw/card>

## 用途

记录访问量

## 怎么用

0. `clone`本项目, 或下载压缩包并解压

1. 创建数据库, 目前只支持`postgresql`:

```sql
CREATE USER "[用户名]" PASSWORD '[密码]';
CREATE DATABASE [数据库名] OWNER "[用户名]";
```

进入 `[数据库名]` 数据库

```sql
DROP TABLE public.counter;

CREATE TABLE IF NOT EXISTS public.counter
(
    id text COLLATE pg_catalog."default" NOT NULL,
    db text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone,
    ip text COLLATE pg_catalog."default" NOT NULL
)

TABLESPACE [表空间, 默认为 pg_default];

ALTER TABLE public.counter
    OWNER to "[用户名]";
```

2. 在 config 目录下, 复制`config.example.toml`到`config.toml`, 按需编辑`config.toml`:

- 更改数据库地址, 端口号, 名称, 用户名和密码
- 更改监听端口
- 更改`oh-my-duo`设置

3. 编译

- 准备好`golang 1.20+`和`gfcli`环境
- 拷贝`manifest/config/config.yaml`到`./config.yaml`
- `./build.sh`

4. 部署

- 确保`oh-my-duo-bin`和`config.yaml`在同一目录下
- 运行`oh-my-duo-bin`

5. 使用

- 写入模式, 每次访问都会计数: `http://127.0.0.1:端口号/根目录/用户名/rw/card`
- 只读模式, 只能查看, 不会记录: `http://127.0.0.1:端口号/根目录/用户名/ro/card`

## 配置文件是什么鬼

每个合法的链接都包含一个标识名（称为`path`），访问计数器将根据这个标识名来记录访问次数。

此外，每个路径都归属于一个分组，每个分组的访问记录都将存储到数据库中。

如果多个路径属于同一分组，它们的访问计数将会共享。例如，如果两个路径都属于 sbchild-profile 分组，它们的访问计数将会合并计算。

```yaml
sbcounter:
  # 省略
  user:
    - path: "sbchild" # 这个是路径的名字 例子: http://127.0.0.1:3264/_sbcounter/sbchild/rw/card
      db: "sbchild-profile" # 如果上述链接被访问过的话, 访问记录会保存在这个分组(sbchild-profile)里
    - path: "testuser" # 具有相同分组的链接, 看到的访问计数是一样的
      db: "sbchild-profile"
    - path: "aaa" # 这是另一个分组，它和上一个路径互不影响
      db: "2"
```

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
