# colly
博文标题搜索

搜索范围包含至今为止所有的文章（共712条）

## Setup
```
git clone https://github.com/NullpointerW/colly.git
cd colly
```

```
export GOPROXY=goproxy.io
go mod download
```

* 运行

```
 go run 
```

## Use
 * 搜索
 ```
curl http://localhost:8964/v1/{search title}
```
 * 正则搜索
 ```
curl http://localhost:8964/v2/{regexp}
```
* 示例

![example](https://raw.githubusercontent.com/NullpointerW/colly/master/example.png)


## 注
* 数据文件保存在/cache/dump.json ,如果不存在则会爬取博客网页转存为dump.json(墙内需要代理）
* 搜索匹配到多条时显示列表,可通过link访问；只有一条则直接跳转到对应博客页
