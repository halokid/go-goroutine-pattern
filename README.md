###  golang协程的应用场景和范式


**纪念逝去的时光和迎接下一刻的重生**

这里只示范golang并行的范式写法， 并不讨论性能区别

> 文件说明
- just_run_tasks           单纯并行执行任务
- get_run_tasks_res        需要获取执行结果 
- run_tasks_with_order     按照顺序返回执行结果 
- goroutine_timeout        设置goroutine超时时间
- cancel_goroutine         根据某一个goroutine的执行结果取消其他goroutine 






