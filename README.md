# citycloud.cf-deploy-ui

使用golang开发的cloudfoundry v2 图形化部署应用。

## config gopath

```
$ pwd

/path/to/path

$ export GOPATH=/path/to/path
```

## init workdir and checkout master

```
$ cd $GOPATH/src
$ go get github.com/gorilla/websocket
$ go get github.com/astaxie/beego
$ go get github.com/beego/bee

#for BDD-style testing
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega

# for more information [ginkgo](http://onsi.github.io/ginkgo/)
```

### attention(important!)

beego 目前使用1.5.0
go get安装完成后，需要检出v1.5.0重新编译安装beego，v1.6.0 缺失AppPath参数，影响应用编译。

```
$ cd $GOPATH/src/github.com/astaxie/beego
$ git checkout v1.5.0
$ go build 
$ go install

$ bee version
bee   :1.4.1
beego :1.5.0
Go    :go version go1.4.2 linux/amd64
```

## init workdir and checkout dev

```
$ mkdir github.com/citycloud
$ cd github.com/citycloud
$ git clone https://github.com/alex8023/citycloud.cf-deploy-ui.git
```

## checkout development version of the branch

```
$ git checkout dev
```

## update

```
$ git fetch
$ git pull
```



## commit and pull

```
$ git add <file>
$ git add .
$ git commit -m "commit message"
$ git push origin dev
```

## commit to gogit
```
$ git remote add gogit http://10.10.82.100:8800/cci-paas/citycloud.cf-deploy-ui.git

$ git push gogit dev --tag
#or
$ git push gogit master --tag
```

## init db

import createDb.sql,mysql.sql to mysql server 5.6