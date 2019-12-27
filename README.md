# ddbfc
Distributed DNS brute force cracking   分布式Dns暴力破解
![Master](https://s2.ax1x.com/2019/12/27/lVmgm9.png)

work初始化 注册到Discory    
master 初始化字典  将字典进行分割成任务队列  下发个给work  将当前下发任务放入执行任务中

work完成任务 向master返回   master向执行任务中表删除结束任务

如果执行中任务超时  广播删除此任务 再重新下发任务

### 这是Cli 本地并发版
字典放入dic目录下 会自动检测 (如果没有会使用基础字典)

``` 
dollarkiller@dollarkiller-virtual-machine:~/Github/ddbfc$ ./ddbfc -h
NAME:
   DdbFC - Distributed DNS brute force cracking CLi

USAGE:
   ddbfc [global options] command [command options] [arguments...]

VERSION:
   0.1

AUTHOR:
   DollarKiller <adapawang@gmail.com>

COMMANDS:
   scan     start to crack weak password
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --domain value, -d value   domain
   --timeout value, -t value  Single DNS query timeout Millisecond (default: 400)
   --tryNum value, -r value   Number of attempts (default: 3)
   --max value, -m value      Maximum number of concurrency (default: 200)
   --help, -h                 show help
   --version, -v              print the version
```

### 引入依赖
- 发现全球高可用公共dnsList  github.com/dollarkillerx/publicDns

后面想了想就把这个做成单独的服务抽离出去

黑名单 (伊朗  奥地利 微软 乱解析到的IP)  要过滤调
"208.91.112.55"  "10.10.34.35"  "213.94.80.190"

### 分支
- cli-new  cli发现严重的性能问题 而开的重构分支