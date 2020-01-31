# Go Distributed Reptiles
- 相信这个项目，对于学习分布式系统或者爬虫的你来说，帮助是巨大的。
- 项目的思路来自于慕课网的ccmouse老师[链接](https://coding.imooc.com/learn/list/180.html)
- 微服务教程[学习吧](https://study.163.com/course/courseMain.htm?courseId=1209482821)
## 技术栈
Go, Protobuf, Consul, Docker, Elasticsearch
### 必须要做的
- git clone https://github.com/apple-han/island.git
- cd island
- 全局搜索192.168.31.231 换成你主机的IP地址(这里因为有json文件,不好做全局的配置)

### Docker的方式部署
- cd crawler_distributed/persist
    - make build
    - make docker
    
- cd crawler_distributed/worker
    - make build
    - make docker
    
- cd crawler_distributed
    - make build
    - make docker
    
- cd crawler/frontend
    - make build
    - make docker
    
- cd island  
- docker-compose up -d
- http://192.168.31.231:8888/search?q=大众(自己的ip)


### 小贴士
1. 由于系统是一个分布式的,所以整体下来 还是有一点难度
2. 大家好好看一下吧,欢迎大家多多的PR

