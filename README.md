# go-slim 
## 使用成熟开源组件组合形成的 web 框架。
#### · 基础服务（如db、redis等）通过 NewServiceContext 注入至 domain.ServiceContext 容器中。
#### · delivery.Register() 实例化handler，注册路由（restful）并绑定相应handler。
## 开源组件：gin、gorm、yaml.v3、zap
