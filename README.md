# bubble
> gin框架记事本
> 参考地址：[https://github.com/Q1mi/bubble](https://github.com/Q1mi/bubble)


- grom\gin框架

流程处理

- 编写简单web执行程序
- 引入前端资源
   - 配置路由访问前端资源
- 设置路由组
- 配置数据源db连接池
- 编写web执行接口
- 进行模块拆分处理
   - MVC
## 项目结构
```
├── README.md
├── controller // 控制层
├── dao // 数据库相关基础服务
├── logic // 复杂业务逻辑处理
├── modules // 模型层的增删改查
├── routers // 路由层
├── static // 静态文件
├── templates // 模版文件
├── utils // 工具库
├── go.mod // 依赖文件
└── go.sum
```
### 执行流程
url ---> controller   ---> logic  ---> model

请求来了  -->  控制器-->  业务逻辑-->  模型层的增删改查

### 执行逻辑
