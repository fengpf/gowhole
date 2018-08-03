
1 启动调试程序（gdb）
   gdb main

2  在main函数上设置断点
   (gdb) b main.main

3  带参数启动程序（r）
   (gdb) r arg1 arg2

4  在文件dbgTest.go上通过行号设置断点（b）
   (gdb) b main.go:16

5  查看断点设置情况（info b）
    (gdb) info b

6  禁用断点（dis n）
    (gdb) dis 1   
    (gdb) info b

7  删除断点（del n)
   (gdb) del 1
   (gdb) info b

8  断点后继续执行（c）
  (gdb) c

9  显示代码（l)
   (gdb) l

10 单步执行（n）
   (gdb) n
   DBGTestRun Begin!

11 打印变量信息（print/p）
    (gdb) l 17
    (gdb) p var1 
    $3 = 1

12 查看调用栈（bt），切换调用栈（f n），显示当前栈变量信息
    (gdb) bt
    (gdb) f 1
    (gdb) l
    (gdb) print var1 
    $5 = 1
    (gdb) print var2
    $6 = "golang dbg test"
    (gdb) print var3
    $7 =  []int = {1, 2, 3}

13 显示goroutine列表（info goroutines）
   (gdb) n
    23        waiter.Add(1)
    (gdb) info goroutines
    * 1 running  runtime.systemstack_switch
    2 waiting  runtime.gopark
    17 waiting  runtime.gopark
    18 waiting  runtime.gopark
    19 runnable GoWorks/GoDbg/mylib.RunFunc1

14  查看goroutine的具体情况（goroutine n cmd）
    (gdb) goroutine 19 bt
    (gdb) goroutine 19 info args
        variable = 1
        waiter = 0xc8200721f0
    (gdb) goroutine 19 p waiter 
        $1 = (struct sync.WaitGroup *) 0xc8200721f0
    (gdb) goroutine 19 p *waiter 
        $2 = {state1 = "\000\000\000\000\001\000\000\000\000\000\000", sema = 0}
    (gdb) n
        26  waiter.Add(1)
    (gdb) info goroutines
    (gdb) goroutine 19 bt

15 查看本地变量信息（info locals)
    (gdb) info locals