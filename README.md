# citycloud.cf-deploy-ui

## config gopath

```
$ pwd

/path/to/path

$ export GOPATH=/path/to/path
```

## init workdir and checkout master

```
$ cd $GOPATH/src
$ mkdir github.com/citycloud
$ cd github.com/citycloud
$ git clone https://github.com/alex8023/citycloud.cf-deploy-ui.git
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
$ git commit -m "message"
$ git push origin master
```