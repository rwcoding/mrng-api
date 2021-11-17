# MRNG网关管理面板及一个简单地配置中心

## 后端Api
### 依赖
> MySql   
> Redis （可选，作配置中的集群方案时可使用redis提高性能）  

### 构建
```
go build -gcflags=-m -ldflags="-w -s" -o tmp/mrng.exe main.go static_handle.go
```

> 构建结果在 `tmp` 目录下  
> 遵循go项目的常规构建方式，可自行编写  
> 构建时会将前端文件一起打包  
> 所有接口使用 `post` 请求  
> 后端配置，请参照 `mrng.sample.toml` 文件  
> 数据库请自行导入，可查询发行版本的SQL文件  
> 默认账号密码 "admin" "admin"  

### 运行
```
mrng --conf /etc/mrng.toml
```

## 配置中心
+ 环境  必须，指 开发、生产、灰度等自定义环境，每个环境配备 `访问密钥`
+ 工程  必须，配置必须从属于工程，每个工程配备 `访问密钥`
+ 配置  具体的配置项目，每一此配置变更都会将该环境下的配置整体版本号 `+1`
+ 配置中心提供 http 接口，客户端查询都是从缓存中获取
+ 配置缓存
    + MySql 默认MySql会保存一份缓存表 `config_kv`
    + Redis 如果配置了Redis，每个Redis都会保存一份
    + 所有缓存都是同步的
