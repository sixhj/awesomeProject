package main

/* 2019-09-20 liu

当你命名一个包时，通过使用关键字package，你只需要提供单个值，而不是一个完整的层次结构

但是当你导入一个包时，你需要指定一个全路径

包名和文件夹名是相同的

避免出现依赖循环的情况，可以一个目录，类似 common

go语言使用了一种简单的规则来规定类型或者函数是否对外部的包可见。
如果你命名  类型或者函数  时以一个  大写字母  开头，那么这个类型和函数就是   可见的。
如果使用一个  小写字母  开头，那么就是  不可见  的。

结构体的字段也使用相同的方式。
如果一个结构体的字段名是以小写字母开头的，那么只有在同一个包中的代码才能访问这些字段

执行go get -u将更新你的包（或者你可以通过go get -u FULL_PACKAGE_NAME更新指定的包）

*/

/* 2019-10-12 liu
包名尽量与所在文件夹的名字一致

变量名、函数名、方法名   =》 驼峰法
首字母大写，表示公有的
*/

func main() {

}