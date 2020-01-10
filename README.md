# exporter-controller

exporter-controller作为CRD，旨在为提供了metrics接口的业务pod注入nginx作为反向代理，来实现拉取metrics的认证。

本地编译运行
```shell script
# make && make install && make run
```

制作镜像/发布/运行
```shell script
打包镜像
# make docker-build

#发布镜像
# make docker-push

部署
# make deploy
```
