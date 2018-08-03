
1 带参数启动程序（dlv exec ./main -- arg1 arg2）
  $ dlv exec ./main -- arg1 arg2 

2 在main函数上设置断点（b）
  (dlv) b main.main

3 启动调试，断点后继续执行（c）

4 在文件main.go上通过行号设置断点（b）
  (dlv) b main.go:17

5 显示所有断点列表（bp）
  (dlv) bp

6 删除某个断点（clear x）
  (dlv) clear 5
  (dlv) bp

7 显示当前运行的代码位置（ls）
  (dlv) ls

8 查看当前调用栈信息（bt）
  (dlv) bt

9 输出变量信息（print/p）
  (dlv) print var1

10 在第n层调用栈上执行相应指令（frame n cmd）
  (dlv) frame 1 ls

11 查看goroutine的信息（goroutines）
  (dlv) goroutines

12 进一步查看goroutine信息（goroutine x）
   (dlv) goroutine 6
   (dlv) bt
   (dlv) bt 13 //通过bt加depth参数，设定bt的输出深度，进而找到我们自己的调用栈
   (dlv) frame 12 ls //通过frame x cmd就可以输出我们想要的调用栈信息了

13 查看当前是在哪个goroutine上（goroutine）
   (dlv) goroutine