# wxcloudrun-golang



## note

- [Limitations of the GET method in HTTP - Dropbox](https://dropbox.tech/developers/limitations-of-the-get-method-in-http)
- 勿在 `init` 中进行数据插入
- GORM

https://gorm.io/zh_CN/docs/

https://www.bilibili.com/video/BV1E64y1472a



## 目录结构说明

~~~
.
├── Dockerfile                Dockerfile 文件
├── LICENSE                   LICENSE 文件
├── README.md                 README 文件
├── container.config.json     微信云托管流水线配置
├── db                        数据库逻辑目录
├── go.mod                    go.mod 文件
├── go.sum                    go.sum 文件
├── index.html                主页 html 
├── main.go                   主函数入口
└── service                   接口服务逻辑目录
~~~



## 服务 API 文档

### `POST /api/book`

获取书籍信息

#### 请求参数

- `action` - `string` 类型
  - `exact` - 精确
  - `fuzzy` - 模糊
  - `category` - 分类
- `hint` - `string` 类型
  - `name`
  - `categoryId` - int 字面量

##### 请求参数示例

```json
{
  "action": "exact",
  "hint": "Models of Computation"
}
```

#### 响应结果

- `code` 错误码
- `data` 对象或对象数组
- `errorMsg` 错误信息

##### 响应结果示例

```json
{
  "code": -1,
  "data": null,
  "errorMsg": "record not found"
}
```

#### 调用示例

```
curl -X POST -H 'content-type: application/json' -d '{"action": "exact", "hint": "Models of Computation"}' https://<云托管服务域名>/api/book
```

使用更友好的 `httpie`

```
http https://<云托管服务域名>/api/book action=exact hint="Models of Computation"
```

### `POST /api/category`

获取书籍分类信息

#### 请求参数

- `action` - `string` 类型
  - `single` - 根据 ID 返回对应分类
  - `all` - 返回全部分类
- `hint` - `string` 类型
  - `categoryId` - int 字面量

#### 响应结果

- `code` 错误码
- `data` 对象或对象数组
- `errorMsg` 错误信息
