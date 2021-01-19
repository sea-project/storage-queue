
# storage-queue

基于栈实现的队列

## 目录
 序号 | Go文件/函数或方法 | 作用 
---|---|---
 1 | queue.go | 基于栈实现的队列存储
 2 | `Single`  | 单例初始化
 3 | `NewQueue`| 初始化队列
 4 | `Push`    | 塞入队列
 5 | `Pop`     | 从队列弹出
 6 | `Len`     | 获取队列长度


## 单元测试 

序号 | Go文件/测试用例方法 | 说明
---|---|---
 1 | queue_test.go | 数据库操作测试用例
 2 | `TestQueuePool_Push`  | 功能测试
 3 | `BenchmarkPool_Push`  | 性能测试