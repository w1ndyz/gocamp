#### 问题一：总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用

1. `fix length`

> 定长型包, 每个包有固定大小，主要适用于单一类型的包

优点: 打包、解包算法简单

缺点：

* 包长度小于定长时会有浪费

* 包长度修改时，服务端和客户端需要同时修改

2. `delimiter based`

> 变长型包，每个包使用特殊字符为结束符

优点：基本可以按照实际大小传输

缺点：

* 特殊符号不能在body中出现，否则会导致解包和封包不一致
* 传递二进制内容容易出错

3. `length field based frame decoder`

> 变长型包，固定位+变长位

优点：灵活，组包没有限制，适用于二进制传输

缺点：每个包都有固定位，包含了包标识符，包长度，变长的body，增加了许多要传的内容

#### 问题二: 实现一个从 socket connection 中解码出 goim 协议的解码器。
查看`ch09/goim`项目