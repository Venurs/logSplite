# logSplite
该文件夹的可执行文件用来将服务器上产生的action.log按照年/月/日/action.log  文件的格式进行分割，便于日志的分析
  1. `spliteLogFilePool`  双线程，读 写，使用文件池的方式，适合没有顺序的文件读取
  2. `spliteLogSplice`   三线程，读 日志按照日期分割保存到切片 写，适合有顺序的日志文件，但是处理线程会大量将数据保存至内存，如果文件很大， 服务器内存不够大，最好还是用第一个吧

```
 ~/go_script/spliteLogFilePool -d `date -d yesterday "+%Y-%m-%d"` /mnt/logs/action-sw-prod1.log -s /mnt/prod-sw-action-log
```

测试：

|方式|文件大小|用时|

|`spliteLogSplice`|1.9G  | 28.52043948s|

|`spliteLogFilePool`|9.2G|5m35.955957661s|