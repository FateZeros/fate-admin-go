# fateAdmin 后台代码

## 技术栈

golang、cobra

## 项目实现功能

基于 casbin 的 RBAC 权限控制，借鉴了 go-admin 项目的前端权限管理，可以在页面对 API、菜单、页面按钮等操作，进行灵活且简单的配置。

- [x] 登录
  - [ ] 岗位权限
  - [ ] 角色权限
- [x] 注册
- [ ] 登出
- [x] ldap 统一认证
- [x] 集成 swagger api 文档

## 参考

[Go 语言搬砖 ldap 统一认证服务](https://juejin.cn/post/7030968139924013087) </br>
[go-ldap-admin](https://github.com/eryajf/go-ldap-admin) </br>
