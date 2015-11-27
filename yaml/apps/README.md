# 云上应用部署

## 配置文件
### 配置文件名称：manifest.yml,内容如下
```
---
applications:
- name: app-name
  memory: 128M
  instances: 1
  host: hello
  domain: app.example.com
  path: path/to/app
```
### 多个app时

```
---
applications:
- name: app-name
  memory: 128M
  instances: 1
  host: hello
  domain: app.example.com
  path: path/to/app
- name: app-name2
  memory: 128M
  instances: 1
  host: hello
  domain: app.example.com
  path: path/to/app
```

## 配置文件属性

### name
```
app 名称
```
### memory
```
app 内存
```
### instances
```
app 实例数量
```
### host || hosts
```
app host、hosts
e.g host: hello

hosts:
- hello
- world
```
### domain || domains

domain 必须先创建，才可以配置

```
app domain，domains
e.g domain: app.example.com

domains:
- hello.example.com
- world.example.com
```
### path
```
app war路径
path: ./example.war
```
### stack
```
app 运行时系统
默认 lucid64
可选windows2012
```
### services
```
services:
- services1
- services2
```
### buildpack
```
app 运行时
```
### disk_quota
```
app 磁盘配额
```
### command
```
app 其他命令
```
### timeout
```
app 启动超时时间
```
### no-route
```
app
bool值
```
### no-hostname
```
app
bool值
```
### random-route
```
app
bool值
```

### health-check-type
### env
```
app 环境变量
e.g
env: 
  foo1: boo1
  foo2: boo2
```
所有属性都可以通过cf push --help查看具体说明