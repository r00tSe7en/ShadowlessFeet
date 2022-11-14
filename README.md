# ShadowlessFeet 

## 无影脚 - 命令行下的日志文件处理工具

渗透过程中遇到删除日志文件中包含(IP\路径\UA...)的文本行时，Linux下有各种自带命令去处理(当然本程序也同样适用Linux)。

然而Windows下遇到较大的日志文件处理起来却十分棘手，此工具正是为了解决这个痛点。

![](img/1.jpeg)

## 使用帮助

-file 日志文件的路径

-key  要匹配的关键字，匹配成功的文本行会自动删除

```
# ./main -h
Usage of ./main:
  -file string
        log file path
  -key string
        keywords to match
```

## 示例效果

```
# cat test.txt
aaaaaaaaaaaa
bbbbbbbbbbbb
abc ,123~!@#$%^&*()_+{}|
cccccccccccc
dddddddddddd

# ./main -file test.txt -key "abc ,123"
delete :  1  //这里显示的是删除的行数
spend :  1.061807ms

# cat test.txt
aaaaaaaaaaaa
bbbbbbbbbbbb
cccccccccccc
dddddddddddd
```
## 更新

2022/11/15 - 重构代码逻辑，优化初始版本的大量重复开关文件操作，大文件处理速度得到明显提升。
