# wxcloudrun-golang

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
- `data` json 对象
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