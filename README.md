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
```

## init workdir and checkout dev

```
$ mkdir github.com/citycloud
$ cd github.com/citycloud
$ git clone https://github.com/alex8023/citycloud.cf-deploy-ui.git
```

#checkout development version of the branch
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
