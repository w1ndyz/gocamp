#### 问题一
使用`redis-benchmark` 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
#### 解答
redis-benchmark命令行工具的使用
| 选项      | 描述                                   | 默认值    |
| --------- | -------------------------------------- | --------- |
| -h        | 指定服务器主机名                       | 127.0.0.1 |
| -p        | 端口                                   | 6379      |
| -s        | 指定服务器socket                       |           |
| -c        | 指定并发连接数                         | 50        |
| -n        | 指定请求数                             | 10000     |
| -k        | 1=keep alive; 0=reconnect              | 1         |
| -d        | 以字节的形式指定GET/SET值的数据大小    | 2         |
| -r        | SET/GET/INCR使用随机key,SADD使用随机值 |           |
| -P        | 通过管道传输<numreq>请求               | 1         |
| --csv     | 以csv格式输出                          |           |
| -q        | 强制退出redis。仅显示query/sec值       |           |
| -l        | 生成循环，永久进行测试                 |           |
| -t        | 仅运行以逗号分隔的测试命令列表         |           |
| -I(大写i) | idle模式。仅打开n个idle连接并等待      |           |


```shell
# 清空redis
redis-cli -p 6379 flushall
# 测试10字节
redis-benchmark -q -d 10 -r 1000000 -n 500000 -t set,get -p 6379 
SET: 46507.30 requests per second, p50=0.839 msec
GET: 43286.30 requests per second, p50=0.791 msec
# 测试20字节 
redis-benchmark -q -d 20 -r 1000000 -n 500000 -t set,get -p 6379 
SET: 43029.26 requests per second, p50=0.879 msec
GET: 46829.63 requests per second, p50=0.791 msec
# 测试50字节
redis-benchmark -q -d 50 -r 1000000 -n 500000 -t set,get -p 6379 
SET: 40371.42 requests per second, p50=0.935 msec
GET: 31772.26 requests per second, p50=0.847 msec
# 测试100字节
redis-benchmark -q -d 100 -r 1000000 -n 500000 -t set,get -p 6379 
SET: 41781.57 requests per second, p50=0.927 msec
GET: 41431.89 requests per second, p50=0.839 msec
# 测试200字节
redis-benchmark -q -d 200 -r 1000000 -n 500000 -t set,get -p 6379 
SET: 35589.72 requests per second, p50=0.919 msec
GET: 45884.19 requests per second, p50=0.815 msec
# 测试1k字节
redis-benchmark -q -d 1024 -r 1000000 -n 500000 -t set,get -p 6379
SET: 27791.68 requests per second, p50=1.023 msec
GET: 35767.94 requests per second, p50=0.887 msec 
# 测试5k字节
redis-benchmark -q -d 5120 -r 1000000 -n 500000 -t set,get -p 6379 
SET: 29139.23 requests per second, p50=1.287 msec
GET: 36284.47 requests per second, p50=1.015 msec
```
观察的结果：
1. 随着GET/SET的key大小越来越大，所耗费的读写时间也会越来越大
2. redis的读的性能还是很高的
3. 在并发读写的场景下，大key的读写很影响性能

#### 问题二
写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
#### 解答
分别写入10字节和1k查看内存平均每个key的内存占用
| 请求数 | key大小 | keys数量 | 内存    | 均值 |
| ------ | ------- | -------- | ------- | ---- |
| 50w    | 10      | 393346   | 41.15M  | 109  |
| 50w    | 1024    | 393183   | 611.07M | 1629 |

#### Redis存储细节
##### Redis并没有实现自己的内存池。
  * dicEntry: redis中的每个键值对都会对应一个dicEntry，里面存储了`key`/`value`对应的指针，指针指向相应的位置。`next`指向下一个dicEntry。
  * `key`存储在SDS结构中
  * 在redis中的value的多种类型都是通过redisObject存储的。里面包含`type`，即redis的数据类型。以及`ptr`指针，指向内存地址。该地址仍然是SDS存储结构。

##### Redis在编译时会指定内存分配器，可以是libc、jemalloc或者tcmalloc。
* jemalloc是redis默认的内存分配器，由facebook推出。在减小内存分配碎片方面做得很好。jemalloc在64位系统中，将内存空间划分为小、大、巨大三个范围。每个范围内又划分了许多小的内存块，当存储数据时，会选择大小最合适的内存块进行存储。
* tcmalloc是google推出的。
* libc是标准的内存分配库malloc和free。

