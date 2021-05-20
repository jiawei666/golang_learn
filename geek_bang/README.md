## Geek Bang
> 极客时间go语言入门到实战学习

#### 切片
1. 切片结构组成
    ![](./img/切片结构.png)
   
2. 切片共享内存原理
    ![](./img/切片共享内存.png)
   
#### 字符
1. 字符本质是[]byte的数组，但是只读，所以不能通过修改数组的值来改变字符
2. len()会统计字符的byte长度
3. 遍历数组时会把每个字符存放到rune中，因为这样可以兼容utf8

#### 函数
1. 函数是一等公民
![](./img/函数是一等公民.png)

#### 面向对象
1. GO是一门面向对象编程语言？
![](./img/官方说法.png)

2. 接口的结构
    >注意：数据保存方式是一个地址
![](./img/接口结构.png)
   
3. 接口使用的是Duck Type
![](./img/ducktype.png)

3. 接口最佳实践
![](./img/接口最佳实践.png)
   1. 尽量试用小接口
   2. 使用小接口组合成大接口
   3. 只依赖必要功能的小接口
    
#### 错误机制
1. 要点
![](./img/错误机制.png)
   
2. panic
![](./img/panic.png)
   
#### 包管理
1. package
![](./img/package.png)
   
2. init函数
![](./img/package_init.png)
   
3. 获取远程package
```shell
go get ...
```

#### 并发编程
1. channel机制
![](./img/channel.png)